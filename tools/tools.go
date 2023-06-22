//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/k0kubun/sqldef/cmd/psqldef"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
)
