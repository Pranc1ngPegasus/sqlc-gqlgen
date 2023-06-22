package repository

import (
	domain "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/repository"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
	"github.com/google/wire"
)

var _ domain.Repository = (*Repository)(nil)

var NewRepositorySet = wire.NewSet(
	wire.Bind(new(domain.Repository), new(*Repository)),
	NewRepository,
)

type Repository struct {
	queries *model.Queries
}

func NewRepository(
	queries *model.Queries,
) *Repository {
	return &Repository{
		queries: queries,
	}
}
