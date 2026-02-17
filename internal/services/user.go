package services

import (
	"context"
	"errors"
	"time"

	"github.com/abdul-burale/multi-tenant-order-management-api/internal/models"
	"github.com/abdul-burale/multi-tenant-order-management-api/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo repository.UsersRespository
}

func NewUserService(repo repository.PostgresUserRepository) *UserService {
	return &UserService{repo: &repo}
}

func (s *UserService) CreateUser(ctx context.Context, tenantID uuid.UUID, email, name string, role models.UserRole) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	user := &models.User{
		UserID:    uuid.New(),
		TenantID:  tenantID,
		Email:     email,
		Name:      name,
		Role:      role,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	return s.GetUserById(ctx, id)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.GetByEmail(ctx, email)
}
