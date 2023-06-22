package repository

import (
	"context"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
)

type User interface {
	CreateUser(context.Context, *model.User) (*model.User, error)
}
