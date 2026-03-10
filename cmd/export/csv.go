package export

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// tableExport defines a dataset to export.
type tableExport struct {
	Filename string
	Query    string
}

// allExports returns the full list of datasets to export.
// Base tables first, then materialized views.
func allExports() []tableExport {
	return []tableExport{
		// --- Core entities ---
		{
			Filename: "daos.csv",
			Query:    `SELECT id, symbol, name FROM daos ORDER BY id`,
		},
		{
			Filename: "blockchains.csv",
			Query:    `SELECT id, chain_id, name, type FROM blockchains ORDER BY id`,
		},
		// --- Governance ---
		{
			Filename: "proposals.csv",
			Query: `SELECT p.id, d.symbol AS dao_symbol, p.dao_id, p.proposal_id, p.chain_id,
				a_creator.address AS creator_address, a_contract.address AS contract_address,
				p.title, p.description, p.type, p.status,
				p.creation_height, p.creation_tx_hash, p.creation_ts,
				p.gas_used, (p.gas_fees).amount AS gas_fees_amount, (p.gas_fees).denom AS gas_fees_denom,
				p.start_height, p.start_time, p.end_height, p.end_time,
				p.quorum, p.extra_metadata
			FROM proposals p
			JOIN daos d ON d.id = p.dao_id
			JOIN addresses a_creator ON a_creator.id = p.creator_address_id
			JOIN addresses a_contract ON a_contract.id = p.contract_address_id
			ORDER BY p.dao_id, p.creation_ts`,
		},
		{
			Filename: "proposal_finalizations.csv",
			Query: `SELECT pf.id, d.symbol AS dao_symbol, p.proposal_id AS onchain_proposal_id,
				pf.tx_hash, pf.height, pf.ts,
				pf.gas_used, (pf.gas_fees).amount AS gas_fees_amount,
				pf.status_triggered, pf.extra_metadata
			FROM proposal_finalizations pf
			JOIN proposals p ON p.id = pf.proposal_id
			JOIN daos d ON d.id = p.dao_id
			ORDER BY pf.ts`,
		},
		{
			Filename: "vote_actions.csv",
			Query: `SELECT va.id, d.symbol AS dao_symbol, p.proposal_id AS onchain_proposal_id,
				a_voter.address AS voter_address,
				a_contract.address AS contract_address,
				COALESCE(a_delegator.address, '') AS delegator_address,
				va.tx_hash, va.height, va.ts,
				va.action_type, va.vote, va.voting_power,
				va.extra_metadata
			FROM vote_actions va
			JOIN proposals p ON p.id = va.proposal_id
			JOIN daos d ON d.id = p.dao_id
			JOIN addresses a_voter ON a_voter.id = va.sender_address_id
			JOIN addresses a_contract ON a_contract.id = va.contract_address_id
			LEFT JOIN addresses a_delegator ON a_delegator.id = va.delegator_address_id
			ORDER BY va.ts`,
		},
		// --- Token events ---
		{
			Filename: "token_transfers.csv",
			Query: `SELECT tt.id, d.symbol AS dao_symbol,
				tt.chain_id, tt.token_address,
				tt.from_address, tt.to_address, tt.amount,
				tt.tx_hash, tt.block_height, tt.log_index, tt.ts
			FROM token_transfers tt
			JOIN daos d ON d.id = tt.dao_id
			ORDER BY tt.ts`,
		},
		{
			Filename: "delegation_events.csv",
			Query: `SELECT de.id, d.symbol AS dao_symbol,
				de.chain_id, de.token_address,
				de.delegator, de.from_delegate, de.to_delegate,
				de.tx_hash, de.block_height, de.log_index, de.ts
			FROM delegation_events de
			JOIN daos d ON d.id = de.dao_id
			ORDER BY de.ts`,
		},
		{
			Filename: "delegate_votes_changed.csv",
			Query: `SELECT dvc.id, d.symbol AS dao_symbol,
				dvc.chain_id, dvc.token_address,
				dvc.delegate, dvc.previous_balance, dvc.new_balance,
				dvc.tx_hash, dvc.block_height, dvc.log_index, dvc.ts
			FROM delegate_votes_changed dvc
			JOIN daos d ON d.id = dvc.dao_id
			ORDER BY dvc.ts`,
		},
		// --- Governance parameter changes ---
		{
			Filename: "governance_param_changes.csv",
			Query: `SELECT gpc.id, d.symbol AS dao_symbol,
				gpc.chain_id, gpc.contract_address,
				gpc.param_name, gpc.old_value, gpc.new_value,
				gpc.tx_hash, gpc.block_height, gpc.log_index, gpc.ts
			FROM governance_param_changes gpc
			JOIN daos d ON d.id = gpc.dao_id
			ORDER BY gpc.ts`,
		},
		// --- Prices ---
		{
			Filename: "token_prices.csv",
			Query: `SELECT id, token_symbol, price_usd, market_cap_usd, volume_24h_usd,
				price_date, source, fetched_at
			FROM token_prices
			ORDER BY token_symbol, price_date`,
		},
		// --- Treasury ---
		{
			Filename: "dao_treasury_addresses.csv",
			Query: `SELECT dta.id, d.symbol AS dao_symbol,
				a.address AS treasury_address, dta.label,
				b.chain_id, b.name AS chain_name
			FROM dao_treasury_addresses dta
			JOIN daos d ON d.id = dta.dao_id
			JOIN addresses a ON a.id = dta.treasury_address_id
			JOIN blockchains b ON b.id = dta.blockchain_id
			ORDER BY d.symbol`,
		},
		// --- Raw events ---
		{
			Filename: "raw_events.csv",
			Query: `SELECT re.id, d.symbol AS dao_symbol,
				b.chain_id, a.address AS contract_address,
				re.tx_hash, re.block_height, re.log_index, re.ts,
				re.topics, re.data
			FROM raw_events re
			JOIN daos d ON d.id = re.dao_id
			JOIN blockchains b ON b.id = re.chain_id
			JOIN addresses a ON a.id = re.contract_address_id
			ORDER BY re.block_height, re.log_index`,
		},
		// --- Decoded events (ABI-decoded raw events) ---
		{
			Filename: "decoded_events.csv",
			Query: `SELECT de.id, d.symbol AS dao_symbol,
				de.chain_id, de.contract_address,
				de.event_name, de.event_signature,
				de.decoded_params,
				de.tx_hash, de.block_height, de.log_index, de.ts
			FROM decoded_events de
			JOIN daos d ON d.id = de.dao_id
			ORDER BY de.block_height, de.log_index`,
		},
		// --- Materialized views (pre-computed KPI summaries) ---
		{
			Filename: "mv_governance_participation.csv",
			Query:    `SELECT * FROM mv_governance_participation ORDER BY dao_id`,
		},
		{
			Filename: "mv_proposal_voting_stats.csv",
			Query:    `SELECT * FROM mv_proposal_voting_stats ORDER BY dao_id, creation_ts`,
		},
		{
			Filename: "mv_token_transfer_stats.csv",
			Query:    `SELECT * FROM mv_token_transfer_stats ORDER BY dao_id`,
		},
		{
			Filename: "mv_delegation_stats.csv",
			Query:    `SELECT * FROM mv_delegation_stats ORDER BY dao_id`,
		},
		{
			Filename: "mv_voting_power_latest.csv",
			Query:    `SELECT * FROM mv_voting_power_latest ORDER BY dao_id, current_voting_power DESC`,
		},
		{
			Filename: "mv_dao_overview.csv",
			Query:    `SELECT * FROM mv_dao_overview ORDER BY dao_id`,
		},
	}
}

