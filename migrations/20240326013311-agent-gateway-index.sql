-- +migrate Up
CREATE INDEX IF NOT EXISTS idx_agent_gateway_active ON session(gateway, agent, valid_until, closed_at, state);

-- +migrate Down
DROP INDEX IF EXISTS idx_agent_gateway_active;