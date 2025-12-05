//go:build wireinject

//go:generate wire

package main

import (
	"car-rental-backend/config"
	"car-rental-backend/internal/providers"
	"car-rental-backend/pkg/database"
	"car-rental-backend/pkg/jwt"
	"car-rental-backend/pkg/minio"

	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*App, error) {
	wire.Build(
		database.NewDatabase,
		minio.NewMinioClient,
		jwt.NewJWTManager,
		providers.RepositorySet,
		providers.ServiceSet,
		providers.HandlerSet,
		NewApp,
	)
	return nil, nil
}
