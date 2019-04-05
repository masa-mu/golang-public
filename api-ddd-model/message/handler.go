package message

import (
	"github.com/labstack/echo"
	"net/http"
)

type Handler interface {
	GetMessage(c echo.Context) (err error)
}

type handler struct {
	model Model
}

func NewHandler(model Model) Handler {
	return &handler{model}
}

func (h *handler) GetMessage(c echo.Context) (err error){
	return c.JSON(http.StatusOK,err)
}
