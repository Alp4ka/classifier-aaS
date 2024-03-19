-- +migrate Up

CREATE TABLE IF NOT EXISTS schema_variant
(
    id          UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    description JSONB,
    editable    BOOL                                      NOT NULL,
    created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);

CREATE TABLE IF NOT EXISTS schema
(
    id                UUID PRIMARY KEY            DEFAULT uuid_generate_v4(),
    gateway           TEXT                                      NOT NULL,
    actual_variant_id UUID REFERENCES schema_variant (id)       NOT NULL,
    created_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL,
    updated_at        TIMESTAMP WITHOUT TIME ZONE DEFAULT now() NOT NULL
);


-- +migrate Down
DROP TABLE IF EXISTS schema_variant CASCADE;
DROP TABLE IF EXISTS schema CASCADE;