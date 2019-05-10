package main

import (
	"github.com/go-playground/validator"
	"golang-public/api-ddd-model/message"
	"golang-public/api-ddd-model/mysql"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func App() *echo.Echo {
	messageDB, err := mysql.GetConnectionDB()
	if err != nil {
		panic(err)
	}

	messageHandler := message.NewHandler(message.NewModel(message.NewRepository(messageDB)))

	e := echo.New()

	e.Validator = NewValidator()

	g := e.Group("/ddd/v1")
	g.GET("/message", messageHandler.GetMessage)

	return e
}

func main() {
	e := App()

	e.Start(":8000")
}
