package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ifrunruhin12/money-manager/internal/domain"
)

// CategoryRepository defines persistence operations for categories.
type CategoryRepository interface {
	Insert(ctx context.Context, c domain.Category) error
	Update(ctx context.Context, c domain.Category) error
	Delete(ctx context.Context, id string) error
	ListByUser(ctx context.Context, userID string) ([]domain.Category, error)
	IsReferencedByTransactions(ctx context.Context, id string) (bool, error)
	IsReferencedByBigBuys(ctx context.Context, id string) (bool, error)
	// InsertTx inserts a category within an existing transaction.
	InsertTx(ctx context.Context, tx pgx.Tx, c domain.Category) error
}

type categoryRepository struct {
	db *pgxpool.Pool
}

// NewCategoryRepository creates a new CategoryRepository backed by the given pool.
func NewCategoryRepository(db *pgxpool.Pool) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Insert(ctx context.Context, c domain.Category) error {
	panic("not implemented")
}

func (r *categoryRepository) Update(ctx context.Context, c domain.Category) error {
	panic("not implemented")
}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented")
}

func (r *categoryRepository) ListByUser(ctx context.Context, userID string) ([]domain.Category, error) {
	panic("not implemented")
}

func (r *categoryRepository) IsReferencedByTransactions(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

func (r *categoryRepository) IsReferencedByBigBuys(ctx context.Context, id string) (bool, error) {
	panic("not implemented")
}

// InsertTx inserts a category row within the provided transaction.
func (r *categoryRepository) InsertTx(ctx context.Context, tx pgx.Tx, c domain.Category) error {
	_, err := tx.Exec(ctx,
		`INSERT INTO categories (id, user_id, name, created_at) VALUES ($1, $2, $3, $4)`,
		c.ID, c.UserID, c.Name, c.CreatedAt,
	)
	return err
}
