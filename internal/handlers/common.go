package handlers

import "github.com/labstack/echo"

type errResponse struct {
	Message string `json:"msg"`
}

func notOk(c echo.Context, status int, err error) error {
	c.Logger().Error(err)
	return c.JSON(status, &errResponse{
		Message: err.Error(),
	})
}
