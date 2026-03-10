/* ------------------------------------------------------------------
   Enum that describes the current status of a proposal.
ACTIVE
Meaning: The proposal is currently open for voting. The voting period has started and has not yet ended.

Typical Scenario: A proposal has just been created and submitted to the DAO. Voters can now cast their votes ('VOTE', 'ABSTAIN') on it.

Next Possible Statuses: CANCELED, EXECUTED, DEFEATED, EXPIRED.

CANCELED
Meaning: The proposal has been withdrawn before the end of its voting period.

Typical Scenario: This action is usually performed by the original creator of the proposal. They might cancel it if they realize there's a flaw in it, if community sentiment is strongly negative, or for other strategic reasons. No further voting is possible.

This is a terminal state.

EXECUTED
Meaning: The proposal was successful. It received the required number of votes to pass, and the on-chain action it proposed has been successfully carried out.

Typical Scenario: A proposal to transfer funds from the DAO treasury passes. After the voting period ends, a transaction is triggered that moves the funds. Once that transaction is confirmed, the proposal's status is updated to EXECUTED. The details of this transaction are logged in the proposal_finalizations table.

This is a terminal state.

DEFEATED
Meaning: The proposal failed to get enough positive votes to pass.

Typical Scenario: The voting period ends, and the number of "no" votes is too high, or the number of "yes" votes is below the required threshold. The proposal is rejected and no on-chain action is taken.

This is a terminal state.

EXPIRED
Meaning: The proposal's voting period has ended, but it was neither executed nor defeated. This typically happens when a proposal fails to meet the minimum participation requirement (quorum).

Typical Scenario: A DAO requires that at least 10% of the total token supply must vote on a proposal for it to be valid. A proposal gets 90% "yes" votes, but only 5% of tokens participated in the vote. Since the 10% quorum wasn't met, the proposal cannot be executed, but it wasn't technically defeated by votes either. It simply expires.

This is a terminal state.
   ------------------------------------------------------------------ */
CREATE TYPE PROPOSAL_STATUS AS ENUM ('PENDING', 'ACTIVE', 'VOTE_CLOSED', 'CANCELED', 'EXECUTED', 'DEFEATED', 'EXPIRED');


-- Proposals table stores the definitive state of each proposal.
CREATE TABLE proposals (
    -- A unique internal ID for the proposal
    id BIGSERIAL PRIMARY KEY,

-- A composite key that uniquely identifies the proposal on-chain
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    proposal_id NUMERIC NOT NULL, -- Changed to NUMERIC to support very large numbers
    chain_id BIGINT NOT NULL REFERENCES blockchains(id),

-- Foreign keys for better integrity and performance
    creator_address_id BIGINT NOT NULL REFERENCES addresses(id),
    contract_address_id BIGINT NOT NULL REFERENCES addresses(id),

-- Core proposal details
    title TEXT NOT NULL,
    description TEXT,
    -- Type of proposal, this will be populated later on using an LLM to infer the type
    -- from the description
    type VARCHAR(64),

-- Current status of the proposal for easy querying
    status PROPOSAL_STATUS NOT NULL DEFAULT 'ACTIVE',

-- Creation event details
    creation_height BIGINT NOT NULL,
    creation_tx_hash TEXT NOT NULL,
    creation_ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,

-- The gas used to create the proposal
    gas_used BIGINT NOT NULL,
    -- The fees paid for the gas used to create the proposal
    gas_fees COIN NOT NULL,

-- Start proposal details
    start_height BIGINT NOT NULL,
    start_time TIMESTAMP WITHOUT TIME ZONE, -- This can be null since DAOs may express start time using a block height

-- End proposal details
    end_height BIGINT NOT NULL,
    end_time TIMESTAMP WITHOUT TIME ZONE, -- This can be null since DAOs may express end time using a block height

-- The voting power required to consider a proposal valid
    quorum NUMERIC NOT NULL,

    extra_metadata JSONB,

    CONSTRAINT unique_proposal_on_chain UNIQUE (dao_id, proposal_id, chain_id)
);
-- To quickly find proposals within a DAO
CREATE INDEX idx_proposals_dao ON proposals(dao_id);
-- To quickly find proposals by creator
CREATE INDEX idx_proposals_creator ON proposals(creator_address_id);
-- To efficiently filter proposals by their current status
CREATE INDEX idx_proposals_status ON proposals(status);
-- To quickly find the proposals by contract
CREATE INDEX idx_proposals_contracts_address_id ON proposals(contract_address_id);


