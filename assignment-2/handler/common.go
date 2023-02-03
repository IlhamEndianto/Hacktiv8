package handler

import (
	"net/http"

	"Hacktiv8project/assignment-2/config"

	"github.com/antonlindstrom/pgstore"
	"github.com/labstack/echo"
)

func storeSessionHelper(c echo.Context, pgs *pgstore.PGStore, username string) error {
	session, err := pgs.Get(c.Request(), config.SessionId)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	session.Values["username"] = username
	if err = session.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
