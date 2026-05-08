package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ifrunruhin12/money-manager/internal/domain"
)

// AccountRepository defines persistence operations for accounts.
type AccountRepository interface {
	GetByUserID(ctx context.Context, userID string) (*domain.Account, error)
	UpdateStartingBalance(ctx context.Context, userID string, balance int) error
	UpdateTimezone(ctx context.Context, userID string, tz string) error
	AdjustBalance(ctx context.Context, userID string, delta int) error
	SetDirty(ctx context.Context, userID string, dirty bool) error
	SetReconciled(ctx context.Context, userID string) error
	// InsertTx inserts a new account within an existing transaction.
	InsertTx(ctx context.Context, tx pgx.Tx, account domain.Account) error
}

type accountRepository struct {
	db *pgxpool.Pool
}

// NewAccountRepository creates a new AccountRepository backed by the given pool.
func NewAccountRepository(db *pgxpool.Pool) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) GetByUserID(ctx context.Context, userID string) (*domain.Account, error) {
	panic("not implemented")
}

func (r *accountRepository) UpdateStartingBalance(ctx context.Context, userID string, balance int) error {
	panic("not implemented")
}

func (r *accountRepository) UpdateTimezone(ctx context.Context, userID string, tz string) error {
	panic("not implemented")
}

func (r *accountRepository) AdjustBalance(ctx context.Context, userID string, delta int) error {
	panic("not implemented")
}

func (r *accountRepository) SetDirty(ctx context.Context, userID string, dirty bool) error {
	panic("not implemented")
}

func (r *accountRepository) SetReconciled(ctx context.Context, userID string) error {
	panic("not implemented")
}

// InsertTx inserts a new account row within the provided transaction.
func (r *accountRepository) InsertTx(ctx context.Context, tx pgx.Tx, account domain.Account) error {
	_, err := tx.Exec(ctx,
		`INSERT INTO accounts (id, user_id, starting_balance, current_balance, balance_dirty, currency, timezone, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		account.ID, account.UserID, account.StartingBalance, account.CurrentBalance,
		account.BalanceDirty, account.Currency, account.Timezone, account.CreatedAt,
	)
	return err
}
