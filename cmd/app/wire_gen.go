// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/handler"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/resolver"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/server"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/repository"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/usecase"
	"net/http"
)

// Injectors from wire.go:

func initialize() (*app, error) {
	contextContext := context.Background()
	db, err := postgres.NewPostgres(contextContext)
	if err != nil {
		return nil, err
	}
	queries := model.NewConn(db)
	repositoryRepository := repository.NewRepository(queries)
	signup := usecase.NewSignup(repositoryRepository)
	listOrganizations := usecase.NewListOrganizations(repositoryRepository)
	executableSchema := resolver.NewSchema(signup, listOrganizations)
	handlerHandler := handler.NewHandler(executableSchema)
	httpServer := server.NewServer(handlerHandler)
	mainApp := &app{
		server: httpServer,
	}
	return mainApp, nil
}

// wire.go:

type app struct {
	server *http.Server
}
