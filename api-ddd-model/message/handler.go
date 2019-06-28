package message

import (
	"net/http"

	"github.com/labstack/echo"
)

type getMessageForm struct {
	ID int `query:"id" validate:"required"`
}

type responseMessage struct {
	message string `json:"message"`
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

func (h *handler) GetMessage(c echo.Context) (err error) {
	req := new(getMessageForm)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusOK, NewResponse(req))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusOK, NewResponse(req))
	}

	return c.JSON(http.StatusOK, NewResponse(req))
}

func NewResponse(body interface{}) Response {
	return Response{
		Header: ResponseHeader{
			ResultCode: "test",
		},
		Body: body,
	}
}

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Body   interface{}    `json:"body"`
	}
	ResponseHeader struct {
		ResultCode string `json:"result_code"`
	}
)
