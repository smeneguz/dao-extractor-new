-- ================================================================
-- Treasury Address Seed Data
-- Populates dao_treasury_addresses with all known treasury,
-- timelock, multisig, executor, agent and vault addresses.
--
-- Labels:
--   timelock    = TimelockController / executor with delay
--   multisig    = Gnosis Safe / multi-signature wallet
--   treasury    = Dedicated treasury / vault contract
--   executor    = Proposal executor (may hold funds)
--   agent       = Aragon Agent (acts as DAO treasury)
--   vault       = Specialized asset vault
-- ================================================================

-- Helper: insert an address if it doesn't exist, upsert treasury row.
-- We use a DO block with a function for reuse.

CREATE OR REPLACE FUNCTION seed_treasury(
    p_dao_symbol TEXT,
    p_chain_id TEXT,
    p_address TEXT,
    p_label TEXT,
    p_used_from TIMESTAMP,
    p_used_until TIMESTAMP DEFAULT NULL
) RETURNS void AS $$
DECLARE
    v_dao_id BIGINT;
    v_blockchain_id BIGINT;
    v_address_id BIGINT;
BEGIN
    -- Get DAO id
    SELECT id INTO v_dao_id FROM daos WHERE symbol = p_dao_symbol;
    IF v_dao_id IS NULL THEN
        RAISE NOTICE 'DAO % not found, skipping', p_dao_symbol;
        RETURN;
    END IF;

    -- Get blockchain id
    SELECT id INTO v_blockchain_id FROM blockchains WHERE chain_id = p_chain_id;
    IF v_blockchain_id IS NULL THEN
        RAISE NOTICE 'Blockchain chain_id=% not found, skipping', p_chain_id;
        RETURN;
    END IF;

    -- Upsert address (all treasury addresses are contracts on EVM chains)
    INSERT INTO addresses (address, is_contract, encoding)
    VALUES (LOWER(p_address), true, 'hex')
    ON CONFLICT (address) DO NOTHING;

    SELECT id INTO v_address_id FROM addresses WHERE address = LOWER(p_address);

    -- Upsert treasury entry
    INSERT INTO dao_treasury_addresses (dao_id, treasury_address_id, blockchain_id, label, used_from, used_until)
    VALUES (v_dao_id, v_address_id, v_blockchain_id, p_label, p_used_from, p_used_until)
    ON CONFLICT ON CONSTRAINT unique_treasury_entry DO NOTHING;
END;
$$ LANGUAGE plpgsql;

-- ================================================================
-- 1. UNI (Uniswap) - Ethereum
-- ================================================================
SELECT seed_treasury('UNI', '1', '0x1a9C8182C09F50C8318d769245beA52c32BE35BC', 'timelock', '2020-09-17');

-- ================================================================
-- 2. ARB (Arbitrum) - Arbitrum
-- ================================================================
-- Core timelock (governance proposals execute here)
SELECT seed_treasury('ARB', '42161', '0x34d45e99f7D8c45ed05B5cA72D54bbD1fb3F98f0', 'timelock', '2023-03-16');
-- Treasury-specific timelock
SELECT seed_treasury('ARB', '42161', '0xbFc1FECa8B09A5c5D3EFfE7429eBE24b9c09EF58', 'timelock', '2023-03-16');
-- Treasury wallet (FixedDelegateErc20Wallet)
SELECT seed_treasury('ARB', '42161', '0xF3FC178157fb3c87548bAA86F9d24BA38E649B58', 'treasury', '2023-03-16');

-- ================================================================
-- 3. AAVE - Ethereum
-- ================================================================
-- Current executors
SELECT seed_treasury('AAVE', '1', '0x5300A1a15135EA4dc7aD5a167152C01EFc9b192A', 'executor', '2023-09-28');
SELECT seed_treasury('AAVE', '1', '0x17Dd33Ed0e3dD2a80E37489B8A63063161BE6957', 'executor', '2023-09-28');
-- Old executors (short + long)
SELECT seed_treasury('AAVE', '1', '0xEE56e2B3D491590B5b31738cC34d5232F378a8D5', 'executor', '2020-12-09', '2023-09-28');
SELECT seed_treasury('AAVE', '1', '0x61910EcD7e8e942136CE7Fe7943f956cea1CC2f7', 'executor', '2020-12-09', '2023-09-28');
-- Timelock (PayloadsController)
SELECT seed_treasury('AAVE', '1', '0xdAbad81aF85554E9ae636395611C58F7eC1aAEc5', 'timelock', '2023-09-28');
-- Ecosystem Reserve (treasury)
SELECT seed_treasury('AAVE', '1', '0x25f2226b597e8f9514b3f68f00f494cf4f286491', 'treasury', '2020-10-02');
-- Collector (treasury)
SELECT seed_treasury('AAVE', '1', '0x464C71f6c2F760DdA6093dCB91C24c39e5d6e18c', 'treasury', '2020-11-15');