func newCSVCmd(rootCmd *cobra.Command) *cobra.Command {
	var dbURL string
	var outputDir string
	var tables string
	var refreshViews bool

	cmd := &cobra.Command{
		Use:   "csv",
		Short: "Export all indexed data to CSV files",
		Long: `Exports all base tables and materialized views to CSV files.
Each table/view becomes a separate CSV file with headers.
Joins are resolved so output contains human-readable values (DAO symbols, addresses) instead of internal IDs.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCSVExport(cmd.Context(), dbURL, outputDir, tables, refreshViews)
		},
	}

	cmd.Flags().StringVar(&dbURL, "db", "postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable", "Database URL")
	cmd.Flags().StringVar(&outputDir, "output", "./dataset", "Output directory for CSV files")
	cmd.Flags().StringVar(&tables, "tables", "", "Comma-separated list of filenames to export (empty = all)")
	cmd.Flags().BoolVar(&refreshViews, "refresh", true, "Refresh materialized views before exporting")

	return cmd
}

func runCSVExport(ctx context.Context, dbURL, outputDir, tablesFilter string, refreshViews bool) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	// Refresh materialized views before export for fresh data.
	if refreshViews {
		fmt.Println("Refreshing materialized views...")
		if _, err := db.ExecContext(ctx, "SELECT refresh_analysis_views()"); err != nil {
			fmt.Printf("WARNING: failed to refresh views: %v (continuing with stale data)\n", err)
		} else {
			fmt.Println("Views refreshed.")
		}
	}

	// Create output directory.
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	exports := allExports()

	// Filter if requested.
	if tablesFilter != "" {
		filterSet := make(map[string]bool)
		for _, t := range strings.Split(tablesFilter, ",") {
			name := strings.TrimSpace(t)
			if !strings.HasSuffix(name, ".csv") {
				name += ".csv"
			}
			filterSet[name] = true
		}
		var filtered []tableExport
		for _, e := range exports {
			if filterSet[e.Filename] {
				filtered = append(filtered, e)
			}
		}
		if len(filtered) == 0 {
			return fmt.Errorf("no matching tables found for filter: %s", tablesFilter)
		}
		exports = filtered
	}

	start := time.Now()
	totalRows := 0
	errors := 0

	for _, exp := range exports {
		path := filepath.Join(outputDir, exp.Filename)
		fmt.Printf("Exporting %s... ", exp.Filename)

		n, err := exportQueryToCSV(ctx, db, exp.Query, path)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			errors++
			continue
		}

		fmt.Printf("%d rows\n", n)
		totalRows += n
	}

	fmt.Printf("\nExport complete: %d files, %d total rows in %s\n",
		len(exports)-errors, totalRows, time.Since(start).Round(time.Millisecond))

	if errors > 0 {
		return fmt.Errorf("%d exports failed", errors)
	}

	fmt.Printf("Output: %s\n", outputDir)
	return nil
}

// exportQueryToCSV runs a query and writes all rows to a CSV file.
func exportQueryToCSV(ctx context.Context, db *sql.DB, query, path string) (int, error) {
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return 0, fmt.Errorf("columns: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	// Write header.
	if err := w.Write(cols); err != nil {
		return 0, fmt.Errorf("write header: %w", err)
	}

	// Prepare scan destinations.
	vals := make([]interface{}, len(cols))
	ptrs := make([]interface{}, len(cols))
	for i := range vals {
		ptrs[i] = &vals[i]
	}

	count := 0
	for rows.Next() {
		if err := rows.Scan(ptrs...); err != nil {
			return count, fmt.Errorf("scan row %d: %w", count, err)
		}

		record := make([]string, len(cols))
		for i, v := range vals {
			record[i] = formatValue(v)
		}
		if err := w.Write(record); err != nil {
			return count, fmt.Errorf("write row %d: %w", count, err)
		}
		count++
	}

	return count, rows.Err()
}

// formatValue converts a database value to its CSV string representation.
func formatValue(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case []byte:
		return string(val)
	case time.Time:
		if val.IsZero() {
			return ""
		}
		return val.UTC().Format(time.RFC3339)
	case int64:
		return fmt.Sprintf("%d", val)
	case float64:
		return fmt.Sprintf("%g", val)
	case bool:
		if val {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%v", val)
	}
}
