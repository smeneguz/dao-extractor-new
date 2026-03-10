-- Decoded events from raw_events using contract ABIs.
-- Provides structured, queryable governance data for ALL DAOs,
-- including those with custom governance mechanisms.
CREATE TABLE IF NOT EXISTS decoded_events (
    id BIGSERIAL PRIMARY KEY,
    raw_event_id BIGINT REFERENCES raw_events(id),
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    chain_id TEXT NOT NULL,
    contract_address TEXT NOT NULL,
    event_name TEXT NOT NULL,           -- e.g. "ProposalCreated", "VoteCast", "StartVote"
    event_signature TEXT NOT NULL,      -- full sig e.g. "ProposalCreated(uint256,address,...)"
    decoded_params JSONB NOT NULL,      -- all decoded parameters as key-value
    tx_hash TEXT NOT NULL,
    block_height BIGINT NOT NULL,
    log_index INT NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_decoded_event UNIQUE (tx_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_decoded_events_dao ON decoded_events(dao_id);
CREATE INDEX IF NOT EXISTS idx_decoded_events_name ON decoded_events(event_name);
CREATE INDEX IF NOT EXISTS idx_decoded_events_height ON decoded_events(block_height);
CREATE INDEX IF NOT EXISTS idx_decoded_events_ts ON decoded_events(ts);
CREATE INDEX IF NOT EXISTS idx_decoded_events_contract ON decoded_events(contract_address);
CREATE INDEX IF NOT EXISTS idx_decoded_events_params ON decoded_events USING GIN (decoded_params);
