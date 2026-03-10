INSERT INTO blockchains (chain_id, name, type) VALUES
    ('1', 'Ethereum', 'EVM'),
    ('3', 'Polygon', 'EVM'),
    ('cosmoshub-4', 'Cosmos Hub', 'Cosmos');

INSERT INTO addresses (address, label, is_contract, encoding) VALUES
    ('0x408ed6354d4973f66138c91495f2f2fcbd8724c3', 'Uniswap V3 Governor', true, 'hex'),
    ('0xdAC17F958D2ee523a2206206994597C13D831ec7', 'Tether USD', true, 'hex'),
    ('0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97', 'Uniswap Treasury', false, 'hex'),
    ('0xFF38B106FCe9647Bdf1E7877BF73cE8B0BAD5f43', 'Uniswap Treasury 2', false, 'hex')
ON CONFLICT (address) DO NOTHING;

INSERT INTO address_blockchains (address_id, blockchain_id)
SELECT
    a.id AS address_id,
    b.id AS blockchain_id
FROM
    addresses AS a,
    blockchains AS b
WHERE
    (a.address = '0x408ed6354d4973f66138c91495f2f2fcbd8724c3' AND b.chain_id = '1') -- Link Uniswap V3 Governor to Ethereum
    OR
    (a.address = '0xdAC17F958D2ee523a2206206994597C13D831ec7' AND b.chain_id = '1') -- Link Tether USD contract to Ethereum
ON CONFLICT (address_id, blockchain_id) DO NOTHING;
