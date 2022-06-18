package handler

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	validate "github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"net/http"
)

type BlogUserHandler struct {
	service  BlogUserService
	validate *validate.DtoValidator
}

func NewBlogUserHandler(service BlogUserService, validate *validate.DtoValidator) *BlogUserHandler {
	return &BlogUserHandler{
		service:  service,
		validate: validate,
	}
}

type BlogUserContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID() (int32, error)
	UserID() int32
}

type BlogUserService interface {
	FindOne(int32) (*proto.BlogUser, *dto.ResponseErr)
	Create(*dto.BlogUserDto) (*proto.BlogUser, *dto.ResponseErr)
	Update(int32, *dto.BlogUserDto) (*proto.BlogUser, *dto.ResponseErr)
	Delete(int32) (*proto.BlogUser, *dto.ResponseErr)
	AddBookmark(int32, int32) ([]*uint32, *dto.ResponseErr)
	FindAllBookmark(int32) ([]*uint32, *dto.ResponseErr)
	DeleteBookmark(int32, int32) (bool, *dto.ResponseErr)
	Read(int32, int32) (bool, *dto.ResponseErr)
}

// FindOne is a function that get the specific BUsers with id
// @Summary Get specific BUser with id
// @Description Return the blog user dto if successfully
// @Param id path int true "id"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BlogUser
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /blog/user/{id} [get]
func (h *BlogUserHandler) FindOne(c BlogUserContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	BUser, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, BUser)
	return
}

// Create is a function that create the BUser
// @Summary Create the BUser
// @Description Return the blog user dto if successfully
// @Param BUser body dto.BlogUserDto true "BUser dto"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 201 {object} proto.BlogUser
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user [post]
func (h *BlogUserHandler) Create(c BlogUserContext) {
	BUserDto := dto.BlogUserDto{}
	err := c.Bind(&BUserDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse blog user dto",
		})
		return
	}

	if errors := h.validate.Validate(BUserDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	BUser, errRes := h.service.Create(&BUserDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, BUser)
	return
}

// Update is a function that update the BUser
// @Summary Update the existing BUser
// @Description Return the blog user dto if successfully
// @Param id path int true "id"
// @Param BUser body dto.BlogUserDto true "BUser dto"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BlogUser
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/{id} [patch]
func (h *BlogUserHandler) Update(c BlogUserContext) {
	BUserDto := dto.BlogUserDto{}
	err := c.Bind(&BUserDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse blog user dto",
		})
		return
	}

	if errors := h.validate.Validate(BUserDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	BUser, errRes := h.service.Update(id, &BUserDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, BUser)
	return
}

// Delete is a function that delete the BUser
// @Summary Delete the BUser
// @Description Return the blog user dto if successfully
// @Param id path int true "id"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BlogUser
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/{id} [delete]
func (h *BlogUserHandler) Delete(c BlogUserContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	BUser, errRes := h.service.Delete(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, BUser)
	return
}

// AddBookmark is a function that add blog post to user's bookmark
// @Summary Add blog to user's bookmark
// @Description Return true if successfully
// @Param id path int true "id"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BookmarkResponse
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/bookmark/{id} [post]
func (h *BlogUserHandler) AddBookmark(c BlogUserContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	userId := c.UserID()

	bookmarkList, errRes := h.service.AddBookmark(id, userId)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, bookmarkList)
	return
}

// FindAllBookmark is a function that get all blog post from user's bookmark
// @Summary Find all blog post from user's bookmark
// @Description Return true if successfully
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BookmarkResponse
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/bookmark [get]
func (h *BlogUserHandler) FindAllBookmark(c BlogUserContext) {
	userId := c.UserID()

	BUser, errRes := h.service.FindAllBookmark(userId)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, BUser)
	return
}

// DeleteBookmark is a function that delete blog post from user's bookmark
// @Summary Delete blog post from user's bookmark
// @Description Return true if successfully
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BookmarkStatusResponse
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/bookmark/{id} [delete]
func (h *BlogUserHandler) DeleteBookmark(c BlogUserContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	userId := c.UserID()

	isDelete, errRes := h.service.DeleteBookmark(id, userId)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, isDelete)
	return
}

// Read is a function that update user read blog post
// @Summary Update that user read blog post
// @Description Return true if successfully
// @Param id path int true "id"
// @Tags Blog User
// @Accept json
// @Produce json
// @Success 200 {object} proto.BookmarkResponse
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found User"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /blog/user/read/{id} [post]
func (h *BlogUserHandler) Read(c BlogUserContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	userId := c.UserID()

	isRead, errRes := h.service.Read(id, userId)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, isRead)
	return
}
