package services

import (
	"context"
	"errors"
	"time"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type TenantService struct {
	repo repository.TenantRepository
}

func NewTenantService(repo repository.PostgresTenantRepository) *TenantService {
	return &TenantService{repo: &repo}
}

func (s *TenantService) Create(ctx context.Context, businessName string) (*models.Tenant, error) {
	if businessName == "" {
		return nil, errors.New("business name cannot be empty")
	}

	tenant := &models.Tenant{
		TenantID:     uuid.New(),
		BusinessName: businessName,
		Created_At:   time.Now(),
	}

	err := s.repo.Create(ctx, tenant)
	if err != nil {
		return nil, err
	}

	return tenant, err
}

func (s *TenantService) GetTenantByID(ctx context.Context, id uuid.UUID) (*models.Tenant, error) {
	return s.repo.GetByID(ctx, id)
}
