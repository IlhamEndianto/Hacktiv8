package config

const (
	PostgresUsername = "postgresuser"
	PostgresPassword = "postgrespassword"
	PostgresAddress  = "postgres://postgresuser:postgrespassword@postgres:5432/postgres?sslmode=disable"
	WebServerPort    = ":8080"
	SessionId        = "test-id"
)

var (
	SessionAuthKey    = []byte("my-auth-key-very-secret")
	SessionEncryptKey = []byte("my-encryption-key-very-secret123")
)
