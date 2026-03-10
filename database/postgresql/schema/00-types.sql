-- This file contains the types definitions that are used in the database.

-- Type that defines a coin
CREATE TYPE COIN AS (
    -- The coin denom
    denom VARCHAR(64),
    -- The coin amount
    amount NUMERIC
);
