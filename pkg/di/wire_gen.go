// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"jerseyhub/pkg/api"
	"jerseyhub/pkg/api/handler"
	"jerseyhub/pkg/config"
	"jerseyhub/pkg/db"
	"jerseyhub/pkg/repository"
	"jerseyhub/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewAdminRepository(gormDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)

	categoryRepository := repository.NewCategoryRepository(gormDB)
	categoryUseCase := usecase.NewCategoryUseCase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	inventoryRepository := repository.NewInventoryRepository(gormDB)
	inventoryUseCase := usecase.NewInventoryUseCase(inventoryRepository)
	inventoryHandler := handler.NewInventoryHandler(inventoryUseCase)

	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUseCase(cfg, otpRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)

	userRepository := repository.NewUserRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository,cfg,otpRepository,inventoryRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	couponRepository := repository.NewCouponRepository(gormDB)
	couponUseCase := usecase.NewCouponUseCase(couponRepository)
	couponHandler := handler.NewCouponHandler(couponUseCase)



	orderRepository := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepository,couponRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)

	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository,inventoryRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)

	
	serverHTTP := http.NewServerHTTP(userHandler,adminHandler,categoryHandler,inventoryHandler,otpHandler,orderHandler,cartHandler,couponHandler)



	return serverHTTP, nil
}
