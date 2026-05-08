CREATE TABLE transactions (
    id              TEXT PRIMARY KEY,
    user_id         TEXT NOT NULL REFERENCES users(id),
    type            TEXT NOT NULL CHECK (type IN ('rule_generated', 'manual', 'override')),
    category_id     TEXT NOT NULL REFERENCES categories(id),
    amount          INTEGER NOT NULL CHECK (amount != 0),
    is_skipped      BOOLEAN NOT NULL DEFAULT FALSE,
    is_overridden   BOOLEAN NOT NULL DEFAULT FALSE,
    source_id       TEXT,
    source_type     TEXT CHECK (source_type IN ('rule', 'consumable', 'transaction')),
    note            TEXT NOT NULL DEFAULT '',
    date            TIMESTAMPTZ NOT NULL,
    generation_date DATE,
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at      TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX uniq_tx_source_date
    ON transactions(source_id, source_type, generation_date)
    WHERE source_id IS NOT NULL AND deleted_at IS NULL;

CREATE INDEX idx_transactions_user_date       ON transactions(user_id, date DESC) WHERE deleted_at IS NULL;
CREATE INDEX idx_transactions_override_source ON transactions(source_id) WHERE source_type = 'transaction' AND deleted_at IS NULL;
CREATE INDEX idx_transactions_active          ON transactions(user_id) WHERE is_skipped = FALSE AND is_overridden = FALSE AND deleted_at IS NULL;
