package services

import (
	"context"
	"errors"
	"time"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, storeID uuid.UUID, name, description string, price int, status models.ProductStatus) (*models.Product, error) {
	if name == "" {
		return nil, errors.New("product name cannot be empty")
	}

	product := &models.Product{
		ProductID:   uuid.New(),
		StoreID:     storeID,
		Name:        name,
		Description: &description,
		Price:       price,
		Status:      status,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (*models.Product, error) {
	return s.repo.GetByID(ctx, id)
}
