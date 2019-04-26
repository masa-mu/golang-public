package main

import (
	"golang-public/api-ddd-model/message"
	"golang-public/api-ddd-model/mysql"

	"github.com/labstack/echo"

	_ "github.com/go-sql-driver/mysql"
)

func App() *echo.Echo {
	messageDB, err := mysql.GetConnectionDB()
	if err != nil {
		panic(err)
	}

	messageHandler := message.NewHandler(message.NewModel(message.NewRepository(messageDB)))

	e := echo.New()
	g := e.Group("/ddd/v1")
	g.GET("/message", messageHandler.GetMessage)

	return e
}

func main() {
	e := App()

	e.Start(":8000")
}
