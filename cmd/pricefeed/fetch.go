package pricefeed

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// CoinGecko ID mapping for our governance tokens.
// COMP-BRAVO and COMP-OZ share the same COMP token.
// UDT and UP share the same Unlock Protocol token (migrated from Ethereum to Base).
// NOUNS-PUB uses NFT-based governance (no fungible token on CoinGecko).
var tokenCoinGeckoIDs = map[string]string{
	"UNI":        "uniswap",
	"COMP":       "compound-governance-token",
	"COMP-BRAVO": "compound-governance-token",
	"COMP-OZ":    "compound-governance-token",
	"ENS":        "ethereum-name-service",
	"AAVE":       "aave",
	"SKY":        "sky",
	"CRV":        "curve-dao-token",
	"LDO":        "lido-dao",
	"DEXE":       "dexe",
	"W":          "wormhole",
	"DIVA":       "diva-staking",
	"AGLD":       "adventure-gold",
	"PAPER":      "dope-wars-paper",
	"HOP":        "hop-protocol",
	"UDT":        "unlock-protocol",
	"UP":         "unlock-protocol",
	"FLT":        "fluence-token",
	"GTC":        "gitcoin",
	"RAD":        "radicle",
	"UNION":      "union-finance",
	"T":          "threshold-network-token",
	"TRU":        "truefi",
	"POOL":       "pooltogether",
	"ATF":        "antfarm-governance-token",
	"INV":        "inverse-finance",
	"HIFI":       "hifi-finance",
	"FORTH":      "ampleforth-governance-token",
	"IDLE":       "idle",
	"CTX":        "cryptex-finance",
	"ANVIL":      "anvil",
	"MPL":        "maple",
	"BTRST":      "braintrust",
	"UMA":        "uma",
	"POOH":       "pooh",
	"RARI":       "rarible",
	"TORN":       "tornado-cash",
	"KNC":        "kyber-network-crystal",
	"AUDIO":      "audius",
	"GRG":        "rigoblock",
	"ARB":        "arbitrum",
	"OCA":        "onchainaustria",
	"GMX":        "gmx",
	"OD":         "open-dollar-governance",
	"HAI":        "hai",
	"SEAM":       "seamless-protocol",
	"SUMMER":     "summer-fi",
	"FAME":       "fame-lady-society",
	"REG":        "realtoken-ecosystem-governance",
	"LSK":        "lisk",
	"DEGEN-DOGS": "degen-dogs",
}

func newFetchCmd(rootCmd *cobra.Command) *cobra.Command {
	var dbURL string
	var daysBack int
	var symbols string

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch historical token prices from CoinGecko",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runFetch(cmd.Context(), dbURL, daysBack, symbols)
		},
	}

	cmd.Flags().StringVar(&dbURL, "db", "postgres://dao_portal:password@localhost:5432/dao_portal?sslmode=disable", "Database URL")
	cmd.Flags().IntVar(&daysBack, "days", 365, "Number of days of history to fetch")
	cmd.Flags().StringVar(&symbols, "symbols", "", "Comma-separated token symbols to fetch (empty = all)")

	return cmd
}

