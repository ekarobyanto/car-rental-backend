package providers

import (
	"car-rental-backend/internal/repository"

	"github.com/google/wire"
)

var RepositorySet = wire.NewSet(
	repository.NewUserRepository,
	repository.NewCarRepository,
	repository.NewRenterRepository,
	repository.NewRentalTransactionRepository,
)
