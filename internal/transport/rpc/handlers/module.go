package handlers

import (
	proto "Crash-Billing-service/pkg/proto"
	"go.uber.org/fx"
)

var Module = fx.Module("handlers",
	fx.Provide(
		NewBillingHandler,
		func(h *BillingHandler) proto.BillingServiceServer { return h },
	),
)
