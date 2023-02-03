package driver

import (
	"Hacktiv8project/assignment-2/config"
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, config.PostgresAddress)
}
