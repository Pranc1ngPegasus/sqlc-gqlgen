package marshaller

import (
	domainmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	recordmodel "github.com/Pranc1ngPegasus/sqlc-gqlgen/infra/postgres/model"
)

func UserToRecord(model *domainmodel.User) *recordmodel.User {
	return &recordmodel.User{
		ID:             model.ID,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
		OrganizationID: model.OrganizationID,
		Name:           model.Name,
	}
}

func RecordToUser(record *recordmodel.User) *domainmodel.User {
	return &domainmodel.User{
		ID:             record.ID,
		CreatedAt:      record.CreatedAt,
		UpdatedAt:      record.UpdatedAt,
		OrganizationID: record.OrganizationID,
		Name:           record.Name,
	}
}
