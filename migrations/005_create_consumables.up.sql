CREATE TABLE consumables (
    id                TEXT PRIMARY KEY,
    user_id           TEXT NOT NULL REFERENCES users(id),
    name              TEXT NOT NULL,
    stock             INTEGER NOT NULL DEFAULT 0,
    usage_per_day     INTEGER NOT NULL CHECK (usage_per_day > 0),
    restock_amount    INTEGER NOT NULL CHECK (restock_amount > 0),
    restock_cost      INTEGER NOT NULL CHECK (restock_cost > 0),
    restock_threshold INTEGER NOT NULL CHECK (restock_threshold >= 0),
    is_depleted       BOOLEAN NOT NULL DEFAULT FALSE,
    last_restock_date DATE,
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_consumables_user ON consumables(user_id);
