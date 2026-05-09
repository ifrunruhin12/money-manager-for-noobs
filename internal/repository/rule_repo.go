package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ifrunruhin12/money-manager/internal/domain"
)

// RuleRepository defines persistence operations for fixed rules.
// Mutating methods accept a DBTX so callers can pass either a *pgxpool.Pool
// or a pgx.Tx (inside a caller-managed transaction).
type RuleRepository interface {
	// InsertFixed persists a new fixed rule.
	InsertFixed(ctx context.Context, db DBTX, rule domain.FixedRule) error
	// UpdateFixed persists changes to an existing fixed rule.
	UpdateFixed(ctx context.Context, db DBTX, rule domain.FixedRule) error
	// DeleteFixed soft-deletes a fixed rule by setting is_active = false.
	DeleteFixed(ctx context.Context, db DBTX, id string) error
	// ListActiveFixed returns all active fixed rules for a user.
	ListActiveFixed(ctx context.Context, userID string) ([]domain.FixedRule, error)
}

type ruleRepository struct {
	db *pgxpool.Pool
}

// NewRuleRepository creates a new RuleRepository backed by the given pool.
func NewRuleRepository(db *pgxpool.Pool) RuleRepository {
	return &ruleRepository{db: db}
}

const ruleSelectCols = `id, user_id, name, category_id, amount, frequency, is_active, created_at`

func scanRuleRow(rows pgx.Rows) (domain.FixedRule, error) {
	var r domain.FixedRule
	err := rows.Scan(
		&r.ID, &r.UserID, &r.Name, &r.CategoryID,
		&r.Amount, &r.Frequency, &r.IsActive, &r.CreatedAt,
	)
	return r, err
}

func (r *ruleRepository) InsertFixed(ctx context.Context, db DBTX, rule domain.FixedRule) error {
	_, err := db.Exec(ctx,
		`INSERT INTO rules (id, user_id, name, category_id, amount, frequency, is_active, created_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		rule.ID, rule.UserID, rule.Name, rule.CategoryID,
		rule.Amount, rule.Frequency, rule.IsActive, rule.CreatedAt,
	)
	return err
}

func (r *ruleRepository) UpdateFixed(ctx context.Context, db DBTX, rule domain.FixedRule) error {
	tag, err := db.Exec(ctx,
		`UPDATE rules
		 SET name = $1, category_id = $2, amount = $3, frequency = $4
		 WHERE id = $5 AND user_id = $6 AND is_active = TRUE`,
		rule.Name, rule.CategoryID, rule.Amount, rule.Frequency,
		rule.ID, rule.UserID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *ruleRepository) DeleteFixed(ctx context.Context, db DBTX, id string) error {
	tag, err := db.Exec(ctx,
		`UPDATE rules
		 SET is_active = FALSE
		 WHERE id = $1 AND is_active = TRUE`,
		id,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}

func (r *ruleRepository) ListActiveFixed(ctx context.Context, userID string) ([]domain.FixedRule, error) {
	rows, err := r.db.Query(ctx,
		`SELECT `+ruleSelectCols+`
		 FROM rules
		 WHERE is_active = TRUE AND user_id = $1
		 ORDER BY created_at ASC`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules []domain.FixedRule
	for rows.Next() {
		rule, err := scanRuleRow(rows)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, rows.Err()
}
