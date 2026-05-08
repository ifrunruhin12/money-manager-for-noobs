CREATE TABLE big_buys (
    id          TEXT PRIMARY KEY,
    user_id     TEXT NOT NULL REFERENCES users(id),
    title       TEXT NOT NULL,
    amount      INTEGER NOT NULL CHECK (amount < 0),
    category_id TEXT NOT NULL REFERENCES categories(id),
    note        TEXT NOT NULL DEFAULT '',
    date        TIMESTAMPTZ NOT NULL,
    deleted_at  TIMESTAMPTZ,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_big_buys_user_date ON big_buys(user_id, date ASC) WHERE deleted_at IS NULL;
