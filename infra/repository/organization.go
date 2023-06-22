package repository

import (
	"context"
	"fmt"

	domainmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/marshaller"
	recordmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
)

func (r *Repository) CreateOrganization(ctx context.Context, model *domainmodel.Organization) (*domainmodel.Organization, error) {
	record, err := r.queries.CreateOrganization(ctx, &recordmodel.CreateOrganizationParams{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create organization: %w", err)
	}

	return marshaller.RecordToOrganization(record), nil
}
