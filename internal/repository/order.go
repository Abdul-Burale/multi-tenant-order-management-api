type OrderRepository interface {
	Create(ctx context.Context, order *models.Order) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Order, error)
}

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Create(ctx context.Context, order *models.Order) error {
	query := `
		INSERT INTO orders (orderID, store_id, user_id, status, total_amount, stripe_payment_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.ExecContext(ctx, query,
		order.OrderID,
		order.StoreID,
		order.UserID,
		order.Status,
		order.TotalAmount,
		order.StripePaymentID,
		order.CreatedAt,
		order.UpdatedAt,
	)

	return err
}
