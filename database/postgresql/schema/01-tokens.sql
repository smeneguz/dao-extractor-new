-- This file contains the tables that contain data related to the tokens.

-- Table that contains information about the tokens
CREATE TABLE tokens (
    -- A unique internal ID for the token
    id BIGSERIAL PRIMARY KEY,
    -- The symbol of the token (e.g., "UNI")
    symbol VARCHAR(32) NOT NULL UNIQUE,
    -- The name of the token (e.g., "Uniswap")
    name VARCHAR(128) NOT NULL
);
-- To quickly find tokens by symbol
CREATE INDEX idx_tokens_symbol ON tokens (symbol);

-- Defines the type of token
-- erc20: ERC20 token
-- erc721: ERC721 token
-- erc1155: ERC1155 token
CREATE TYPE TOKEN_CONTRACT_STANDARD AS ENUM ('erc20', 'erc721', 'erc1155');

-- Table that contains information about the contracts that represent the tokens
-- and where the contract is deployed on
CREATE TABLE token_contracts (
    -- A unique internal ID for the token
    token_id BIGINT NOT NULL REFERENCES tokens(id),
    -- The blockchain where the contract is deployed on
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(id),
    -- The address of the contract that represents the token
    contract_address_id BIGINT NOT NULL REFERENCES addresses(id),
    -- The token standard.
    standard TOKEN_CONTRACT_STANDARD NOT NULL,
    -- The number of decimals used by the token
    decimals SMALLINT NOT NULL,
    PRIMARY KEY (token_id, contract_address_id, blockchain_id)
);
-- To quickly find the contracts that represent a token
CREATE INDEX idx_token_contracts_token_id ON token_contracts (token_id);
-- To quickly find the contracts that represent a token
CREATE INDEX idx_token_contracts_contract_address_id ON token_contracts (contract_address_id);

-- Table that contains information about the native tokens
CREATE TABLE token_natives(
    -- A unique internal ID for the token
    token_id BIGINT NOT NULL REFERENCES tokens(id),
    -- The blockchain where the token is available on
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(id),
    -- The token denomination, e.g. uatom (for Cosmos), or wei (for Ethereum)
    denom varchar(128) NOT NULL,
    -- The number of decimals used by the token
    decimals SMALLINT NOT NULL,
    PRIMARY KEY (token_id, blockchain_id, denom)
);
-- To quickly find the native tokens
CREATE INDEX idx_token_natives_token_id ON token_natives (token_id);
-- To quickly find native token information by its blockchain
CREATE INDEX idx_token_natives_blockchain_id ON token_natives (blockchain_id);
-- To quickly find native token information by its denomination
CREATE INDEX idx_token_natives_denom ON token_natives (denom);

-- View that contains the unified tokens representation
CREATE OR REPLACE VIEW tokens_unified_view AS
SELECT
    t.id                               AS token_id,
    t.symbol                           AS symbol,
    t.name                             AS name,
    'contract'::varchar(16)            AS token_type,
    tc.blockchain_id                   AS blockchain_id,
    b.chain_id                         AS chain_id,
    b.name                             AS chain_name,
    b.type                             AS chain_type,
    tc.contract_address_id             AS contract_address_id,
    a.address                          AS contract_address,
    a.encoding                         AS contract_address_encoding,
    tc.standard                        AS contract_standard,
    NULL::varchar(128)                 AS denom,
    tc.decimals                        AS decimals
FROM token_contracts tc
JOIN tokens t ON t.id = tc.token_id
JOIN addresses a ON a.id = tc.contract_address_id
JOIN blockchains b ON b.id = tc.blockchain_id

UNION ALL

SELECT
    t.id                               AS token_id,
    t.symbol                           AS symbol,
    t.name                             AS name,
    'native'::varchar(16)              AS token_type,
    tn.blockchain_id                   AS blockchain_id,
    b.chain_id                         AS chain_id,
    b.name                             AS chain_name,
    b.type                             AS chain_type,
    NULL                               AS contract_address_id,
    NULL                               AS contract_address,
    NULL::ADDRESS_ENCODING             AS contract_address_encoding,
    NULL::TOKEN_CONTRACT_STANDARD      AS contract_standard,
    tn.denom                           AS denom,
    tn.decimals                        AS decimals
FROM token_natives tn
JOIN tokens t ON t.id = tn.token_id
JOIN blockchains b ON b.id = tn.blockchain_id;

