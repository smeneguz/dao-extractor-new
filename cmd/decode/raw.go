package decode

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

const batchSize = 1000

// contractInfo holds the parsed ABI and metadata for a contract.
type contractInfo struct {
	Address string
	ChainID string
	ABI     abi.ABI
}

func newRawCmd(rootCmd *cobra.Command) *cobra.Command {
	var dbURL string
	var daoFilter string

	cmd := &cobra.Command{
		Use:   "raw",
		Short: "Decode raw events into structured decoded_events table",
		Long: `Reads raw_events from the database, matches each event against its
contract ABI (from the abis table), decodes the event parameters,
and stores the result in the decoded_events table.

Requires ABIs to be populated first. Use 'abi download' or 'abi fetch-all'
to download ABIs from Etherscan for each contract.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDecode(cmd.Context(), dbURL, daoFilter)
		},
	}

	cmd.Flags().StringVar(&dbURL, "db", "postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable", "Database URL")
	cmd.Flags().StringVar(&daoFilter, "dao", "", "Comma-separated DAO symbols to decode (empty = all)")

	return cmd
}

func runDecode(ctx context.Context, dbURL, daoFilter string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	// Find all unique contracts in raw_events that have ABIs.
	contracts, err := loadContractsWithABIs(ctx, db, daoFilter)
	if err != nil {
		return fmt.Errorf("load contracts: %w", err)
	}

	if len(contracts) == 0 {
		fmt.Println("No contracts found with ABIs. Run 'abi fetch-all' first to download ABIs from Etherscan.")
		return nil
	}

	fmt.Printf("Found %d contracts with ABIs\n", len(contracts))

	totalDecoded := 0
	totalSkipped := 0
	totalErrors := 0

	for _, ci := range contracts {
		n, skipped, errs, err := decodeContractEvents(ctx, db, ci)
		if err != nil {
			return fmt.Errorf("decode events for %s: %w", ci.Address, err)
		}
		totalDecoded += n
		totalSkipped += skipped
		totalErrors += errs
	}

	fmt.Printf("\nDone: %d decoded, %d skipped (already decoded or unknown event), %d errors\n",
		totalDecoded, totalSkipped, totalErrors)
	return nil
}

// loadContractsWithABIs finds all unique (contract_address, chain_id) pairs
// in raw_events that have a matching ABI in the abis table.
func loadContractsWithABIs(ctx context.Context, db *sql.DB, daoFilter string) ([]contractInfo, error) {
	query := `
	SELECT DISTINCT a.address, b.chain_id, ab.abi
	FROM raw_events re
	JOIN addresses a ON a.id = re.contract_address_id
	JOIN blockchains b ON b.id = re.chain_id
	JOIN abis ab ON LOWER(ab.contract_address) = LOWER(a.address) AND ab.chain_id = b.chain_id
	`
	if daoFilter != "" {
		symbols := strings.Split(daoFilter, ",")
		quoted := make([]string, len(symbols))
		for i, s := range symbols {
			quoted[i] = fmt.Sprintf("'%s'", strings.TrimSpace(strings.ToUpper(s)))
		}
		query += fmt.Sprintf(`
		JOIN daos d ON d.id = re.dao_id
		WHERE UPPER(d.symbol) IN (%s)`, strings.Join(quoted, ","))
	}

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []contractInfo
	for rows.Next() {
		var addr, chainID string
		var abiJSON []byte
		if err := rows.Scan(&addr, &chainID, &abiJSON); err != nil {
			return nil, err
		}

		parsed, err := abi.JSON(bytes.NewReader(abiJSON))
		if err != nil {
			fmt.Printf("WARNING: failed to parse ABI for %s on chain %s: %v\n", addr, chainID, err)
			continue
		}

		contracts = append(contracts, contractInfo{
			Address: addr,
			ChainID: chainID,
			ABI:     parsed,
		})
	}

	return contracts, rows.Err()
}

// decodeContractEvents decodes all raw events for a specific contract.
func decodeContractEvents(ctx context.Context, db *sql.DB, ci contractInfo) (decoded, skipped, errors int, err error) {
	// Count total events to process.
	var total int
	err = db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM raw_events re
		JOIN addresses a ON a.id = re.contract_address_id
		JOIN blockchains b ON b.id = re.chain_id
		WHERE LOWER(a.address) = LOWER($1) AND b.chain_id = $2
		AND NOT EXISTS (
			SELECT 1 FROM decoded_events de
			WHERE de.tx_hash = re.tx_hash AND de.log_index = re.log_index
		)
	`, ci.Address, ci.ChainID).Scan(&total)
	if err != nil {
		return 0, 0, 0, err
	}

	if total == 0 {
		return 0, 0, 0, nil
	}

	fmt.Printf("  Decoding %s (chain %s): %d events... ", ci.Address[:10]+"...", ci.ChainID, total)

	// Process in batches.
	offset := 0
	for offset < total {
		rows, err := db.QueryContext(ctx, `
			SELECT re.id, d.symbol, b.chain_id, a.address,
				re.topics, re.data,
				re.tx_hash, re.block_height, re.log_index, re.ts, re.dao_id
			FROM raw_events re
			JOIN addresses a ON a.id = re.contract_address_id
			JOIN blockchains b ON b.id = re.chain_id
			JOIN daos d ON d.id = re.dao_id
			WHERE LOWER(a.address) = LOWER($1) AND b.chain_id = $2
			AND NOT EXISTS (
				SELECT 1 FROM decoded_events de
				WHERE de.tx_hash = re.tx_hash AND de.log_index = re.log_index
			)
			ORDER BY re.block_height, re.log_index
			LIMIT $3 OFFSET $4
		`, ci.Address, ci.ChainID, batchSize, offset)
		if err != nil {
			return decoded, skipped, errors, err
		}

		batchDecoded := 0
		for rows.Next() {
			var (
				rawID       int64
				daoSymbol   string
				chainID     string
				contractAddr string
				topicsJSON  []byte
				data        string
				txHash      string
				blockHeight int64
				logIndex    int
				ts          time.Time
				daoID       int64
			)
			if err := rows.Scan(&rawID, &daoSymbol, &chainID, &contractAddr,
				&topicsJSON, &data, &txHash, &blockHeight, &logIndex, &ts, &daoID); err != nil {
				rows.Close()
				return decoded, skipped, errors, err
			}

			// Parse topics from JSONB.
			var topics []string
			if err := json.Unmarshal(topicsJSON, &topics); err != nil {
				errors++
				continue
			}

			// Must have at least topic0.
			if len(topics) == 0 {
				skipped++
				continue
			}

			// Decode the event.
			eventName, eventSig, params, ok := decodeEvent(ci.ABI, topics, data)
			if !ok {
				skipped++
				continue
			}

			paramsJSON, err := json.Marshal(params)
			if err != nil {
				errors++
				continue
			}

			// Store decoded event.
			_, err = db.ExecContext(ctx, `
				INSERT INTO decoded_events(
					raw_event_id, dao_id, chain_id, contract_address,
					event_name, event_signature, decoded_params,
					tx_hash, block_height, log_index, ts)
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
				ON CONFLICT (tx_hash, log_index) DO NOTHING
			`, rawID, daoID, chainID, contractAddr,
				eventName, eventSig, paramsJSON,
				txHash, blockHeight, logIndex, ts)
			if err != nil {
				errors++
				continue
			}

			decoded++
			batchDecoded++
		}
		rows.Close()

		// If no new events were decoded in this batch, stop.
		if batchDecoded == 0 {
			break
		}

		offset += batchSize
	}

	fmt.Printf("%d decoded\n", decoded)
	return decoded, skipped, errors, nil
}

