package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
	OrganizationID uuid.UUID
	Name           string
}

func NewUser(
	organizationID uuid.UUID,
	name string,
) *User {
	now := time.Now()

	return &User{
		ID:             uuid.New(),
		CreatedAt:      now,
		UpdatedAt:      now,
		OrganizationID: organizationID,
		Name:           name,
	}
}
