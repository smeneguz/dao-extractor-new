-- This file contains the tables that contain data related to the blockchains
-- like chain information and user/contract addresses.

-- Blockchains table contains all the blockchains that have been indexed.
-- This help reduce the amount of data that we store by avoiding storing 
-- the same blockchain multiple times.
CREATE TABLE blockchains (
    id     BIGSERIAL PRIMARY KEY,
    chain_id VARCHAR(32) NOT NULL UNIQUE,  -- e.g. '1', 'cosmoshub-4'
    name   VARCHAR(64) NOT NULL UNIQUE,   -- e.g. 'Ethereum', 'Binance Smart Chain'
    type   VARCHAR(32) NOT NULL          -- e.g. 'EVM', 'Cosmos'
);
-- To quickly find blockchains by chain ID
CREATE INDEX idx_blockchains_chain_id ON blockchains (chain_id);

-- Defines the encoding type of an address
-- hex: address is encoded in hexadecimal
-- bech32: address is encoded in bech32
CREATE TYPE ADDRESS_ENCODING AS ENUM ('hex', 'bech32');

-- Addresses table contains all the addresses that have been indexed.
-- This help reduce the amount of data that we store by avoiding storing 
-- the same address multiple times and provides a way to label addresses.
CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY,
    address VARCHAR(64) NOT NULL UNIQUE,
    label VARCHAR(128),
    is_contract BOOLEAN NOT NULL,
    encoding ADDRESS_ENCODING NOT NULL
);
-- To quickly find addresses by address
CREATE INDEX idx_addresses_address ON addresses (address);
-- To quickly find contract addresses
CREATE INDEX idx_addresses_is_contract ON addresses (is_contract);

-- Table that associates a blockchain with an address.
-- This is used to track where an address has been used.
CREATE TABLE address_blockchains (
    address_id BIGINT NOT NULL REFERENCES addresses(id),
    blockchain_id BIGINT NOT NULL REFERENCES blockchains(id),
    PRIMARY KEY (address_id, blockchain_id)
);
CREATE INDEX idx_address_blockchains_address_id ON address_blockchains (address_id);
CREATE INDEX idx_address_blockchains_blockchain_id ON address_blockchains (blockchain_id);

