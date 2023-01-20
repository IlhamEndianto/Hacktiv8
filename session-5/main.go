package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type M struct {
	Message string
	Counter int
}

var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("from action index"))
}
var ActionHome = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action home"))
	},
)
var ActionAbout = echo.WrapHandler(
	http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action about"))
		},
	),
)

type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=0,lte=80"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	r := echo.New()
	r.Validator = &CustomValidator{validator: validator.New()}
	// r.GET("/", func(ctx echo.Context) error {
	// 	data := "hello from /index"
	// 	return ctx.String(http.StatusOK, data)
	// })

	// r.GET("/index", func(ctx echo.Context) error {
	// 	data := "Hello from /index"
	// 	return ctx.String(http.StatusOK, data)
	// })
	// r.GET("/html", func(ctx echo.Context) error {
	// 	data := "Hello from /html"
	// 	return ctx.HTML(http.StatusOK, data)
	// })
	// r.GET("/index", func(ctx echo.Context) error {
	// 	return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	// })
	// r.GET("/json", func(ctx echo.Context) error {
	// 	data := M{Message: "Hello", Counter: 2}
	// 	return ctx.JSON(http.StatusOK, data)
	// })
	// r.GET("/page1", func(ctx echo.Context) error {
	// 	name := ctx.QueryParam("name")
	// 	data := fmt.Sprintf("Hello %s", name)
	// 	return ctx.HTML(http.StatusOK, data)
	// })
	// r.GET("/page2/:name", func(ctx echo.Context) error {
	// 	name := ctx.Param("name")
	// 	data := fmt.Sprintf("Hello %s", name)
	// 	return ctx.HTML(http.StatusOK, data)
	// })
	// r.GET("/page3/:name/*", func(ctx echo.Context) error {
	// 	name := ctx.Param("name")
	// 	message := ctx.Param("*")
	// 	data := fmt.Sprintf("Hello %s, I have message for you: %s", name, message)
	// 	return ctx.HTML(http.StatusOK, data)
	// })
	// r.GET("/page4/", func(ctx echo.Context) error {
	// 	name := ctx.FormValue("name")
	// 	message := ctx.FormValue("message")
	// 	data := fmt.Sprintf("Hello %s, I have message for you: %s", name, strings.Replace(message, "/", "", 1))
	// 	return ctx.String(http.StatusOK, data)
	// })

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	r.Static("/static", "assets")

	r.Any("/user", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}
		return c.JSON(http.StatusOK, u)
	})

	r.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, true)
	})

	r.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// if castedObject, ok := err.(validator.ValidationErrors); ok {
		// 	for _, err := range castedObject {
		// 		switch err.Tag() {
		// 		case "required":
		// 			report.Message = fmt.Sprintf("%s is required", err.Field())
		// 		case "email":
		// 			report.Message = fmt.Sprintf("%s is not a valid email address", err.Field())
		// 		case "gte":
		// 			report.Message = fmt.Sprintf("%s value must be greater than or equal %s", err.Field(), err.Param())
		// 		case "lte":
		// 			report.Message = fmt.Sprintf("%s value must be less than or equal %s", err.Field(), err.Param())
		// 		}
		// 	}
		// }

		// c.Logger().Error(report)
		// c.JSON(report.Code, report)

		errPage := fmt.Sprintf("%d.html", report.Code)
		if err := c.File(errPage); err != nil {
			c.HTML(report.Code, "Errrooooorrrr!!!")
		}
	}

	r.Start("localhost:9000")
}
