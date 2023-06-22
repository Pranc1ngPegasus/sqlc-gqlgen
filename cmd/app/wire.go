//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"net/http"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/handler"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/resolver"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/adapter/server"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres"
	sqlc "github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/repository"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/usecase"
	"github.com/google/wire"
)

type app struct {
	server *http.Server
}

func initialize() (*app, error) {
	wire.Build(
		context.Background,
		handler.NewHandlerSet,
		postgres.NewPostgresSet,
		repository.NewRepositorySet,
		resolver.NewSchema,
		server.NewServer,
		sqlc.NewConn,
		usecase.NewSignupSet,
		usecase.NewListOrganizationsSet,

		wire.Struct(new(app), "*"),
	)

	return nil, nil
}
