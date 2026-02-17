package services

import (
	"context"
	"errors"
	"time"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type StoreService struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (s *StoreService) CreateStore(ctx context.Context, tenantID uuid.UUID, name string, storeType models.StoreType, description *string) (*models.Store, error) {
	if name == "" {
		return nil, errors.New("store name cannot be empty")
	}

	store := &models.Store{
		StoreID:     uuid.New(),
		TenantID:    tenantID,
		Name:        name,
		Type:        storeType,
		Description: description,
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, store); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StoreService) GetStoreByID(ctx context.Context, id uuid.UUID) (*models.Store, error) {
	return s.repo.GetByID(ctx, id)
}
