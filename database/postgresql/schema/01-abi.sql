-- Stores the ABI of the contracts
CREATE TABLE abis
(
    -- Address of the contract
    contract_address TEXT NOT NULL,
    -- ID of the chain where the contract has been deployed
    chain_id TEXT NOT NULL,
    -- ABI of the contract
    abi JSONB NOT NULL,
    CONSTRAINT unique_abi UNIQUE (contract_address, chain_id)
);
