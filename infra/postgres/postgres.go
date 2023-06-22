package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"

	"github.com/google/wire"
	_ "github.com/lib/pq"
)

var NewPostgresSet = wire.NewSet(
	NewPostgres,
)

func NewPostgres(
	ctx context.Context,
) (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD")),
		Host:   net.JoinHostPort(os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT")),
		Path:   os.Getenv("DATABASE_NAME"),
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	db, err := sql.Open("postgres", dsn.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make connection for postgres: %w", err)
	}

	return db, nil
}
