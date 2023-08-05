// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/phamphihungbk/go-graphql-api/config"
	"github.com/phamphihungbk/go-graphql-api/internal/app"
	"github.com/phamphihungbk/go-graphql-api/internal/repository"
	"github.com/phamphihungbk/go-graphql-api/internal/resolver"
	"github.com/phamphihungbk/go-graphql-api/internal/service"
)

// Injectors from wire.go:

func InitializeApp(dbCfg *config.DBCfg) (*Application, error) {
	db := app.NewDBConnection(dbCfg)
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	resolverResolver := resolver.NewResolver(userService)
	engine := app.NewRoute(resolverResolver)
	application := NewApplication(engine)
	return application, nil
}
