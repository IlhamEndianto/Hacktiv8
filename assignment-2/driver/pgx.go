package driver

import (
	"context"

	"github.com/IlhamEndianto/Hacktiv8/assignment-2/config"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, config.PostgresAddress)
}
