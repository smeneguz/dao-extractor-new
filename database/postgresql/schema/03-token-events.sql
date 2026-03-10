-- Token transfer events from governance token contracts.
-- Captures ERC20 Transfer(address indexed from, address indexed to, uint256 value).
CREATE TABLE IF NOT EXISTS token_transfers (
    id BIGSERIAL PRIMARY KEY,
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    chain_id TEXT NOT NULL,
    token_address TEXT NOT NULL,
    from_address TEXT NOT NULL,
    to_address TEXT NOT NULL,
    amount NUMERIC NOT NULL,
    tx_hash TEXT NOT NULL,
    block_height BIGINT NOT NULL,
    log_index INT NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_token_transfer UNIQUE (tx_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_token_transfers_dao ON token_transfers(dao_id);
CREATE INDEX IF NOT EXISTS idx_token_transfers_from ON token_transfers(from_address);
CREATE INDEX IF NOT EXISTS idx_token_transfers_to ON token_transfers(to_address);
CREATE INDEX IF NOT EXISTS idx_token_transfers_height ON token_transfers(block_height);
CREATE INDEX IF NOT EXISTS idx_token_transfers_ts ON token_transfers(ts);
CREATE INDEX IF NOT EXISTS idx_token_transfers_chain_id ON token_transfers(chain_id);

-- Delegation change events from governance token contracts.
-- Captures DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate).
CREATE TABLE IF NOT EXISTS delegation_events (
    id BIGSERIAL PRIMARY KEY,
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    chain_id TEXT NOT NULL,
    token_address TEXT NOT NULL,
    delegator TEXT NOT NULL,
    from_delegate TEXT NOT NULL,
    to_delegate TEXT NOT NULL,
    tx_hash TEXT NOT NULL,
    block_height BIGINT NOT NULL,
    log_index INT NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_delegation_event UNIQUE (tx_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_delegation_events_dao ON delegation_events(dao_id);
CREATE INDEX IF NOT EXISTS idx_delegation_events_delegator ON delegation_events(delegator);
CREATE INDEX IF NOT EXISTS idx_delegation_events_to_delegate ON delegation_events(to_delegate);
CREATE INDEX IF NOT EXISTS idx_delegation_events_height ON delegation_events(block_height);
CREATE INDEX IF NOT EXISTS idx_delegation_events_ts ON delegation_events(ts);
CREATE INDEX IF NOT EXISTS idx_delegation_events_chain_id ON delegation_events(chain_id);

-- Voting power change events from governance token contracts.
-- Captures DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance).
CREATE TABLE IF NOT EXISTS delegate_votes_changed (
    id BIGSERIAL PRIMARY KEY,
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    chain_id TEXT NOT NULL,
    token_address TEXT NOT NULL,
    delegate TEXT NOT NULL,
    previous_balance NUMERIC NOT NULL,
    new_balance NUMERIC NOT NULL,
    tx_hash TEXT NOT NULL,
    block_height BIGINT NOT NULL,
    log_index INT NOT NULL,
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT unique_delegate_votes_changed UNIQUE (tx_hash, log_index)
);
CREATE INDEX IF NOT EXISTS idx_delegate_votes_changed_dao ON delegate_votes_changed(dao_id);
CREATE INDEX IF NOT EXISTS idx_delegate_votes_changed_delegate ON delegate_votes_changed(delegate);
CREATE INDEX IF NOT EXISTS idx_delegate_votes_changed_height ON delegate_votes_changed(block_height);
CREATE INDEX IF NOT EXISTS idx_delegate_votes_changed_ts ON delegate_votes_changed(ts);
CREATE INDEX IF NOT EXISTS idx_delegate_votes_changed_chain_id ON delegate_votes_changed(chain_id);
