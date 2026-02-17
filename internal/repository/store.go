package repository

import (
	"context"
	"database/sql"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/google/uuid"
)

type StoreRepository interface {
	Create(ctx context.Context, store *models.Stores) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Stores, error)
}

type PostgresStoreRepository struct {
	db *sql.DB
}

func NewPostgresStoreRepository(db *sql.DB) *PostgresStoreRepository {
	return &PostgresStoreRepository{db: db}
}

func (r *PostgresStoreRepository) Create(ctx context.Context, store *models.Stores) error {
	query := `
		INSERT INTO stores (storeID, tenant_id, name, description, type, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		store.StoreID,
		store.TenantID,
		store.Name,
		store.Description,
		store.Type,
		store.CreatedAt,
	)

	return err
}

func (r *PostgresStoreRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Stores, error) {
	query := `
		SELECT storeID, tenant_id, name, description, type, created_at
		FROM stores
		WHERE storeID = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var store models.Stores
	err := row.Scan(
		&store.StoreID,
		&store.TenantID,
		&store.Name,
		&store.Description,
		&store.Type,
		&store.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &store, nil
}
