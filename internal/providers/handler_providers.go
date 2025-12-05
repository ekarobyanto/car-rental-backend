package providers

import (
	"car-rental-backend/internal/handler"

	"github.com/google/wire"
)

var HandlerSet = wire.NewSet(
	handler.NewAuthHandler,
	handler.NewCarHandler,
	handler.NewRenterHandler,
	handler.NewRentalTransactionHandler,
)
