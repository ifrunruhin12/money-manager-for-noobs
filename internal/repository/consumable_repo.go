package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ifrunruhin12/money-manager/internal/domain"
)

// ConsumableRepository defines persistence operations for consumable rules.
// Mutating methods accept a DBTX so callers can pass either a *pgxpool.Pool
// or a pgx.Tx (inside a caller-managed transaction).
type ConsumableRepository interface {
	// Insert persists a new consumable rule.
	Insert(ctx context.Context, db DBTX, c domain.ConsumableRule) error
	// Update persists changes to an existing consumable rule.
	Update(ctx context.Context, db DBTX, c domain.ConsumableRule) error
	// GetByID returns a consumable by ID, or ErrNotFound if it doesn't exist or is deleted.
	GetByID(ctx context.Context, id string) (*domain.ConsumableRule, error)
	// GetByIDForUpdate returns a consumable by ID using SELECT FOR UPDATE, locking the row.
	// Must be called within a transaction (db must be a pgx.Tx).
	GetByIDForUpdate(ctx context.Context, db DBTX, id string) (*domain.ConsumableRule, error)
	// ListActive returns all non-deleted consumables for a user.
	ListActive(ctx context.Context, userID string) ([]domain.ConsumableRule, error)
}

type consumableRepository struct {
	db *pgxpool.Pool
}

// NewConsumableRepository creates a new ConsumableRepository backed by the given pool.
func NewConsumableRepository(db *pgxpool.Pool) ConsumableRepository {
	return &consumableRepository{db: db}
}

const consumableSelectCols = `id, user_id, name, stock, usage_per_day, restock_amount, restock_cost, restock_threshold, is_depleted, last_restock_date, created_at`

func scanConsumableRow(row pgx.Row) (*domain.ConsumableRule, error) {
	var c domain.ConsumableRule
	err := row.Scan(
		&c.ID, &c.UserID, &c.Name, &c.Stock,
		&c.UsagePerDay, &c.RestockAmount, &c.RestockCost, &c.RestockThreshold,
		&c.IsDepleted, &c.LastRestockDate, &c.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func scanConsumableRows(rows pgx.Rows) (domain.ConsumableRule, error) {
	var c domain.ConsumableRule
	err := rows.Scan(
		&c.ID, &c.UserID, &c.Name, &c.Stock,
		&c.UsagePerDay, &c.RestockAmount, &c.RestockCost, &c.RestockThreshold,
		&c.IsDepleted, &c.LastRestockDate, &c.CreatedAt,
	)
	return c, err
}

func (r *consumableRepository) Insert(ctx context.Context, db DBTX, c domain.ConsumableRule) error {
	_, err := db.Exec(ctx,
		`INSERT INTO consumables
		 (id, user_id, name, stock, usage_per_day, restock_amount, restock_cost, restock_threshold, is_depleted, last_restock_date, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		c.ID, c.UserID, c.Name, c.Stock,
		c.UsagePerDay, c.RestockAmount, c.RestockCost, c.RestockThreshold,
		c.IsDepleted, c.LastRestockDate, c.CreatedAt,
	)
	return err
}

func (r *consumableRepository) Update(ctx context.Context, db DBTX, c domain.ConsumableRule) error {
	tag, err := db.Exec(ctx,
		`UPDATE consumables
		 SET name = $1, stock = $2, usage_per_day = $3, restock_amount = $4,
		     restock_cost = $5, restock_threshold = $6, is_depleted = $7, last_restock_date = $8
		 WHERE id = $9 AND user_id = $10`,
		c.Name, c.Stock, c.UsagePerDay, c.RestockAmount,
		c.RestockCost, c.RestockThreshold, c.IsDepleted, c.LastRestockDate,
		c.ID, c.UserID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *consumableRepository) GetByID(ctx context.Context, id string) (*domain.ConsumableRule, error) {
	row := r.db.QueryRow(ctx,
		`SELECT `+consumableSelectCols+`
		 FROM consumables
		 WHERE id = $1`,
		id,
	)
	c, err := scanConsumableRow(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return c, nil
}

func (r *consumableRepository) GetByIDForUpdate(ctx context.Context, db DBTX, id string) (*domain.ConsumableRule, error) {
	row := db.QueryRow(ctx,
		`SELECT `+consumableSelectCols+`
		 FROM consumables
		 WHERE id = $1
		 FOR UPDATE`,
		id,
	)
	c, err := scanConsumableRow(row)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return c, nil
}

func (r *consumableRepository) ListActive(ctx context.Context, userID string) ([]domain.ConsumableRule, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+consumableSelectCols+`
		 FROM consumables
		 WHERE user_id = $1
		 ORDER BY created_at ASC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consumables []domain.ConsumableRule
	for rows.Next() {
		c, err := scanConsumableRows(rows)
		if err != nil {
			return nil, err
		}
		consumables = append(consumables, c)
	}
	return consumables, rows.Err()
}
