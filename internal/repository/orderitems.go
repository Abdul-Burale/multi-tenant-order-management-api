type OrderItemRepository interface {
	Create(ctx context.Context, item *models.OrderItem) error
	GetByOrderID(ctx context.Context, orderID uuid.UUID) ([]models.OrderItem, error)
}

type PostgresOrderItemRepository struct {
	db *sql.DB
}

func NewPostgresOrderItemRepository(db *sql.DB) *PostgresOrderItemRepository {
	return &PostgresOrderItemRepository{db: db}
}

func (r *PostgresOrderItemRepository) Create(ctx context.Context, item *models.OrderItem) error {
	query := `
		INSERT INTO order_items (orderItemID, order_id, product_id, quantity, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		item.OrderItemID,
		item.OrderID,
		item.ProductID,
		item.Quantity,
		item.CreatedAt,
		item.UpdatedAt,
	)

	return err
}
