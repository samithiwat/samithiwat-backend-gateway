package handler

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

type UserContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID(*int32) error
	PaginationQueryParam(*dto.PaginationQueryParams) error
}

// FindAll is a function that get all users in database
// @Summary Get all users
// @Description Return the arrays of user dto if successfully
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} proto.User
// @Failure 400 {object} dto.ResponseErr Invalid query param
// @Failure 503 {object} dto.ResponseErr Service is down
// @Router /user [get]
func (h *UserHandler) FindAll(c UserContext) {
	query := dto.PaginationQueryParams{}

	err := c.PaginationQueryParam(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    []string{"Cannot parse query param"},
		})
		return
	}

	res, err := h.service.FindAll(query)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    []string{"Service is down"},
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Data:       res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, res.Data)
	return
}

// FindOne is a function that get the specific users with id
// @Summary Get specific user with id
// @Description Return the user dto if successfully
// @Param id path int true "id"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} proto.User
// @Failure 400 {object} dto.ResponseErr Invalid ID
// @Failure 404 {object} dto.ResponseErr "Not found user"
// @Failure 503 {object} dto.ResponseErr Service is down
// @Router /user/{id} [get]
func (h *UserHandler) FindOne(c UserContext) {
	var id int32
	err := c.ID(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    []string{"Invalid ID"},
		})
		return
	}

	res, err := h.service.FindOne(id)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    []string{"Service is down"},
		})
		return
	}

	c.JSON(int(res.StatusCode), res)
	return
}

// Create is a function that create the user
// @Summary Create the user
// @Description Return the user dto if successfully
// @Param user body dto.UserDto true "user dto"
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} proto.User
// @Failure 400 {object} dto.ResponseErr Invalid ID
// @Failure 404 {object} dto.ResponseErr Not found user
// @Failure 503 {object} dto.ResponseErr Service is down
// @Router /user [post]
func (h *UserHandler) Create(c UserContext) {
	userDto := dto.UserDto{}
	err := c.Bind(&userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    []string{"Cannot parse user dto"},
		})
		return
	}

	errors := dto.ValidateUser(userDto)
	if errors != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    []string{"Invalid body request"},
			Data:       errors,
		})
		return
	}

	res, err := h.service.Create(userDto)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    []string{"Service is down"},
		})
		return
	}

	c.JSON(int(res.StatusCode), res.Data)
	return
}

// Update is a function that update the user
// @Summary Update the existing user
// @Description Return the user dto if successfully
// @Param id path int true "id"
// @Param user body proto.User true "user dto"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} proto.User
// @Failure 400 {object} dto.ResponseErr Invalid ID
// @Failure 404 {object} dto.ResponseErr Not found user
// @Failure 503 {object} dto.ResponseErr Service is down
// @Router /user/{id} [patch]
func (h *UserHandler) Update(c UserContext) {
	userDto := dto.UserDto{}
	err := c.Bind(&userDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    []string{"Cannot parse user dto"},
		})
		return
	}

	errors := dto.ValidateUser(userDto)
	if errors != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    []string{"Invalid body request"},
			Data:       errors,
		})
		return
	}

	var id int32
	err = c.ID(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    []string{"Invalid ID"},
		})
		return
	}

	res, err := h.service.Update(id, userDto)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    []string{"Service is down"},
		})
		return
	}

	c.JSON(int(res.StatusCode), res.Data)
	return
}

// Delete is a function that delete the user
// @Summary Delete the user
// @Description Return the user dto if successfully
// @Param id path int true "id"
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} proto.User
// @Failure 400 {object} dto.ResponseErr Invalid ID
// @Failure 404 {object} dto.ResponseErr Not found user
// @Failure 503 {object} dto.ResponseErr Service is down
// @Router /user/{id} [delete]
func (h *UserHandler) Delete(c UserContext) {
	var id int32
	err := c.ID(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    []string{"Invalid ID"},
		})
		return
	}

	res, err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    []string{"Service is down"},
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    res.Errors,
		})
		return
	}

	c.JSON(int(res.StatusCode), res.Data)
	return
}
