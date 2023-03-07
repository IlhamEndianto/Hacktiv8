package repository

import (
	"context"

	"github.com/IlhamEndianto/Hacktiv8/assignment-2/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	log "github.com/sirupsen/logrus"
)

// PgxPoolIface defines a little interface for pgxpool functionality.
// Since in the real implementation we can use pgxpool.Pool,
// this interface exists mostly for testing purpose.
type PgxPoolIface interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Ping(ctx context.Context) error
}

// User is responsible to connect user entity with users table in PostgreSQL.
type User struct {
	pool PgxPoolIface
}

// NewUser creates an instance of User.
func NewUser(pool PgxPoolIface) *User {
	return &User{pool: pool}
}

// Insert inserts the user into the users table and return the user id.
func (u *User) Insert(ctx context.Context, user *entity.User) error {
	query := "INSERT INTO " +
		"users (username, first_name, last_name, password) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	row := u.pool.QueryRow(ctx, query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Password,
	)

	if err := row.Scan(&user.ID); err != nil {
		log.Warn("Error when scan rows")
		return err
	}

	return nil
}

// GetByUsernamePassword gets a user from PostgreSQL.
// If there isn't any data, it returns error.
func (u *User) GetByUsernamePassword(ctx context.Context, username, password string) (*entity.User, error) {
	log.Info("Start GetByUsernamePassword", username, password)
	query := "SELECT id,username,first_name,last_name,password " +
		"FROM users WHERE username = $1 AND password = $2"

	row := u.pool.QueryRow(ctx, query, username, password)

	user := entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		log.WithError(err).Warn("Error when scan rows")
		return nil, err
	}
	log.Infof("Success retrieve user %+v\n", user)
	return &user, nil
}
