-- Governance parameter change events from Governor Bravo contracts.
-- Captures VotingDelaySet, VotingPeriodSet, ProposalThresholdSet events.
-- These track the evolution of governance rules over time.
CREATE TABLE IF NOT EXISTS governance_param_changes (
    id BIGSERIAL PRIMARY KEY,
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    chain_id TEXT NOT NULL,
    contract_address TEXT NOT NULL,
    param_name TEXT NOT NULL,        -- 'voting_delay', 'voting_period', 'proposal_threshold'
    old_value NUMERIC NOT NULL,
    new_value NUMERIC NOT NULL,
    tx_hash TEXT NOT NULL,
    block_height BIGINT NOT NULL,
    log_index INT NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_governance_param_change UNIQUE (tx_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_gov_param_changes_dao ON governance_param_changes(dao_id);
CREATE INDEX IF NOT EXISTS idx_gov_param_changes_param ON governance_param_changes(param_name);
CREATE INDEX IF NOT EXISTS idx_gov_param_changes_height ON governance_param_changes(block_height);
CREATE INDEX IF NOT EXISTS idx_gov_param_changes_ts ON governance_param_changes(ts);
