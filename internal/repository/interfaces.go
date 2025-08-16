package repository

import (
	"Crash-Billing-service/internal/entities"
	"context"
)

type BillingRepository interface {
	CreateWallet(ctx context.Context, userId, currencyCode string) (*entities.Wallet, error)
	GetWallet(ctx context.Context, userId string) (*entities.Wallet, error)
}
