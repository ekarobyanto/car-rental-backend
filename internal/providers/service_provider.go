package providers

import (
	"car-rental-backend/internal/service"

	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewMediaService,
	service.NewUserService,
	service.NewCarService,
	service.NewRenterService,
	service.NewRentalTransactionService,
)
