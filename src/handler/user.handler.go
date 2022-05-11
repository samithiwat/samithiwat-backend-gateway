package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	validate "github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"net/http"
)

type UserHandler struct {
	service  UserService
	validate *validator.Validate
}

func NewUserHandler(service UserService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validate,
	}
}

type UserContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID(*int32) error
	PaginationQueryParam(*dto.PaginationQueryParams) error
}

type UserService interface {
	FindAll(*dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr)
	FindOne(int32) (*proto.User, *dto.ResponseErr)
	Create(*dto.UserDto) (*proto.User, *dto.ResponseErr)
	Update(int32, *dto.UserDto) (*proto.User, *dto.ResponseErr)
	Delete(int32) (*proto.User, *dto.ResponseErr)
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
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Cannot parse query param",
		})
		return
	}

	users, errRes := h.service.FindAll(&query)
	if users.Items == nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, users)
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
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	user, errRes := h.service.FindOne(id)
	if user.Id == 0 {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, user)
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
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse user dto",
		})
		return
	}

	if errors := h.validate.Struct(userDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       validate.Format(errors.(validator.ValidationErrors)),
		})
		return
	}

	user, errRes := h.service.Create(&userDto)
	if user.Id == 0 {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, user)
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
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse user dto",
		})
		return
	}

	if errors := h.validate.Struct(userDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       validate.Format(errors.(validator.ValidationErrors)),
		})
		return
	}

	var id int32
	err = c.ID(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	user, errRes := h.service.Update(id, &userDto)
	if user.Id == 0 {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, user)
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
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	user, errRes := h.service.Delete(id)
	if user.Id == 0 {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
