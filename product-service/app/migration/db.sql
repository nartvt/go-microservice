CREATE TABLE IF NOT EXISTS products(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    image       VARCHAR(100) NOT NULL
    active      BOOL DEFAULT TRUE,
    price       DECIMAL(10, 2) NOT NULL,
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    deleted_at  TIMESTAMPTZ
);