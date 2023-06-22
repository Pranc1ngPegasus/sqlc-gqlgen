package repository

import (
	"context"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
)

type Organization interface {
	CreateOrganization(context.Context, *model.Organization) (*model.Organization, error)
}
