package handlers

import (
	"Crash-Billing-service/internal/service"
	proto "Crash-Billing-service/pkg/proto"
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type BillingHandler struct {
	proto.UnimplementedBillingServiceServer
	service service.BillingServiceServer
	log     *zap.Logger
}

func NewBillingHandler(
	service service.BillingServiceServer,
	log *zap.Logger,
) *BillingHandler {
	return &BillingHandler{
		service: service,
		log:     log,
	}
}

func (h *BillingHandler) CreateWallet(ctx context.Context, req *proto.CreateWalletRequest) (*proto.WalletResponse, error) {
	createCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	user, err := h.service.CreateWallet(createCtx, req.UserId, req.CurrencyCode)
	if err != nil {
		h.log.Error("Failed to create wallet",
			zap.String("userId", req.UserId),
			zap.String("currency", req.CurrencyCode),
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Internal, "Failed to create wallet: %v", err)
	}

	h.log.Info("Wallet created successfully",
		zap.String("walletId", user.Id),
	)
	return &proto.WalletResponse{
		Id: user.Id,
	}, nil
}

func (h *BillingHandler) GetWallet(ctx context.Context, req *proto.GetWalletRequest) (*proto.WalletResponse, error) {
	createCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	wallet, err := h.service.GetWallet(createCtx, req.UserId)
	if err != nil {
		h.log.Error("Failed to get wallet",
			zap.String("userId", req.UserId),
			zap.Error(err),
		)
		return nil, status.Errorf(codes.Internal, "Failed to get wallet: %v", err)
	}

	h.log.Info("Get  wallet successfully",
		zap.String("walletId", wallet.Id),
		zap.String("userId", wallet.UserId),
		zap.String("currency", wallet.CurrencyCode),
		zap.Float64("balance", wallet.Balance),
	)
	return &proto.WalletResponse{
		UserId:       wallet.UserId,
		Id:           wallet.Id,
		CurrencyCode: wallet.CurrencyCode,
		Balance:      wallet.Balance,
	}, nil
}
