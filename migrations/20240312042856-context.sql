-- +migrate Up
CREATE TABLE IF NOT EXISTS context
(
    id                UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    session_id        UUID                                      NOT NULL REFERENCES session (id),
    -- Since we are going to migrate to microservice architecture, we don't really need to specify FK here.
    schema_variant_id UUID                                      NOT NULL,
    schema_x_path     TEXT,
    metadata          JSONB,
    state             TEXT,
    created_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS context;