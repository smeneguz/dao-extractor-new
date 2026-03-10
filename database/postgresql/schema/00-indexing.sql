-- Table to store operations that should be executed when we reach a certain height.
CREATE TABLE IF NOT EXISTS height_deferred_operations (
    -- Key used to identify who created the operation
    creator_key VARCHAR(64) NOT NULL,
    -- Operation type
    type VARCHAR(64) NOT NULL,
    -- Height at which the operation should be performed
    height BIGINT NOT NULL,
    -- Operation payload
    payload JSONB,

    UNIQUE (creator_key, type, height)
);

-- Index to quickly find operations by creator key.
CREATE INDEX IF NOT EXISTS height_deferred_operations_creator_key_idx ON height_deferred_operations (creator_key);

