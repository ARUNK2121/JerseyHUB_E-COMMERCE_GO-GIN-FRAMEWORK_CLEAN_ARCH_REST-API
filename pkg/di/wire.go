//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "jerseyhub/pkg/api"
	handler "jerseyhub/pkg/api/handler"
	config "jerseyhub/pkg/config"
	db "jerseyhub/pkg/db"
	repository "jerseyhub/pkg/repository"
	usecase "jerseyhub/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserRepository, usecase.NewUserUseCase, handler.NewUserHandler, http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