-- ================================================================
-- 4. SKY (MakerDAO) - Ethereum
-- ================================================================
SELECT seed_treasury('SKY', '1', '0xbE286431454714F511008713973d3B053A2d38f3', 'timelock', '2019-10-20');
-- DSPauseProxy (the actual address that holds funds and executes spells)
SELECT seed_treasury('SKY', '1', '0xBE8E3e3618f7474F8cB1d074A26afFef007E98FB', 'treasury', '2019-10-20');

-- ================================================================
-- 5. CRV (Curve DAO) - Ethereum
-- ================================================================
-- Dedicated treasury contract
SELECT seed_treasury('CRV', '1', '0x6508eF65b0Bd57eaBD0f1D52685A70433B2d290B', 'treasury', '2024-12-01');
-- Ownership Agent (Aragon Agent - holds DAO assets, executes ownership votes)
SELECT seed_treasury('CRV', '1', '0x40907540d8a6C65c637785e8f8B742ae6b0b9968', 'agent', '2020-08-13');
-- Parameter Agent (Aragon Agent)
SELECT seed_treasury('CRV', '1', '0x4EEb3bA4f221cA16ed4A0cC7254E2E32DF948c5f', 'agent', '2020-08-13');
-- Community Fund
SELECT seed_treasury('CRV', '1', '0xe3997288987E6297Ad550A69B31439504F513267', 'treasury', '2020-08-14');
-- Emergency DAO (Gnosis Safe multisig)
SELECT seed_treasury('CRV', '1', '0x467947EE34aF926cF1DCac093870f613C96B1E0c', 'multisig', '2021-01-13');

