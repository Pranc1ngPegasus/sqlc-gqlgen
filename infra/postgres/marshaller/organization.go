package marshaller

import (
	domainmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	recordmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
)

func OrganizationToRecord(model *domainmodel.Organization) *recordmodel.Organization {
	return &recordmodel.Organization{
		ID:        model.ID,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Name:      model.Name,
	}
}

func RecordToOrganization(record *recordmodel.Organization) *domainmodel.Organization {
	return &domainmodel.Organization{
		ID:        record.ID,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		Name:      record.Name,
	}
}
