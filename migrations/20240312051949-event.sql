-- +migrate Up
CREATE TABLE IF NOT EXISTS event
(
    id             BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    session_id     UUID                                      NOT NULL REFERENCES session (id),
    req            TEXT,
    resp           TEXT,
    schema_node_id UUID,
    created_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at     TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS event;