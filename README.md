# Extractor

Blockchain indexer for DAO governance data. Built on the Flux framework, it indexes proposals, votes, token events, and raw contract logs from 50+ DAOs across 7 EVM chains into PostgreSQL.

## Architecture

```
                          50 DAOs x 7 Chains
                                 |
                    Flux Framework (block workers)
                                 |
          +----------+-----------+-----------+
          |          |           |           |
   Gov Bravo    OZ Governor  Raw Events  Token Events
    (8 DAOs)    (21 DAOs)    (16 DAOs)   (all DAOs)
          |          |           |           |
          v          v           v           v
    proposals    proposals   raw_events  token_transfers
    vote_actions vote_actions (topics+   delegation_events
    finalizations finaliz.    data)      delegate_votes_changed
    param_changes                |
                                 | abi fetch-all + decode raw
                                 v
                          decoded_events
                          (JSONB params)

    + token_prices (CoinGecko)
    + dao_treasury_addresses (78 entries)
    + 6 materialized views (KPI summaries)
```

## Prerequisites

- Go 1.24+
- PostgreSQL 16+
- `psql` CLI

## Quick start

```bash
# 1. Database
createuser dao_portal -P           # password: your_password
createdb dao_portal -O dao_portal
psql dao_portal -c "GRANT ALL ON SCHEMA public TO dao_portal;"

# 2. Schema
chmod +x scripts/*.sh
./scripts/init-db.sh "postgres://dao_portal:your_password@localhost:5432/dao_portal?sslmode=disable"

# 3. Config
# Edit config.yaml: set DB URL, RPC endpoints, Etherscan API key
# Flux does NOT support ${VAR} interpolation - hardcode all values

# 4. Build
go build -o build/extractor ./cmd/extractor

# 5. Run
./build/extractor start --home $PWD
```

## CLI commands

| Command | Description |
|---------|-------------|
| `start` | Run all indexers (real-time block processing) |
| `parse range <indexer> <from> [to]` | Backfill historical blocks |
| `price-feed fetch` | Fetch token prices from CoinGecko |
| `abi download <address>` | Download single contract ABI from Etherscan |
| `abi fetch-all --api-key KEY` | Bulk download ABIs for all raw_events contracts |
| `decode raw` | Decode raw_events using stored ABIs |
| `export csv` | Export all tables to CSV files |

### Price feed

```bash
# Full history (365 days, all tokens)
./build/extractor price-feed fetch --db "postgres://..." --days 365

# Specific tokens only
./build/extractor price-feed fetch --db "postgres://..." --symbols "UNI,AAVE,ENS"
```

### ABI download and decode

```bash
# Bulk download ABIs for all contracts in raw_events
./build/extractor abi fetch-all --api-key YOUR_ETHERSCAN_KEY --db "postgres://..."

# Decode raw events into structured decoded_events table
./build/extractor decode raw --db "postgres://..."

# Decode specific DAO only
./build/extractor decode raw --db "postgres://..." --dao "AAVE,CRV"
```

### Export

```bash
# Export all tables + materialized views to CSV
./build/extractor export csv --output ./dataset/ --db "postgres://..."

# Export specific tables
./build/extractor export csv --tables "proposals,vote_actions,token_transfers" --output ./dataset/

# Skip materialized view refresh
./build/extractor export csv --refresh=false --output ./dataset/
```

Produces 19 CSV files: `daos.csv`, `proposals.csv`, `vote_actions.csv`, `proposal_finalizations.csv`, `token_transfers.csv`, `delegation_events.csv`, `delegate_votes_changed.csv`, `governance_param_changes.csv`, `token_prices.csv`, `dao_treasury_addresses.csv`, `raw_events.csv`, `decoded_events.csv`, plus 6 `mv_*.csv` materialized view summaries.

## Indexing workflow

### Phase 1: Historical backfill

```bash
# Run each chain in a separate terminal
./scripts/backfill.sh ethereum-mainnet 4620855  24620000
./scripts/backfill.sh arbitrum        70397646 300000000
./scripts/backfill.sh optimism        116055145 130000000
./scripts/backfill.sh base            5842016  25000000
./scripts/backfill.sh gnosis          35476313 38000000
./scripts/backfill.sh lisk            1        10000000
./scripts/backfill.sh polygon         25818525 70000000

# Verify zero gaps
./scripts/fill-gaps.sh ethereum-mainnet
```

