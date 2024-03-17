-- +migrate Up

CREATE TABLE IF NOT EXISTS schema
(
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    gateway           TEXT                        NOT NULL,
    actual_variant_id UUID                        NOT NULL,
    created_at        TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at        TIMESTAMP WITHOUT TIME ZONE NOT NULL
);


CREATE TABLE IF NOT EXISTS schema_variant
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ref_schema_id UUID                        NOT NULL REFERENCES schema (id),
    description   JSONB                       NOT NULL,
    created_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    updated_at    TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

ALTER TABLE schema
    ADD CONSTRAINT fk_schema_to_schema_variant FOREIGN KEY (actual_variant_id) REFERENCES schema_variant;

-- +migrate Down
DROP TABLE IF EXISTS schema_variant;
DROP TABLE IF EXISTS schema;
