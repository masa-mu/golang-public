package common

import (
	"fmt"
	"github.com/labstack/echo"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) BindValidate(p interface{}) (error) {

	fmt.Println("test1",p)

	if err := c.Bind(p); err != nil {
		return err
	}

	fmt.Println("test2",p)

	if err := c.Validate(p); err != nil {
		return err
	}

	return nil
}