### Phase 2: Real-time monitoring

```bash
./build/extractor start --home $PWD
```

Starts all 7 indexers concurrently. Only use `start` after backfill is complete. See "Why two phases?" below.

### Why two phases?

`start` uses `generate_series(start_height, current_height) EXCEPT SELECT height FROM blocks` to find missing blocks. For chains with millions of blocks (Arbitrum: 230M+), this query is too slow. `parse range` bypasses this with a `RangeHeightProducer` that feeds heights directly.

## Database schema

```
00-types.sql              ENUM types (PROPOSAL_STATUS, VOTE_ACTION_TYPE)
00-blocks.sql             blocks table (Flux internal)
00-indexing.sql           indexing operations (Flux internal)
01-abi.sql                abis table (contract ABIs for decoding)
01-blockchain.sql         blockchains, addresses tables
01-tokens.sql             tokens table
02-dao.sql                daos, dao_contracts, dao_treasury_addresses
02-governance.sql         proposals, vote_actions, proposal_finalizations
02-raw-events.sql         raw_events (uninterpreted EVM logs)
03-token-events.sql       token_transfers, delegation_events, delegate_votes_changed
04-price-feed.sql         token_prices (CoinGecko daily prices)
05-analysis-views.sql     6 materialized views + refresh function
06-treasury-seed.sql      78 treasury addresses (timelock, multisig, treasury, executor, agent, vault)
07-governance-params.sql  governance_param_changes (VotingDelaySet, VotingPeriodSet, ProposalThresholdSet)
08-decoded-events.sql     decoded_events (ABI-decoded raw events)
```

Refresh materialized views after indexing:

```sql
SELECT refresh_analysis_views();
```

## Modules

| Module | DAOs | Events captured |
|--------|------|-----------------|
| **governor-bravo** | 8 (UNI, COMP-BRAVO, INV, HIFI, FORTH, IDLE, CTX, ANVIL) | ProposalCreated, VoteCast, ProposalQueued, ProposalExecuted, ProposalCanceled, VotingDelaySet, VotingPeriodSet, ProposalThresholdSet |
| **oz-governor** | 21 (ENS, COMP-OZ, W, DIVA, ARB, GMX, ...) | ProposalCreated, VoteCast, VoteCastWithParams, ProposalQueued, ProposalExecuted, ProposalCanceled |
| **raw-events** | 16 (AAVE, SKY, CRV, LDO, DEXE, ...) | All contract logs (raw topics + data) |
| **token-events** | All DAOs with governance tokens | Transfer, DelegateChanged, DelegateVotesChanged |

DAOs in raw-events use custom governance (Aragon, DSChief, Aave V2/V3, NounsDAO forks, etc.). Use `abi fetch-all` + `decode raw` to get structured decoded events.

## Chains and indexers

| Indexer | Chain | Node type | Modules |
|---------|-------|-----------|---------|
| `ethereum-mainnet` | Ethereum (1) | evm-rpc-fallback | governor-bravo, oz-governor, raw-events, token-events |
| `arbitrum` | Arbitrum (42161) | evm-rpc-fallback | oz-governor, raw-events, token-events |
| `optimism` | Optimism (10) | evm-rpc-fallback | oz-governor, token-events |
| `base` | Base (8453) | evm-rpc-fallback | oz-governor, token-events |
| `polygon` | Polygon (137) | evm-rpc-fallback | raw-events, token-events |
| `gnosis` | Gnosis (100) | evm-rpc | oz-governor, token-events |
| `lisk` | Lisk (1135) | evm-rpc | oz-governor, token-events |

## Node types

**evm-rpc**: Single RPC endpoint. Used for chains without Alchemy.

**evm-rpc-fallback**: Primary (PublicNode, fast) + fallback (Alchemy, archive). On failure, switches to fallback with 2s cooldown. On 429 errors, exponential backoff up to 5 retries. Blocks are never dropped.

