package zap_db

import (
	"context"
	"fmt"
	"os"

	tools "zap/internal/_shared"

	"github.com/jackc/pgx/v5"
)

func New() *pgx.Conn {

	conn, err := pgx.Connect(
		context.Background(),
		"postgres://postgres:1111@localhost:5432/zap",
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	tools.Logg("Successfully connected to database")

	return conn

}
