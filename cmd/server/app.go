package main

import (
	"car-rental-backend/internal/handler"
	"car-rental-backend/pkg/jwt"
	"car-rental-backend/pkg/minio"

	"gorm.io/gorm"
)

type App struct {
	DB                       *gorm.DB
	MinioClient              *minio.MinioClient
	JWTManager               *jwt.JWTManager
	AuthHandler              *handler.AuthHandler
	CarHandler               *handler.CarHandler
	RenterHandler            *handler.RenterHandler
	RentalTransactionHandler *handler.RentalTransactionHandler
}

func NewApp(
	db *gorm.DB,
	minioClient *minio.MinioClient,
	jwtManager *jwt.JWTManager,
	authHandler *handler.AuthHandler,
	carHandler *handler.CarHandler,
	renterHandler *handler.RenterHandler,
	rentalTransactionHandler *handler.RentalTransactionHandler,
) *App {
	return &App{
		DB:                       db,
		MinioClient:              minioClient,
		JWTManager:               jwtManager,
		AuthHandler:              authHandler,
		CarHandler:               carHandler,
		RenterHandler:            renterHandler,
		RentalTransactionHandler: rentalTransactionHandler,
	}
}
