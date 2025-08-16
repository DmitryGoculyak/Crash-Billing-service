package service

import (
	"Crash-Billing-service/internal/entities"
	"Crash-Billing-service/internal/repository"
	"context"
)

type BillingServiceServer interface {
	CreateWallet(ctx context.Context, userId, currencyCode string) (*entities.Wallet, error)
	GetWallet(ctx context.Context, userId string) (*entities.Wallet, error)
}

type BillingService struct {
	repo repository.BillingRepository
}

func NewBillingService(repo repository.BillingRepository) BillingServiceServer {
	return &BillingService{repo: repo}
}

func (s *BillingService) CreateWallet(ctx context.Context, userId, currencyCode string) (*entities.Wallet, error) {
	return s.repo.CreateWallet(ctx, userId, currencyCode)
}

func (s *BillingService) GetWallet(ctx context.Context, userId string) (*entities.Wallet, error) {
	return s.repo.GetWallet(ctx, userId)
}
