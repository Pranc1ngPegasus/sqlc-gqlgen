package repository

import (
	"context"
	"fmt"

	domainmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/marshaller"
	recordmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
)

func (r *Repository) CreateUser(ctx context.Context, model *domainmodel.User) (*domainmodel.User, error) {
	record, err := r.queries.CreateUser(ctx, &recordmodel.CreateUserParams{
		ID:             model.ID,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
		OrganizationID: model.OrganizationID,
		Name:           model.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return marshaller.RecordToUser(record), nil
}
