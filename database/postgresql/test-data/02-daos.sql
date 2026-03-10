INSERT INTO daos (symbol, name) VALUES
    ('UNI', 'Uniswap DAO')
ON CONFLICT (symbol) DO NOTHING;

INSERT INTO dao_contracts (dao_id, contract_address_id, blockchain_id)
SELECT
    d.id AS dao_id,
    a.id AS contract_address_id,
    b.id AS blockchain_id
FROM
    daos AS d,
    addresses AS a,
    blockchains AS b
WHERE
    d.symbol = 'UNI'
    AND a.address = '0x408ed6354d4973f66138c91495f2f2fcbd8724c3'
    AND b.chain_id = '1'
ON CONFLICT (dao_id, contract_address_id, blockchain_id) DO NOTHING;

INSERT INTO dao_treasury_tokens (dao_id, token_id)
SELECT
    d.id AS dao_id,
    t.id AS token_id
FROM
    daos AS d,
    tokens AS t
WHERE
    d.symbol = 'UNI' -- Target DAO
    AND (
        t.symbol = 'ETH' -- Link Uniswap DAO to ETH
        OR t.symbol = 'USDT' -- Link Uniswap DAO to USDT
    )
ON CONFLICT (dao_id, token_id) DO NOTHING; -- Assuming this is the unique constraint

INSERT INTO dao_treasury_addresses (dao_id, treasury_address_id, blockchain_id, used_from, used_until)
SELECT
    d.id AS dao_id,
    a.id AS treasury_address_id,
    b.id AS blockchain_id,
    CASE
        WHEN a.address = '0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97' AND b.chain_id = '1' THEN '2022-01-01 00:00:00'::timestamp
        WHEN a.address = '0xFF38B106FCe9647Bdf1E7877BF73cE8B0BAD5f43' AND b.chain_id = '1' THEN '2022-01-01 00:00:00'::timestamp
        ELSE NULL
    END AS used_from,
    CASE
        WHEN a.address = '0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97' AND b.chain_id = '1' THEN '2022-12-31 23:59:59'::timestamp
        WHEN a.address = '0xFF38B106FCe9647Bdf1E7877BF73cE8B0BAD5f43' AND b.chain_id = '1' THEN NULL
        ELSE NULL
    END AS used_until
FROM
    daos AS d,
    addresses AS a,
    blockchains AS b
WHERE
    d.symbol = 'UNI' -- Target DAO
    AND b.name = 'Ethereum' -- All treasury addresses are on Ethereum in this example
    AND (
        a.address = '0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97' -- Uniswap Treasury 1
        OR a.address = '0xFF38B106FCe9647Bdf1E7877BF73cE8B0BAD5f43' -- Uniswap Treasury 2
    )
ON CONFLICT DO NOTHING;

