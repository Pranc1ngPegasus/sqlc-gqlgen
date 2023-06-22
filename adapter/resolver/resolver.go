package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/domain/usecase"
	"github.com/Pranc1ngPegasus/sqlc-gqlgen/graph"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	signup           usecase.Signup
	listOrganization usecase.ListOrganizations
}

func NewSchema(
	signup usecase.Signup,
	listOrganization usecase.ListOrganizations,
) graphql.ExecutableSchema {
	return graph.NewExecutableSchema(
		graph.Config{
			Resolvers: &Resolver{
				signup:           signup,
				listOrganization: listOrganization,
			},
		},
	)
}
