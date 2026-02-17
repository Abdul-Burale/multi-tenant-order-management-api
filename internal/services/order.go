package services

import (
	"context"
	"time"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	productRepo   repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, orderItemRepo repository.OrderItemRepository, productRepo repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:     orderRepo,
		orderItemRepo: orderItemRepo,
		productRepo:   productRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, storeID, userID uuid.UUID, items []models.OrderItems) (*models.Order, error) {
	order := &models.Order{
		OrderID:     uuid.New(),
		StoreID:     storeID,
		UserID:      userID,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		TotalAmount: 0,
	}

	// calculate total and validate products
	total := 0
	for i := range items {
		product, err := s.productRepo.GetByID(ctx, items[i].ProductID)
		if err != nil {
			return nil, err
		}
		total += product.Price * items[i].Quantity
		items[i].OrderID = order.OrderID
	}
	order.TotalAmount = total

	if err := s.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	for _, item := range items {
		if err := s.orderItemRepo.Create(ctx, &item); err != nil {
			return nil, err
		}
	}

	return order, nil
}

func (s *OrderService) GetOrderByID(ctx context.Context, id uuid.UUID) (*models.Order, error) {
	return s.orderRepo.GetByID(ctx, id)
}
