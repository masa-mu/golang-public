package main

import (
	"github.com/labstack/echo"
	"golang-public/api-ddd-model/message"
	"golang-public/api-ddd-model/mysql"
)

func App() (*echo.Echo){
	messageDB,err := mysql.GetConnectionDB()
	if err != nil {
		panic(err)
	}

	messageHandler := message.NewHandler(message.NewModel(message.NewRepository(messageDB)))

	e := echo.New()
	g := e.Group("/ddd/v1")
	g.POST("/message", messageHandler.GetMessage)

	return e
}

func main() {
	e := App()

	e.Start(":8000")
}
