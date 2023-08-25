//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api"
	handler "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/api/handler"
	config "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/config"
	db "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/db"
	repository "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/repository"
	usecase "github.com/aarathyaadhiv/ecommerce-fashionsture-cleanarch.git/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP,
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
		repository.NewProductRepository,
		usecase.NewProductUseCase,
		handler.NewProductHandler,
		repository.NewOtpRepository,
		usecase.NewOtpUseCase,
		handler.NewOtpHandler,
	    repository.NewCartRepository,
		usecase.NewCartUseCase,
		handler.NewCartHandler,
		repository.NewOrderRepository,
		usecase.NewOrderUseCase,
		handler.NewOrderHandler)

	return &http.ServerHTTP{}, nil
}
