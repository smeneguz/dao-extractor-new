-- ================================================================
-- Analysis Views for DAO Governance KPIs
-- These views compute metrics from the indexed on-chain data.
-- Refresh materialized views periodically after indexing runs.
-- ================================================================

-- -----------------------------------------------------------------
-- 1. GOVERNANCE PARTICIPATION
-- Per-DAO proposal and voting statistics.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_governance_participation AS
SELECT
    d.id AS dao_id,
    d.symbol AS dao_symbol,
    d.name AS dao_name,
    COUNT(DISTINCT p.id) AS total_proposals,
    COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'EXECUTED') AS approved_proposals,
    COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'DEFEATED') AS defeated_proposals,
    COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'CANCELED') AS canceled_proposals,
    COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'EXPIRED') AS expired_proposals,
    COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'ACTIVE') AS active_proposals,
    CASE WHEN COUNT(DISTINCT p.id) > 0
        THEN ROUND(100.0 * COUNT(DISTINCT p.id) FILTER (WHERE p.status = 'EXECUTED') / COUNT(DISTINCT p.id), 2)
        ELSE 0 END AS approval_rate_pct,
    COUNT(DISTINCT va.sender_address_id) AS unique_voters,
    COUNT(DISTINCT p.creator_address_id) AS unique_proposers,
    COUNT(va.id) AS total_votes,
    ROUND(AVG(EXTRACT(EPOCH FROM (p.end_time - p.start_time)) / 86400.0)::NUMERIC, 2) AS avg_voting_duration_days,
    MIN(p.creation_ts) AS first_proposal_at,
    MAX(p.creation_ts) AS last_proposal_at
FROM daos d
LEFT JOIN proposals p ON p.dao_id = d.id
LEFT JOIN vote_actions va ON va.proposal_id = p.id
GROUP BY d.id, d.symbol, d.name;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_gov_participation_dao
    ON mv_governance_participation(dao_id);

-- -----------------------------------------------------------------
-- 2. VOTER ENGAGEMENT
-- Per-proposal voting metrics.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_proposal_voting_stats AS
SELECT
    p.id AS proposal_id,
    p.dao_id,
    d.symbol AS dao_symbol,
    p.proposal_id AS onchain_proposal_id,
    p.status,
    p.title,
    p.creation_ts,
    p.quorum,
    COUNT(va.id) AS vote_count,
    COUNT(DISTINCT va.sender_address_id) AS unique_voters,
    COALESCE(SUM(va.voting_power) FILTER (WHERE va.vote = 1), 0) AS votes_for,
    COALESCE(SUM(va.voting_power) FILTER (WHERE va.vote = 0), 0) AS votes_against,
    COALESCE(SUM(va.voting_power) FILTER (WHERE va.action_type = 'ABSTAIN'), 0) AS votes_abstain,
    COALESCE(SUM(va.voting_power), 0) AS total_voting_power
FROM proposals p
JOIN daos d ON d.id = p.dao_id
LEFT JOIN vote_actions va ON va.proposal_id = p.id
GROUP BY p.id, p.dao_id, d.symbol, p.proposal_id, p.status, p.title, p.creation_ts, p.quorum;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_proposal_voting_pid
    ON mv_proposal_voting_stats(proposal_id);

-- -----------------------------------------------------------------
-- 3. TOKEN TRANSFER SUMMARY
-- Per-DAO token transfer statistics.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_token_transfer_stats AS
WITH addr_union AS (
    SELECT dao_id, chain_id, token_address, from_address AS addr FROM token_transfers
    UNION
    SELECT dao_id, chain_id, token_address, to_address AS addr FROM token_transfers
)
SELECT
    d.id AS dao_id,
    d.symbol AS dao_symbol,
    tt.chain_id,
    tt.token_address,
    COUNT(*) AS total_transfers,
    COUNT(DISTINCT tt.from_address) AS unique_senders,
    COUNT(DISTINCT tt.to_address) AS unique_receivers,
    au.unique_addresses,
    SUM(tt.amount) AS total_volume,
    COUNT(*) FILTER (WHERE tt.from_address = '0x0000000000000000000000000000000000000000') AS mint_events,
    COUNT(*) FILTER (WHERE tt.to_address = '0x0000000000000000000000000000000000000000') AS burn_events,
    MIN(tt.ts) AS first_transfer_at,
    MAX(tt.ts) AS last_transfer_at
FROM token_transfers tt
JOIN daos d ON d.id = tt.dao_id
JOIN (
    SELECT dao_id, chain_id, token_address, COUNT(DISTINCT addr) AS unique_addresses
    FROM addr_union
    GROUP BY dao_id, chain_id, token_address
) au ON au.dao_id = tt.dao_id AND au.chain_id = tt.chain_id AND au.token_address = tt.token_address
GROUP BY d.id, d.symbol, tt.chain_id, tt.token_address, au.unique_addresses;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_token_transfer_dao_chain
    ON mv_token_transfer_stats(dao_id, chain_id, token_address);

