type ProductRepository interface {
	Create(ctx context.Context, product *models.Product) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error)
}

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) Create(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO products (productID, store_id, name, description, price, active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		product.ProductID,
		product.StoreID,
		product.Name,
		product.Description,
		product.Price,
		product.Active,
		product.CreatedAt,
	)

	return err
}

func (r *PostgresProductRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	query := `
		SELECT productID, store_id, name, description, price, active, created_at
		FROM products
		WHERE productID = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var product models.Product
	err := row.Scan(
		&product.ProductID,
		&product.StoreID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Active,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
