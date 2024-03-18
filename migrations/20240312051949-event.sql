-- +migrate Up
CREATE TABLE IF NOT EXISTS event
(
    id            BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    context_id    UUID                                      NOT NULL REFERENCES context (id),
    req           TEXT,
    resp          TEXT,
    schema_x_path TEXT,
    created_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at    TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS event;
