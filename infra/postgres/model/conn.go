package model

import (
	"database/sql"

	"github.com/google/wire"
)

var NewConnSet = wire.NewSet(
	wire.Bind(new(DBTX), new(*sql.DB)),
	NewConn,
)

func NewConn(
	db *sql.DB,
) *Queries {
	return New(db)
}
