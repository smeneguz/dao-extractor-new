package abi

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// contractEntry represents a unique contract needing an ABI.
type contractEntry struct {
	Address string
	ChainID string
}

func NewFetchAllCmd() *cobra.Command {
	var dbURL string
	var apiKey string

	cmd := &cobra.Command{
		Use:   "fetch-all",
		Short: "Download ABIs for all contracts in raw_events that don't have one yet",
		Long: `Queries the database for all unique contract addresses in raw_events
that don't have a matching ABI in the abis table, then downloads
each ABI from Etherscan and stores it.

Uses Etherscan V2 API which supports all EVM chains with a single key.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFetchAll(cmd.Context(), dbURL, apiKey)
		},
	}

	cmd.Flags().StringVar(&dbURL, "db", "postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable", "Database URL")
	cmd.Flags().StringVar(&apiKey, "api-key", "", "Etherscan API key (required)")
	cmd.MarkFlagRequired("api-key")

	return cmd
}

func runFetchAll(ctx context.Context, dbURL, apiKey string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	// Find contracts missing ABIs.
	rows, err := db.QueryContext(ctx, `
		SELECT DISTINCT a.address, b.chain_id
		FROM raw_events re
		JOIN addresses a ON a.id = re.contract_address_id
		JOIN blockchains b ON b.id = re.chain_id
		WHERE NOT EXISTS (
			SELECT 1 FROM abis ab
			WHERE LOWER(ab.contract_address) = LOWER(a.address)
			AND ab.chain_id = b.chain_id
		)
		ORDER BY b.chain_id, a.address
	`)
	if err != nil {
		return fmt.Errorf("query missing ABIs: %w", err)
	}
	defer rows.Close()

	var missing []contractEntry
	for rows.Next() {
		var ce contractEntry
		if err := rows.Scan(&ce.Address, &ce.ChainID); err != nil {
			return err
		}
		missing = append(missing, ce)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	if len(missing) == 0 {
		fmt.Println("All contracts already have ABIs.")
		return nil
	}

	fmt.Printf("Found %d contracts without ABIs\n\n", len(missing))

	client := &http.Client{Timeout: 15 * time.Second}
	fetched := 0
	errors := 0

	for _, ce := range missing {
		fmt.Printf("  Fetching ABI for %s (chain %s)... ", ce.Address, ce.ChainID)

		abiJSON, err := fetchABIFromEtherscan(ctx, client, apiKey, ce.ChainID, ce.Address)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			errors++
			// Rate limit backoff.
			if strings.Contains(err.Error(), "rate") || strings.Contains(err.Error(), "429") {
				time.Sleep(6 * time.Second)
			} else {
				time.Sleep(1 * time.Second)
			}
			continue
		}

		// Store in DB.
		_, err = db.ExecContext(ctx, `
			INSERT INTO abis(contract_address, chain_id, abi)
			VALUES ($1, $2, $3)
			ON CONFLICT (contract_address, chain_id) DO UPDATE SET abi = $3
		`, ce.Address, ce.ChainID, abiJSON)
		if err != nil {
			fmt.Printf("DB ERROR: %v\n", err)
			errors++
			continue
		}

		fmt.Println("OK")
		fetched++

		// Etherscan free tier: ~5 calls/second.
		time.Sleep(250 * time.Millisecond)
	}

	fmt.Printf("\nDone: %d fetched, %d errors, %d total\n", fetched, errors, len(missing))

	if errors > 0 {
		return fmt.Errorf("%d ABI downloads failed", errors)
	}
	return nil
}

type etherscanResp struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func fetchABIFromEtherscan(ctx context.Context, client *http.Client, apiKey, chainID, address string) ([]byte, error) {
	// Etherscan V2 API supports all chains via chainid parameter.
	u, _ := url.Parse("https://api.etherscan.io/v2/api")
	q := u.Query()
	q.Set("chainid", strings.TrimPrefix(chainID, "0x"))
	q.Set("module", "contract")
	q.Set("action", "getabi")
	q.Set("address", address)
	q.Set("apikey", apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 429 {
		return nil, fmt.Errorf("rate limited (429)")
	}

	var result etherscanResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if result.Status != "1" || result.Message != "OK" {
		return nil, fmt.Errorf("etherscan: %s", result.Result)
	}

	return []byte(result.Result), nil
}
