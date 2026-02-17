package services

import (
	"context"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type OrderItemService struct {
	repo repository.OrderItemRepository
}

func NewOrderItemService(repo repository.OrderItemRepository) *OrderItemService {
	return &OrderItemService{repo: repo}
}

func (s *OrderItemService) GetItemsByOrder(ctx context.Context, orderID uuid.UUID) ([]models.OrderItems, error) {
	return s.repo.GetByOrderID(ctx, orderID)
}