// decodeEvent decodes a raw event using the contract ABI.
// Returns the event name, full signature, decoded parameters, and success flag.
func decodeEvent(contractABI abi.ABI, topics []string, hexData string) (string, string, map[string]interface{}, bool) {
	topic0 := common.HexToHash(topics[0])

	// Find the event by topic0 (event ID = keccak256 of event signature).
	event, err := contractABI.EventByID(topic0)
	if err != nil {
		return "", "", nil, false
	}

	params := make(map[string]interface{})

	// Decode non-indexed parameters from data.
	if len(hexData) > 0 {
		dataBytes := common.Hex2Bytes(hexData)
		if len(dataBytes) > 0 {
			nonIndexed := make(map[string]interface{})
			if err := event.Inputs.UnpackIntoMap(nonIndexed, dataBytes); err == nil {
				for k, v := range nonIndexed {
					params[k] = formatParam(v)
				}
			}
		}
	}

	// Decode indexed parameters from topics[1:].
	indexedArgs := make(abi.Arguments, 0)
	for _, input := range event.Inputs {
		if input.Indexed {
			indexedArgs = append(indexedArgs, input)
		}
	}

	if len(indexedArgs) > 0 && len(topics) > 1 {
		topicHashes := make([]common.Hash, 0, len(topics)-1)
		for _, t := range topics[1:] {
			topicHashes = append(topicHashes, common.HexToHash(t))
		}

		indexedValues := make(map[string]interface{})
		if err := abi.ParseTopicsIntoMap(indexedValues, indexedArgs, topicHashes); err == nil {
			for k, v := range indexedValues {
				params[k] = formatParam(v)
			}
		}
	}

	return event.Name, event.Sig, params, true
}

// formatParam converts ABI-decoded values to JSON-safe representations.
func formatParam(v interface{}) interface{} {
	switch val := v.(type) {
	case common.Address:
		return val.Hex()
	case common.Hash:
		return val.Hex()
	case []byte:
		return common.Bytes2Hex(val)
	case fmt.Stringer:
		return val.String()
	default:
		return v
	}
}