-- -----------------------------------------------------------------
-- 4. DELEGATION SUMMARY
-- Per-DAO delegation activity.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_delegation_stats AS
SELECT
    d.id AS dao_id,
    d.symbol AS dao_symbol,
    de.chain_id,
    de.token_address,
    COUNT(*) AS total_delegation_events,
    COUNT(DISTINCT de.delegator) AS unique_delegators,
    COUNT(DISTINCT de.to_delegate) AS unique_delegates,
    -- Self-delegation: delegator == to_delegate
    COUNT(*) FILTER (WHERE de.delegator = de.to_delegate) AS self_delegations,
    -- External delegation: delegator != to_delegate
    COUNT(*) FILTER (WHERE de.delegator != de.to_delegate) AS external_delegations,
    MIN(de.ts) AS first_delegation_at,
    MAX(de.ts) AS last_delegation_at
FROM delegation_events de
JOIN daos d ON d.id = de.dao_id
GROUP BY d.id, d.symbol, de.chain_id, de.token_address;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_delegation_dao_chain
    ON mv_delegation_stats(dao_id, chain_id, token_address);

-- -----------------------------------------------------------------
-- 5. VOTING POWER DISTRIBUTION (latest snapshot)
-- Current voting power per delegate per DAO.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_voting_power_latest AS
SELECT DISTINCT ON (dao_id, chain_id, delegate)
    d.id AS dao_id,
    d.symbol AS dao_symbol,
    dvc.chain_id,
    dvc.token_address,
    dvc.delegate,
    dvc.new_balance AS current_voting_power,
    dvc.block_height AS last_change_height,
    dvc.ts AS last_change_ts
FROM delegate_votes_changed dvc
JOIN daos d ON d.id = dvc.dao_id
ORDER BY dao_id, chain_id, delegate, dvc.block_height DESC, dvc.log_index DESC;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_vp_latest_dao_delegate
    ON mv_voting_power_latest(dao_id, chain_id, delegate);

-- -----------------------------------------------------------------
-- 6. DAO OVERVIEW (high-level dashboard)
-- Combines governance + token + delegation metrics per DAO.
-- -----------------------------------------------------------------
CREATE MATERIALIZED VIEW IF NOT EXISTS mv_dao_overview AS
SELECT
    d.id AS dao_id,
    d.symbol AS dao_symbol,
    d.name AS dao_name,
    COALESCE(gp.total_proposals, 0) AS total_proposals,
    COALESCE(gp.approved_proposals, 0) AS approved_proposals,
    COALESCE(gp.approval_rate_pct, 0) AS approval_rate_pct,
    COALESCE(gp.unique_voters, 0) AS unique_voters,
    COALESCE(gp.unique_proposers, 0) AS unique_proposers,
    COALESCE(gp.total_votes, 0) AS total_votes,
    gp.avg_voting_duration_days,
    COALESCE(tts.total_transfers, 0) AS total_token_transfers,
    COALESCE(tts.unique_senders, 0) AS unique_token_senders,
    COALESCE(tts.total_volume, 0) AS total_token_volume,
    COALESCE(ds.total_delegation_events, 0) AS total_delegations,
    COALESCE(ds.unique_delegators, 0) AS unique_delegators,
    COALESCE(ds.unique_delegates, 0) AS unique_delegates,
    COALESCE(ds.self_delegations, 0) AS self_delegations,
    COALESCE(ds.external_delegations, 0) AS external_delegations,
    gp.first_proposal_at,
    gp.last_proposal_at,
    tts.first_transfer_at,
    tts.last_transfer_at
FROM daos d
LEFT JOIN mv_governance_participation gp ON gp.dao_id = d.id
LEFT JOIN mv_token_transfer_stats tts ON tts.dao_id = d.id
LEFT JOIN mv_delegation_stats ds ON ds.dao_id = d.id;

CREATE UNIQUE INDEX IF NOT EXISTS idx_mv_dao_overview_dao
    ON mv_dao_overview(dao_id);

-- -----------------------------------------------------------------
-- Helper function to refresh all materialized views.
-- Usage: SELECT refresh_analysis_views();
-- -----------------------------------------------------------------
CREATE OR REPLACE FUNCTION refresh_analysis_views()
RETURNS void AS $$
BEGIN
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_governance_participation;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_proposal_voting_stats;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_token_transfer_stats;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_delegation_stats;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_voting_power_latest;
    REFRESH MATERIALIZED VIEW CONCURRENTLY mv_dao_overview;
END;
$$ LANGUAGE plpgsql;
