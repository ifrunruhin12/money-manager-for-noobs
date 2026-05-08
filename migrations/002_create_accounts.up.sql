CREATE TABLE accounts (
    id                  TEXT PRIMARY KEY,
    user_id             TEXT NOT NULL REFERENCES users(id),
    starting_balance    INTEGER NOT NULL DEFAULT 0,
    current_balance     INTEGER NOT NULL DEFAULT 0,
    balance_dirty       BOOLEAN NOT NULL DEFAULT FALSE,
    last_reconciled_at  TIMESTAMPTZ,
    currency            TEXT NOT NULL DEFAULT 'BDT',
    timezone            TEXT NOT NULL DEFAULT 'UTC',
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE(user_id)
);
