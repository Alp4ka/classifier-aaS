-- +migrate Up
CREATE TABLE IF NOT EXISTS session
(
    id          UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    state       TEXT                                      NOT NULL,
    agent       TEXT                                      NOT NULL,
    gateway     TEXT                                      NOT NULL,
    valid_until TIMESTAMP WITHOUT TIME ZONE               NOT NULL,
    closed_at   TIMESTAMP WITHOUT TIME ZONE,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS session;
