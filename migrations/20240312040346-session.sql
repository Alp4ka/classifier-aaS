-- +migrate Up
CREATE TABLE IF NOT EXISTS session
(
    id          UUID PRIMARY KEY,
    state       TEXT                        NOT NULL,
    agent       TEXT                        NOT NULL,
    gateway     TEXT                        NOT NULL,
    valid_until TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS session;
