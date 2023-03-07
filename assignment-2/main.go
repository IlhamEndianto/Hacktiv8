package main

import (
	"github.com/IlhamEndianto/Hacktiv8/assignment-2/config"
	"github.com/IlhamEndianto/Hacktiv8/assignment-2/driver"
	"github.com/IlhamEndianto/Hacktiv8/assignment-2/handler"
	"github.com/IlhamEndianto/Hacktiv8/assignment-2/repository"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
)

const SESSION_ID = "id"

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	e := echo.New()
	e.Renderer = driver.NewRenderer("template/*", true)
	ctx := context.Background()
	store := driver.NewPostgresStore()
	pgPool, err := driver.NewPostgresConn(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUser(pgPool)

	loginHandler := handler.NewLoginHandler(nil, store, userRepo)
	registerHandler := handler.NewRegisterHandler(nil, store, userRepo)

	e.GET("/", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/home")
	})

	e.GET("/login", loginHandler.LoginHandler)
	e.POST("/login", loginHandler.LoginHandler)

	e.GET("/home", loginHandler.HomeHandler)
	e.POST("/home", loginHandler.HomeHandler)

	e.POST("/register", registerHandler.RegisterHandler)

	e.POST("/logout", loginHandler.LogoutHandler)

	e.Start(config.WebServerPort)
}
