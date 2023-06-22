package usecase

import (
	"context"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
)

type ListOrganizations interface {
	Do(context.Context) ([]*model.Organization, error)
}
