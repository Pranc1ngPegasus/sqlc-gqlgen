package usecase

import (
	"context"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/repository"
	domain "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/usecase"
	"github.com/google/wire"
)

var _ domain.ListOrganizations = (*ListOrganizations)(nil)

var NewListOrganizationsSet = wire.NewSet(
	wire.Bind(new(domain.ListOrganizations), new(*ListOrganizations)),
	NewListOrganizations,
)

type ListOrganizations struct {
	repository repository.Repository
}

func NewListOrganizations(
	repository repository.Repository,
) *ListOrganizations {
	return &ListOrganizations{
		repository: repository,
	}
}

func (u *ListOrganizations) Do(ctx context.Context) ([]*model.Organization, error) {
	return []*model.Organization{}, nil
}
