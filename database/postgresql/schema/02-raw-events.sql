-- Raw events table stores uninterpreted EVM event logs for DAOs that lack
-- a dedicated governance module.  These can be decoded offline later.
CREATE TABLE raw_events (
    id BIGSERIAL PRIMARY KEY,

    -- The DAO and chain this event belongs to.
    dao_id   BIGINT NOT NULL REFERENCES daos(id),
    chain_id BIGINT NOT NULL REFERENCES blockchains(id),

    -- The contract that emitted the event.
    contract_address_id BIGINT NOT NULL REFERENCES addresses(id),

    -- Transaction and block context.
    tx_hash      TEXT   NOT NULL,
    block_height BIGINT NOT NULL,
    log_index    INTEGER NOT NULL,
    ts           TIMESTAMP WITHOUT TIME ZONE NOT NULL,

    -- Raw event payload.
    topics JSONB  NOT NULL DEFAULT '[]',
    data   TEXT   NOT NULL DEFAULT '',

    -- Deduplicate: same tx + same log index = same event.
    CONSTRAINT unique_raw_event UNIQUE (tx_hash, log_index)
);

-- To quickly find events for a specific DAO.
CREATE INDEX idx_raw_events_dao_id ON raw_events(dao_id);
-- To quickly scan events by block range.
CREATE INDEX idx_raw_events_block_height ON raw_events(block_height);
-- To quickly find events by contract.
CREATE INDEX idx_raw_events_contract ON raw_events(contract_address_id);
-- To filter events by topic signature (first topic = event selector).
CREATE INDEX idx_raw_events_topics ON raw_events USING GIN (topics);
