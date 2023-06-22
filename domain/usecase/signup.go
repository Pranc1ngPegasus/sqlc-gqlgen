package usecase

import "context"

type Signup interface {
	Do(context.Context, SignupInput) error
}

type SignupInput struct {
	OrganizationName string
	UserName         string
}
