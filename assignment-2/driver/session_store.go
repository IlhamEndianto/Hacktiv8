package driver

import (
	"log"

	"github.com/IlhamEndianto/Hacktiv8/assignment-2/config"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
)

func NewPostgresStore() *pgstore.PGStore {
	store, err := pgstore.NewPGStore(config.PostgresAddress, config.SessionAuthKey, config.SessionEncryptKey)
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	return store
}

func NewCookieStore() *sessions.CookieStore {
	store := sessions.NewCookieStore(config.SessionAuthKey, config.SessionEncryptKey)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}
