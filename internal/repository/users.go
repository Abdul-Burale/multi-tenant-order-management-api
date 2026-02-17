package repository

import (
	"context"
	"database/sql"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/google/uuid"
)

type UsersRespository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func (r *PostgresUserRepository) NewPostgresUserRespository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
	INSERT INTO users (usersID, tenant_id, email, name, role, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		user.UserID,
		user.TenantID,
		user.Email,
		user.Name,
		user.Role,
		user.CreatedAt,
	)

	return err
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
	SELECT userID, tenantID, email, name, role, created_at
	FROM users
	WHERE usersID = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var user models.User
	err := row.Scan(
		&user.UserID,
		&user.TenantID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
	SELECT userID, tenantID, email, name, role, created_at
	FROM users
	WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var user models.User
	err := row.Scan(
		&user.UserID,
		&user.TenantID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
