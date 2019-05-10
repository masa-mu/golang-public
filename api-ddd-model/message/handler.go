package message

import (
	"github.com/labstack/echo"
	"net/http"
)

type getMessageForm struct {
	ID int `query:"id" validate:"required"`
}

type responseMessage struct {
	message string
}

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
	req := new(getMessageForm)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusOK,responseMessage{message:"Bind Error"})
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusOK,responseMessage{message:"Validate Error"})
	}

	return c.JSON(http.StatusOK,responseMessage{message:"Success"})
}
