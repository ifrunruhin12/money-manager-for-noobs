CREATE TABLE rules (
    id          TEXT PRIMARY KEY,
    user_id     TEXT NOT NULL REFERENCES users(id),
    name        TEXT NOT NULL,
    category_id TEXT NOT NULL REFERENCES categories(id),
    amount      INTEGER NOT NULL CHECK (amount > 0),
    frequency   TEXT NOT NULL CHECK (frequency IN ('daily', 'weekday', 'weekend')),
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
