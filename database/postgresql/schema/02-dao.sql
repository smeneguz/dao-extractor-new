-- This file contains the tables for the DAOs information, their governance contracts and treasuries.

-- This table contains all the DAOs that have been indexed.
CREATE TABLE daos (
    -- A unique internal ID for the DAO
    id BIGSERIAL PRIMARY KEY,
    -- The symbol of the DAO (e.g., "UNI")
    symbol VARCHAR(32) NOT NULL UNIQUE,
    -- The name of the DAO (e.g., "Uniswap DAO")
    name VARCHAR(128) NOT NULL
);
CREATE INDEX idx_daos_symbol ON daos (symbol);

-- This table contains the association between a DAO and the contracts 
-- that are being used to govern it.
CREATE TABLE dao_contracts (
    -- A unique internal ID for the DAO
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    -- The address of the contract that governs the DAO
    contract_address_id BIGINT NOT NULL REFERENCES addresses(id),
    -- The blockchain where the contract is deployed on
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(id),
    PRIMARY KEY (dao_id, contract_address_id, blockchain_id)
);
CREATE INDEX idx_dao_contracts_dao_id ON dao_contracts (dao_id);
CREATE INDEX idx_dao_contracts_contract_address_id ON dao_contracts (contract_address_id);
CREATE INDEX idx_dao_contracts_blockchain_id ON dao_contracts (blockchain_id);

-- Table that contains the association between a DAO and the tokens in its treasury
CREATE TABLE dao_treasury_tokens (
    -- A unique internal ID for the DAO
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    -- The token that is being used to govern the DAO
    token_id BIGINT NOT NULL REFERENCES tokens(id),
    PRIMARY KEY (dao_id, token_id)
);
-- To quickly find the tokens that are being used to govern a DAO
CREATE INDEX idx_dao_treasury_tokens_dao_id ON dao_treasury_tokens (dao_id);

-- Table that contains the association between a DAO and its treasury addresses
CREATE TABLE dao_treasury_addresses (
    id BIGSERIAL PRIMARY KEY,
    -- A unique internal ID for the DAO
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    -- The address of the treasury
    treasury_address_id BIGINT NOT NULL REFERENCES addresses(id),
    -- The blockchain where the treasury is deployed on
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(id),
    -- Type of treasury address: timelock, multisig, treasury, executor, agent, vault
    label TEXT NOT NULL DEFAULT 'treasury',
    -- Timestamp from when the address was used as a treasury for the DAO
    used_from TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    -- Timestamp until when the address was used as a treasury for the DAO
    -- if null, the address is still being used
    used_until TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT unique_treasury_entry UNIQUE (dao_id, treasury_address_id, blockchain_id, label)
);
-- To quickly find the treasury addresses of a DAO
CREATE INDEX idx_dao_treasury_dao_id ON dao_treasury_addresses (dao_id);
-- To quickly find the current active treasury addresses of a DAO
CREATE INDEX idx_dao_treasury_active_treasuries ON dao_treasury_addresses (dao_id, treasury_address_id)
WHERE (used_until IS NULL);
