INSERT INTO tokens (symbol, name) VALUES
    ('ETH', 'Ether'),
    ('USDT', 'Tether USD'),
    ('ATOM', 'Cosmos ATOM')
ON CONFLICT (symbol) DO NOTHING;

INSERT INTO token_contracts (token_id, blockchain_id, contract_address_id, standard, decimals)
SELECT
    t.id AS token_id,
    b.id AS blockchain_id,
    a.id AS contract_address_id,
    'erc20' AS standard,
    18 AS decimals
FROM
    tokens AS t,
    blockchains AS b,
    addresses AS a
WHERE
    t.symbol = 'USDT' AND b.chain_id = '1' AND a.address = '0xdAC17F958D2ee523a2206206994597C13D831ec7'
ON CONFLICT (token_id, blockchain_id, contract_address_id) DO NOTHING;

INSERT INTO token_natives (token_id, blockchain_id, denom, decimals)
SELECT
    t.id AS token_id,
    b.id AS blockchain_id,
    CASE
        WHEN t.symbol = 'ETH' THEN 'eth'
        WHEN t.symbol = 'ATOM' THEN 'uatom'
        ELSE NULL
    END AS denom,
    CASE
        WHEN t.symbol = 'ETH'  THEN 18
        WHEN t.symbol = 'ATOM' THEN 6
        ELSE NULL
    END AS decimals
FROM
    tokens AS t,
    blockchains AS b
WHERE
    (t.symbol = 'ETH' AND b.chain_id = '1')
    OR
    (t.symbol = 'ATOM' AND b.chain_id = 'cosmoshub-4')
ON CONFLICT (token_id, blockchain_id, denom) DO NOTHING;
