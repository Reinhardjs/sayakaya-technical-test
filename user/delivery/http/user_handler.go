package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/reinhardjs/sayakaya/domain"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Code    int           `json:"code"`
	Message string        `json:"msg"`
	Users   []domain.User `json:"records"`
}

type UserHandler struct {
	Usecase domain.UserUsecase
}

func NewUserHandler(e *echo.Echo, us domain.UserUsecase) {
	handler := &UserHandler{
		Usecase: us,
	}
	e.GET("/users", handler.FetchUser)
	e.POST("/users", handler.Store)
	e.GET("/users/:id", handler.GetByID)
	e.DELETE("/users/:id", handler.Delete)
}

func (a *UserHandler) FetchUser(c echo.Context) error {
	ctx := c.Request().Context()

	listAr, err := a.Usecase.Fetch(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

func (a *UserHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.Usecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}

func isRequestValid(m *domain.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *UserHandler) Store(c echo.Context) (err error) {
	var user domain.User
	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	err = a.Usecase.Store(ctx, &user)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// Delete will delete user by given param
func (a *UserHandler) Delete(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	err = a.Usecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
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
