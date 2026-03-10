-- Table used to track the tokens balances of users
CREATE TABLE user_token_balances (
    -- Token identifier, will be the token denom for native tokens
    -- and the contract address for non-native tokens
    denom String,
    -- Address of the user holding the token
    user_address String,
    -- Blockchain id inside the postgresql database
    chain_id UInt64,
    -- The user's balance of the token
    balance UInt256,
    -- Block height of when the balance was recorded
    block_height UInt64,
    -- Timestamp of when the balance was recorded
    timestamp DateTime,
)
ENGINE = MergeTree
ORDER BY (user_address, denom, chain_id, block_height, timestamp);