-- ================================================================
-- 6. LDO (Lido DAO) - Ethereum
-- ================================================================
-- Agent (Aragon Agent - this IS Lido's primary treasury)
SELECT seed_treasury('LDO', '1', '0x3e40D73EB977Dc6a537aF587D48316feE66E9C8c', 'agent', '2020-12-17');
-- Timelock (EmergencyProtectedTimelock)
SELECT seed_treasury('LDO', '1', '0xCE0425301C85c5Ea2A0873A2dEe44d78E02D2316', 'timelock', '2024-11-01');

-- ================================================================
-- 7. ENS (Ethereum Name Service) - Ethereum
-- ================================================================
SELECT seed_treasury('ENS', '1', '0xFe89cc7aBB2C4183683ab71653C4cdc9B02D44b7', 'timelock', '2021-11-09');

-- ================================================================
-- 8. DEXE (DeXe Protocol DAO) - Ethereum
-- ================================================================
-- GovUserKeeper (treasury proxy)
SELECT seed_treasury('DEXE', '1', '0xbE8cB128fBCf13f7F7A362c3820f376b0971B7B2', 'treasury', '2024-03-05');

-- ================================================================
-- 9. COMP (Compound - shared timelock for both Bravo and OZ) - Ethereum
-- ================================================================
SELECT seed_treasury('COMP', '1', '0x6d903f6003cca6255D85CcA4D3B5E5146dC33925', 'timelock', '2019-09-01');
SELECT seed_treasury('COMP-BRAVO', '1', '0x6d903f6003cca6255D85CcA4D3B5E5146dC33925', 'timelock', '2019-09-01');
SELECT seed_treasury('COMP-OZ', '1', '0x6d903f6003cca6255D85CcA4D3B5E5146dC33925', 'timelock', '2019-09-01');

-- ================================================================
-- 10. MPL (Maple Finance) - Ethereum
-- ================================================================
-- DAO Multisig (Gnosis Safe)
SELECT seed_treasury('MPL', '1', '0xd6d4Bcde6c816F17889f1Dd3000aF0261B03a196', 'multisig', '2020-09-08');
-- Timelock
SELECT seed_treasury('MPL', '1', '0x2eFFf88747EB5a3FF00d4d8d0f0800E306C0426b', 'timelock', '2025-01-15');
-- Dedicated treasury (MapleTreasury)
SELECT seed_treasury('MPL', '1', '0xa9466EaBd096449d650D5AEB0dD3dA6F52FD0B19', 'treasury', '2021-05-01');

-- ================================================================
-- 11. W (Wormhole) - Ethereum
-- ================================================================
SELECT seed_treasury('W', '1', '0xfBc580c0289121673EfB7375fF111bD2A4db4654', 'timelock', '2024-12-01');

-- ================================================================
-- 12. DIVA (Diva Staking) - Ethereum
-- ================================================================
SELECT seed_treasury('DIVA', '1', '0x4eBB20995B6264b4b1E25f4473a4636CDB6a9790', 'timelock', '2023-06-18');

-- ================================================================
-- 13. BTRST (BrainTrust) - Ethereum
-- ================================================================
SELECT seed_treasury('BTRST', '1', '0xb6f1F016175588a049fDA12491cF3686De33990B', 'timelock', '2021-09-14');

-- ================================================================
-- 14. UMA - Ethereum
-- No dedicated treasury/timelock. Governor contracts hold voting but UMA
-- uses an oracle-based model. No fund-holding address to seed.
-- ================================================================

-- ================================================================
-- 15. POOH - Ethereum
-- ================================================================
-- POOH Crew treasury-timelock
SELECT seed_treasury('POOH', '1', '0xe8A93664A164D302c83Bbc02CE97502563134aDA', 'treasury', '2023-06-20');
-- PoohGov timelock
SELECT seed_treasury('POOH', '1', '0xfefd45d6b7d04A9935C50Ef9226484A385071566', 'timelock', '2023-06-15');

-- ================================================================
-- 16. AGLD (Adventure Gold) - Ethereum
-- ================================================================
SELECT seed_treasury('AGLD', '1', '0xEd4f981249Dde7Cd3c295fc28CB934D4682d7ef9', 'timelock', '2024-10-01');

-- ================================================================
-- 17. DEGEN-DOGS - Polygon
-- ================================================================
-- DegenDAOExecutor (combined timelock + treasury)
SELECT seed_treasury('DEGEN-DOGS', '137', '0xb6021d0b1e63596911f2cCeEF5c14f2db8f28Ce1', 'executor', '2024-04-01');

-- ================================================================
-- 18. PAPER (Dope Wars) - Ethereum
-- ================================================================
SELECT seed_treasury('PAPER', '1', '0xB57Ab8767CAe33bE61fF15167134861865F7D22C', 'timelock', '2021-09-16');

-- ================================================================
-- 19. HOP - Ethereum
-- ================================================================
SELECT seed_treasury('HOP', '1', '0xeeA8422a08258e73c139Fc32a25e10410c14bd7a', 'timelock', '2022-05-10');

-- ================================================================
-- 20. OCA (onchainaustria) - Arbitrum
-- No treasury/timelock in the contracts. Governor-only model.
-- ================================================================

-- ================================================================
-- 21. RARI (Rari Capital) - Ethereum (v1 contracts)
-- ================================================================
SELECT seed_treasury('RARI', '1', '0x7e9c956e3EFA81Ace71905Ff0dAEf1A71f42CBC5', 'timelock', '2022-10-01');

-- ================================================================
-- 22. INV (Inverse Finance) - Ethereum
-- ================================================================
SELECT seed_treasury('INV', '1', '0x926dF14a23BE491164dCF93f4c468A50ef659D5B', 'timelock', '2021-01-03');

-- ================================================================
-- 23. NOUNS-PUB (Public Nouns) - Ethereum
-- ================================================================
-- Old executor/treasury
SELECT seed_treasury('NOUNS-PUB', '1', '0x0BC3807Ec262cB779b38D65b38158acC3bfedE10', 'executor', '2021-10-20', '2022-09-20');
-- Current executor/treasury (NounsDAOExecutor holds all auction proceeds)
SELECT seed_treasury('NOUNS-PUB', '1', '0x553826Cb0D0Ee63155920F42b4E60aaE6607DFCB', 'executor', '2022-09-20');

-- ================================================================
-- 24. UDT (Unlock Protocol) - Ethereum
-- ================================================================
SELECT seed_treasury('UDT', '1', '0x17EEDFb0a6E6e06E95B3A1F928dc4024240BC76B', 'timelock', '2021-09-03');

-- ================================================================
-- 25. UP (Unlock Protocol on Base) - Base
-- ================================================================
SELECT seed_treasury('UP', '8453', '0xB34567C4cA697b39F72e1a8478f285329A98ed1b', 'timelock', '2023-09-25');

-- ================================================================
-- 26. HIFI (Hifi Finance) - Ethereum
-- ================================================================
SELECT seed_treasury('HIFI', '1', '0xAC46Db50B44BBeF8DC25f778359e1834248147F7', 'timelock', '2022-12-01');

-- ================================================================
-- 27. FLT (Fluence) - Ethereum
-- ================================================================
-- Timelock (Executor)
SELECT seed_treasury('FLT', '1', '0xf5693Bbe961F166a2fE96094d25567f7517f27B7', 'timelock', '2024-03-01');
-- Treasury (Gnosis Safe multisig)
SELECT seed_treasury('FLT', '1', '0x7F629403fDCC02aD83aA5debd1D4B1548982afaC', 'multisig', '2024-02-20');

-- ================================================================
-- 28. GTC (Gitcoin) - Ethereum
-- ================================================================
SELECT seed_treasury('GTC', '1', '0x57a8865cfB1eCEf7253c27da6B4BC3dAEE5Be518', 'timelock', '2021-05-25');

-- ================================================================
-- 29. GMX - Arbitrum
-- ================================================================
SELECT seed_treasury('GMX', '42161', '0x4bd1cdAab4254fC43ef6424653cA2375b4C94C0E', 'timelock', '2024-02-01');

-- ================================================================
-- 30. TORN (Tornado Cash) - Ethereum
-- ================================================================
SELECT seed_treasury('TORN', '1', '0x2F50508a8a3D323B91336FA3eA6ae50E55f32185', 'treasury', '2021-10-07');

-- ================================================================
-- 31. KNC (Kyber Network Crystal) - Ethereum
-- ================================================================
-- Short executor (current)
SELECT seed_treasury('KNC', '1', '0x41f5D722e6471c338392884088bD03340f50b3b5', 'executor', '2021-03-31');
-- Long executor (current)
SELECT seed_treasury('KNC', '1', '0x7d4d05B1a1E5775a9C6ca248ABBE629B52C1D9D9', 'executor', '2022-02-20');
-- Old long executor
SELECT seed_treasury('KNC', '1', '0x6758A66cD25fef7767A44895041678Fc4Ae9AfD0', 'executor', '2021-03-31', '2022-02-20');
-- Treasury (MultiSigWalletWithDailyLimit - multisig!)
SELECT seed_treasury('KNC', '1', '0x91c9D4373B077eF8082F468C7c97f2c499e36F5b', 'multisig', '2021-06-01');
-- Old treasury multisig
SELECT seed_treasury('KNC', '1', '0xE6A7338cba0A1070AdfB22c07115299605454713', 'multisig', '2020-06-01', '2021-06-01');

-- ================================================================
-- 32. AUDIO (Audius) - Ethereum
-- Governor proxy also acts as timelock AND treasury
-- ================================================================
SELECT seed_treasury('AUDIO', '1', '0x4DEcA517D6817B6510798b7328F2314d3003AbAC', 'treasury', '2020-10-18');

-- ================================================================
-- 33. FORTH (Ampleforth) - Ethereum
-- ================================================================
SELECT seed_treasury('FORTH', '1', '0x223592a191ECfC7FDC38a9256c3BD96E771539A9', 'timelock', '2021-06-27');

-- ================================================================
-- 34. RAD (Radworks) - Ethereum
-- ================================================================
SELECT seed_treasury('RAD', '1', '0x8dA8f82d2BbDd896822de723F55D6EdF416130ba', 'timelock', '2021-02-26');

-- ================================================================
-- 35. IDLE - Ethereum
-- ================================================================
-- Timelock
SELECT seed_treasury('IDLE', '1', '0xD6dABBc2b275114a2366555d6C481EF08FDC2556', 'timelock', '2020-11-01');
-- Main treasury (Gnosis Safe multisig)
SELECT seed_treasury('IDLE', '1', '0xFb3bD022D5DAcF95eE28a6B07825D4Ff9C5b3814', 'multisig', '2021-02-01');

-- ================================================================
-- 36. HAI - Optimism
-- ================================================================
SELECT seed_treasury('HAI', '10', '0xd68e7D20008a223dD48A6076AAf5EDd4fe80a899', 'timelock', '2024-01-01');
-- Treasury (Gnosis Safe multisig)
SELECT seed_treasury('HAI', '10', '0xDCb421Cc4Cbb7267F3b2cAcAb44Ec18AEbEd6724', 'multisig', '2024-01-01');

-- ================================================================
-- 37. UNION - Ethereum
-- ================================================================
SELECT seed_treasury('UNION', '1', '0xBBD3321f377742c4b3fe458b270c2F271d3294D8', 'timelock', '2021-09-20');
-- Dedicated treasury contract
SELECT seed_treasury('UNION', '1', '0x6DBDe0E7e563E34A53B1130D6B779ec8eD34B4B9', 'treasury', '2021-09-20');

-- ================================================================
-- 38. OD (Open Dollar) - Arbitrum
-- ================================================================
SELECT seed_treasury('OD', '42161', '0x7A528eA3E06D85ED1C22219471Cf0b1851943903', 'timelock', '2024-01-01');

-- ================================================================
-- 39. CTX (Cryptex) - Ethereum
-- ================================================================
SELECT seed_treasury('CTX', '1', '0xa54074b2cc0e96a43048d4a68472F7F046aC0DA8', 'timelock', '2021-06-01');

-- ================================================================
-- 40. T (Threshold Network) - Ethereum
-- ================================================================
SELECT seed_treasury('T', '1', '0x87F005317692D05BAA4193AB0c961c69e175f45f', 'timelock', '2022-01-19');

-- ================================================================
-- 41. ANVIL - Ethereum
-- ================================================================
SELECT seed_treasury('ANVIL', '1', '0x4eeB7c5BB75Fc0DBEa4826BF568FD577f62cad21', 'timelock', '2024-06-15');
-- CollateralVault (dedicated treasury)
SELECT seed_treasury('ANVIL', '1', '0x5d2725fdE4d7Aa3388DA4519ac0449Cc031d675f', 'vault', '2024-09-15');

-- ================================================================
-- 42. TRU (TrueFi) - Ethereum
-- ================================================================
SELECT seed_treasury('TRU', '1', '0x4f4AC7a7032A14243aEbDa98Ee04a5D7Fe293d07', 'timelock', '2022-02-15');
-- Distributor treasury (RatingAgencyV2Distributor)
SELECT seed_treasury('TRU', '1', '0x6151570934470214592AA051c28805cF4744BCA7', 'treasury', '2021-02-01');

-- ================================================================
-- 43. POOL (PoolTogether) - Ethereum
-- ================================================================
SELECT seed_treasury('POOL', '1', '0x42cd8312D2BCe04277dD5161832460e95b24262E', 'timelock', '2021-02-17');

-- ================================================================
-- 44. SEAM (Seamless) - Base
-- ================================================================
SELECT seed_treasury('SEAM', '8453', '0x639d2dD24304aC2e6A691d8c1cFf4a2665925fee', 'timelock', '2023-11-15');

-- ================================================================
-- 45. ATF (Antfarm) - Ethereum
-- ================================================================
SELECT seed_treasury('ATF', '1', '0x529C78Ee582e4293a20Ab60c848506eADd8723D8', 'timelock', '2023-01-15');

-- ================================================================
-- 46. SUMMER (Summer.fi) - Base
-- ================================================================
SELECT seed_treasury('SUMMER', '8453', '0x447BF9d1485ABDc4C1778025DfdfbE8b894C3796', 'timelock', '2024-11-01');

-- ================================================================
-- 47. REG (RealToken Ecosystem) - Gnosis
-- ================================================================
-- Timelock (named REGTreasuryDAO - combined timelock+treasury)
SELECT seed_treasury('REG', '100', '0x3f2d192F64020dA31D44289d62DB82adE6ABee6c', 'timelock', '2024-06-01');

-- ================================================================
-- 48. FAME (Fame Lady Society) - Base
-- ================================================================
SELECT seed_treasury('FAME', '8453', '0x08f0194e4ca54ebDF4C28cD06273053e0eDc2Da0', 'timelock', '2024-09-01');

-- ================================================================
-- 49. GRG (RigoBlock) - Ethereum
-- ================================================================
-- GrgVault (dedicated treasury vault)
SELECT seed_treasury('GRG', '1', '0xfbd2588b170Ff776eBb1aBbB58C0fbE3ffFe1931', 'vault', '2020-12-01');

-- ================================================================
-- 50. LSK (Lisk) - Lisk
-- ================================================================
SELECT seed_treasury('LSK', '1135', '0x2294A7f24187B84995A2A28112f82f07BE1BceAD', 'timelock', '2024-05-01');

-- Clean up the helper function
DROP FUNCTION IF EXISTS seed_treasury;
