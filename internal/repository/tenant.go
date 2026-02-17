package repository

import (
	"context"
	"database/sql"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/google/uuid"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *models.Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Tenant, error)
}

type PostgresTenantRepository struct {
	db *sql.DB
}

func NewPostgresTenantRepository(db *sql.DB) *PostgresTenantRepository {
	return &PostgresTenantRepository{db: db}
}

func (r *PostgresTenantRepository) Create(ctx context.Context, tenant *models.Tenant) error {
	query := `
	INSERT INTO tenants (tenantID, business_name, created_at)
	VALUES ($1, $2, $3)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		tenant.TenantID,
		tenant.BusinessName,
		tenant.Created_At,
	)

	return err
}

func (r *PostgresTenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Tenant, error) {
	query := `
	SELECT tenantID, business_name, created_at
	FROM tenants
	WHERE tenantID = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var tenant models.Tenant
	err := row.Scan(
		&tenant.TenantID,
		&tenant.BusinessName,
		&tenant.Created_At,
	)

	if err != nil {
		return nil, err
	}

	return &tenant, nil
}
