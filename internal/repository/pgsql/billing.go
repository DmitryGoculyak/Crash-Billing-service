package pgsql

import (
	"Crash-Billing-service/internal/entities"
	"Crash-Billing-service/internal/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type BillingRepo struct {
	db *sqlx.DB
}

func NewBillingRepo(db *sqlx.DB) repository.BillingRepository {
	return &BillingRepo{db: db}
}

func (r *BillingRepo) CreateWallet(ctx context.Context, userId, currencyCode string) (*entities.Wallet, error) {
	var createWallet entities.Wallet
	err := r.db.GetContext(ctx, &createWallet, "INSERT INTO wallets(currency_code, balance) VALUES ($1, $2) RETURNING id",
		currencyCode, userId)
	if err != nil {
		return nil, err
	}
	return &createWallet, nil
}

func (r *BillingRepo) GetWallet(ctx context.Context, userId string) (*entities.Wallet, error) {
	var getWallet entities.Wallet
	err := r.db.GetContext(ctx, &getWallet, "SELECT * FROM wallets WHERE id = $1", userId)
	if err != nil {
		return nil, err
	}
	return &getWallet, nil
}
