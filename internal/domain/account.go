package domain

import "time"

type Account struct {
	ID               string     `db:"id"`
	UserID           string     `db:"user_id"`
	StartingBalance  int        `db:"starting_balance"`
	CurrentBalance   int        `db:"current_balance"`
	BalanceDirty     bool       `db:"balance_dirty"`
	LastReconciledAt *time.Time `db:"last_reconciled_at"`
	Currency         string     `db:"currency"`
	Timezone         string     `db:"timezone"`
	CreatedAt        time.Time  `db:"created_at"`
}
