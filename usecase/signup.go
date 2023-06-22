package usecase

import (
	"context"
	"fmt"

	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/model"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/repository"
	domain "github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/usecase"
	"github.com/google/wire"
)

var _ domain.Signup = (*Signup)(nil)

var NewSignupSet = wire.NewSet(
	wire.Bind(new(domain.Signup), new(*Signup)),
	NewSignup,
)

type Signup struct {
	repository repository.Repository
}

func NewSignup(
	repository repository.Repository,
) *Signup {
	return &Signup{
		repository: repository,
	}
}

func (u *Signup) Do(ctx context.Context, input domain.SignupInput) error {
	newOrganization := model.NewOrganization(input.OrganizationName)

	organization, err := u.repository.CreateOrganization(ctx, newOrganization)
	if err != nil {
		return fmt.Errorf("failed to create organization: %w", err)
	}

	newUser := model.NewUser(organization.ID, input.UserName)

	if _, err := u.repository.CreateUser(ctx, newUser); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