func runFetch(ctx context.Context, dbURL string, daysBack int, symbolsFilter string) error {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	targets := tokenCoinGeckoIDs
	if symbolsFilter != "" {
		targets = make(map[string]string)
		for _, s := range strings.Split(symbolsFilter, ",") {
			s = strings.TrimSpace(strings.ToUpper(s))
			if cgID, ok := tokenCoinGeckoIDs[s]; ok {
				targets[s] = cgID
			} else {
				fmt.Printf("WARNING: unknown symbol %s, skipping\n", s)
			}
		}
	}

	// Deduplicate CoinGecko IDs to avoid fetching the same data multiple times.
	// Multiple DAO symbols may share the same token (e.g., COMP, COMP-BRAVO, COMP-OZ).
	type fetchResult struct {
		points []pricePoint
		err    error
	}
	cgCache := make(map[string]*fetchResult)

	client := &http.Client{Timeout: 30 * time.Second}
	totalErrors := 0
	backoff := 7 * time.Second

	for symbol, cgID := range targets {
		// Check if we already fetched this CoinGecko ID
		cached, haveCached := cgCache[cgID]
		if !haveCached {
			fmt.Printf("Fetching %s (%s)...\n", symbol, cgID)

			points, fetchErr := fetchCoinGeckoHistory(ctx, client, cgID, daysBack)
			cached = &fetchResult{points: points, err: fetchErr}
			cgCache[cgID] = cached

			if fetchErr != nil {
				fmt.Printf("  ERROR: %v (skipping)\n", fetchErr)
				totalErrors++
				// Exponential backoff on rate limit
				if strings.Contains(fetchErr.Error(), "429") {
					backoff = min(backoff*2, 60*time.Second)
				}
				time.Sleep(backoff)
				continue
			}

			// Reset backoff on success
			backoff = 7 * time.Second
			// CoinGecko free tier: ~10 req/min
			time.Sleep(backoff)
		} else if cached.err != nil {
			fmt.Printf("Skipping %s (%s): previous fetch failed\n", symbol, cgID)
			continue
		} else {
			fmt.Printf("Storing %s (%s) from cache...\n", symbol, cgID)
		}

		inserted := 0
		dbErrors := 0
		for _, p := range cached.points {
			_, err := db.ExecContext(ctx, `
				INSERT INTO token_prices(token_symbol, price_usd, market_cap_usd, volume_24h_usd, price_date, source)
				VALUES ($1, $2, $3, $4, $5, 'coingecko')
				ON CONFLICT (token_symbol, price_date, source) DO UPDATE SET
					price_usd = EXCLUDED.price_usd,
					market_cap_usd = EXCLUDED.market_cap_usd,
					volume_24h_usd = EXCLUDED.volume_24h_usd,
					fetched_at = NOW()
			`, symbol, p.Price, p.MarketCap, p.Volume, p.Date)
			if err != nil {
				fmt.Printf("  DB error for %s on %s: %v\n", symbol, p.Date, err)
				dbErrors++
				continue
			}
			inserted++
		}

		if dbErrors > 0 {
			fmt.Printf("  %d stored, %d DB errors\n", inserted, dbErrors)
			totalErrors += dbErrors
		} else {
			fmt.Printf("  %d price points stored\n", inserted)
		}
	}

	if totalErrors > 0 {
		return fmt.Errorf("completed with %d errors", totalErrors)
	}

	fmt.Println("Done.")
	return nil
}

type pricePoint struct {
	Date      string
	Price     float64
	MarketCap float64
	Volume    float64
}

func fetchCoinGeckoHistory(ctx context.Context, client *http.Client, coinID string, days int) ([]pricePoint, error) {
	url := fmt.Sprintf(
		"https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=usd&days=%d&interval=daily",
		coinID, days,
	)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return nil, fmt.Errorf("rate limited (429)")
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	var result struct {
		Prices       [][]float64 `json:"prices"`
		MarketCaps   [][]float64 `json:"market_caps"`
		TotalVolumes [][]float64 `json:"total_volumes"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	points := make([]pricePoint, 0, len(result.Prices))
	for i, p := range result.Prices {
		if len(p) < 2 {
			continue
		}
		ts := time.UnixMilli(int64(p[0]))
		dateStr := ts.UTC().Format("2006-01-02")

		var mcap float64
		if i < len(result.MarketCaps) && len(result.MarketCaps[i]) >= 2 {
			mcap = result.MarketCaps[i][1]
		}

		var vol float64
		if i < len(result.TotalVolumes) && len(result.TotalVolumes[i]) >= 2 {
			vol = result.TotalVolumes[i][1]
		}

		points = append(points, pricePoint{
			Date:      dateStr,
			Price:     p[1],
			MarketCap: mcap,
			Volume:    vol,
		})
	}

	return points, nil
}