-- This table logs the finalization event of a proposal (e.g., execution or
-- cancellation).
CREATE TABLE proposal_finalizations (
    id BIGSERIAL PRIMARY KEY,

-- A proposal can only be finalized once, creating a one-to-one relationship.
    proposal_id BIGINT NOT NULL UNIQUE REFERENCES proposals(id),

-- The transaction that triggered the finalization.
    tx_hash TEXT NOT NULL,
    height BIGINT NOT NULL,
    -- The gas used to execute the transaction
    gas_used BIGINT NOT NULL,
    -- The fees that was paid to execute the transaction
    gas_fees COIN NOT NULL,
    -- The timestamp of the block that has included the transaction
    ts TIMESTAMP WITHOUT TIME ZONE NOT NULL,

-- The status that this event triggered.
    status_triggered PROPOSAL_STATUS NOT NULL,

-- Any extra data related to the finalization (e.g., function call data for an
-- execution).
    extra_metadata JSONB,

-- Ensure that only terminal statuses are logged here.
    CONSTRAINT check_finalization_status CHECK (status_triggered IN ('CANCELED', 'EXECUTED', 'DEFEATED', 'EXPIRED'))
);

-- To quickly find proposal finalizations by proposal ID
CREATE INDEX idx_proposals_finalizations_proposal_id ON proposal_finalizations (proposal_id);

/* ------------------------------------------------------------------
   Enum that describes the kind of vote a user can cast.
   ------------------------------------------------------------------ */
CREATE TYPE VOTE_ACTION_TYPE AS ENUM ('VOTE', 'CANCEL', 'ABSTAIN');

-- Vote actions table logs every vote, cancel, or abstain action on a proposal.
CREATE TABLE vote_actions
(
    -- ID of the vote action.
    id BIGSERIAL PRIMARY KEY,

-- Foreign key to the proposal being voted on.
    proposal_id BIGINT NOT NULL REFERENCES proposals(id),

-- Address of the voter.
    sender_address_id BIGINT NOT NULL REFERENCES addresses(id),
    -- Address of the contract that has been used to cast the vote.
    contract_address_id BIGINT NOT NULL REFERENCES addresses(id),
    -- In the case of a delegated vote, this is the address of the delegator.
    delegator_address_id BIGINT REFERENCES addresses(id),

-- Hash of the transaction that has casted the vote.
    tx_hash  TEXT NOT NULL,
    -- Height at which the vote has been casted.
    height   BIGINT NOT NULL,
    -- Time at which the vote has been casted (with timezone).
    ts       TIMESTAMP WITHOUT TIME ZONE NOT NULL,

-- Type the vote has been casted.
    action_type VOTE_ACTION_TYPE NOT NULL,
    -- Vote that has been casted (NULL for CANCEL or ABSTAIN).
    vote         INTEGER,
    -- User’s voting power at the time of the vote (NULL for CANCEL or ABSTAIN).
    voting_power NUMERIC,
    -- Extra metadata that has been attached to the vote.
    extra_metadata JSONB,

    CONSTRAINT unique_vote_action UNIQUE (tx_hash, proposal_id, sender_address_id)
);
-- To quickly find all votes for a specific proposal
CREATE INDEX idx_vote_actions_proposal ON vote_actions(proposal_id);
-- To quickly find all votes from a specific user
CREATE INDEX idx_vote_actions_sender ON vote_actions(sender_address_id);
-- To efficiently query votes by time
CREATE INDEX idx_vote_actions_ts ON vote_actions(ts);

-- Table that contains the association between a DAO and the tokens that are being
-- used to govern it
CREATE TABLE dao_governance_tokens (
    id BIGSERIAL PRIMARY KEY,
    -- A unique internal ID for the DAO
    dao_id BIGINT NOT NULL REFERENCES daos(id),
    -- The token that is being used to govern the DAO
    token_id BIGINT NOT NULL REFERENCES tokens(id),
    -- Timestamp from when the token was used to govern the DAO
    used_from TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    -- Timestamp until when the token was used to govern the DAO
    -- if null, the token is still being used
    used_until TIMESTAMP WITHOUT TIME ZONE
);
-- To quickly find the tokens that are being used to govern a DAO
CREATE INDEX idx_dao_governance_tokens_dao_id ON dao_governance_tokens (dao_id);

-- To quickly find wich DAOs are using a token to govern them
CREATE INDEX idx_dao_governance_tokens_token_id ON dao_governance_tokens (token_id);

-- To quickly find the current active tokens that are being used to govern a DAO
CREATE INDEX idx_dao_governance_tokens_active ON dao_governance_tokens (dao_id, token_id) 
WHERE (used_until IS NULL);

-- To quickly find tokens by timestamp
CREATE INDEX idx_dao_governance_tokens_used_from ON dao_governance_tokens (used_from);
CREATE INDEX idx_dao_governance_tokens_used_until ON dao_governance_tokens (used_until);

