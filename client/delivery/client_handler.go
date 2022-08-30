package delivery

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nelsonomoi/invoice/domain"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ClientHandler struct {
	CUsecase domain.ClientUseCase
}

func NewClientHandler(e *echo.Echo, us domain.ClientUseCase) {
	handler := &ClientHandler{
		CUsecase: us,
	}
	// e.GET("/clients", handler.FetchArticle)
	e.POST("/client", handler.Store)
}

func isRequestValid(client *domain.Client) {
	panic("unimplemented")
}

func (a *ClientHandler) Store(c echo.Context) (err error) {
	var client domain.Client
	err = c.Bind(&client)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	ctx := c.Request().Context()
	err = a.CUsecase.Store(ctx, &client)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, client)
}


func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
