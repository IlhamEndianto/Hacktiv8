package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// var (
// 	argAppName = kingpin.Arg("name", "Application name").Required().String()
// 	argPort    = kingpin.Arg("port", "Web server port").Default("9000").Int()
// )

func main() {
	// kingpin.Parse()
	viper.SetConfigName("./config") // name of config file (without extension)
	viper.SetConfigType("json")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	e := echo.New()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	// e.Use(middlewareOne)
	// e.Use(middlewareTwo)
	// e.Use(echo.WrapMiddleware(middlewareSomething))

	e.Use(middlewareLogging)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method= ${method}, uri=${uri}, status=${status}\n",
	}))
	e.HTTPErrorHandler = errorHandler

	e.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("threeeeeee!")

		return c.JSON(http.StatusOK, true)
	})
	e.GET("/private", func(c echo.Context) (err error) {
		fmt.Println("threeeeeee!")

		return c.JSON(http.StatusOK, true)
	})
	private := e.Group("/private")
	private.Use(middleware.BasicAuth(BasicAuthMiddleware))
	private.GET("/index", func(c echo.Context) (err error) {
		fmt.Println("threeeeee!")

		return c.JSON(http.StatusOK, true)
	})
	e.Logger.Fatal(e.Start(":9000"))

	// lock := make(chan error)
	// go func(lock chan error) {
	// 	lock <- e.Start(":9000")
	// }(lock)

	// time.Sleep(1 * time.Millisecond)
	// makeLogEntry(nil).Warning("application started without ssl/tls enabled")

	// err := <-lock
	// if err != nil {
	// 	makeLogEntry(nil).Panic("failed to start application")
	// }
}

func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware one")
		return next(c)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("from middleware two")
		return next(c)
	}
}

func middlewareSomething(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("from middleware something")
		next.ServeHTTP(w, r)
	})
}

func makeLogEntry(c echo.Context) *log.Entry {
	if c == nil {
		return log.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})

	}
	return log.WithFields(log.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"ip":     c.Request().RemoteAddr,
	})
}

func middlewareLogging(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if ok {
		report.Message = fmt.Sprintf("http error %d - %v", report.Code, report.Message)
	} else {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	makeLogEntry(c).Error(report.Message)
	c.HTML(report.Code, report.Message.(string))
}
func BasicAuthMiddleware(username, password string, c echo.Context) (bool, error) {
	// Be careful to use constant time comparison to prevent timing attacks
	if username == viper.GetString("basicAuthUser") && password == viper.GetString("basicAuthPass") {
		return true, nil
	}
	return false, nil
}
