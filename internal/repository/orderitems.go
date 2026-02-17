package repository

import (
	"context"
	"database/sql"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/google/uuid"
)

type OrderItemRepository interface {
	Create(ctx context.Context, item *models.OrderItems) error
	GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]models.OrderItems, error)
}

type PostgresOrderItemRepository struct {
	db *sql.DB
}

func NewPostgresOrderItemRepository(db *sql.DB) *PostgresOrderItemRepository {
	return &PostgresOrderItemRepository{db: db}
}

func (r *PostgresOrderItemRepository) Create(ctx context.Context, item *models.OrderItems) error {
	query := `
		INSERT INTO order_items (orderItemID, order_id, product_id, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		item.OrderItemsID,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.CreatedAt,
		item.UpdatedAt,
	)

	return err
}
