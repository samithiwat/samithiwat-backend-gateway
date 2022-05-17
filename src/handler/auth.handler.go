package handler

import (
	"fmt"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	validate "github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"net/http"
)

type AuthHandler struct {
	service  AuthService
	userSrv  UserService
	validate *validate.DtoValidator
}

func NewAuthHandler(s AuthService, u UserService, v *validate.DtoValidator) *AuthHandler {
	return &AuthHandler{
		service:  s,
		validate: v,
		userSrv:  u,
	}
}

type AuthContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	UserID() int32
}

type AuthService interface {
	Register(*dto.Register) (*proto.User, *dto.ResponseErr)
	Login(*dto.Login) (*proto.Credential, *dto.ResponseErr)
	Logout(uint32) (bool, *dto.ResponseErr)
	ChangePassword(*dto.ChangePassword) (bool, *dto.ResponseErr)
	Validate(string) (uint32, *dto.ResponseErr)
	RefreshToken(string) (*proto.Credential, *dto.ResponseErr)
}

// Register is a function that register user account
// @Summary Register user account
// @Description Return the user dto if successfully
// @Param register body dto.Register true "register dto"
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} proto.User
// @Failure 400 {object} dto.ResponseErr "Invalid request body"
// @Failure 422 {object} dto.ResponseErr "Email is already existed"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /auth/register [post]
func (h *AuthHandler) Register(c AuthContext) {
	register := dto.Register{}
	err := c.Bind(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse register dto",
		})
		return
	}

	if errors := h.validate.Validate(register); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	res, errRes := h.service.Register(&register)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, res)
	return
}

// Login is a function that login the user
// @Summary Login user account
// @Description Return the credentials if successfully
// @Param register body dto.Login true "login dto"
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} proto.Credential
// @Failure 400 {object} dto.ResponseErr "Invalid request body"
// @Failure 401 {object} dto.ResponseErr "Invalid email or username"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c AuthContext) {
	login := dto.Login{}
	err := c.Bind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse login dto",
		})
		return
	}

	if errors := h.validate.Validate(login); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	res, errRes := h.service.Login(&login)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// Logout is a function log out from service
// @Summary Logout user from service
// @Description Return the user dto if successfully
// @Tags auth
// @Accept json
// @Produce json
// @Success 204
// @Failure 401 {object} dto.ResponseErr "Invalid token"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /auth/logout [get]
func (h *AuthHandler) Logout(c AuthContext) {
	userId := c.UserID()

	res, errRes := h.service.Logout(uint32(userId))
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// ChangePassword is a function that change password of the user account
// @Summary ChangePassword of user account
// @Description Return the true if successfully
// @Param register body dto.ChangePassword true "change password dto"
// @Tags auth
// @Accept json
// @Produce json
// @Success 204
// @Failure 400 {object} dto.ResponseErr "Invalid request body"
// @Failure 401 {object} dto.ResponseErr "Invalid access token"
// @Failure 403 {object} dto.ResponseErr "Insufficiency permission"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /auth/change-password [post]
func (h *AuthHandler) ChangePassword(c AuthContext) {
	changePassword := dto.ChangePassword{}
	err := c.Bind(&changePassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse changePassword dto",
		})
		return
	}

	if errors := h.validate.Validate(changePassword); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	res, errRes := h.service.ChangePassword(&changePassword)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusNoContent, res)
	return
}

// Validate is a function check the user token and return user dto
// @Summary Check user status and user info
// @Description Return the user dto if successfully
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} proto.User
// @Failure 401 {object} dto.ResponseErr "Invalid token"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /auth/me [get]
func (h *AuthHandler) Validate(c AuthContext) {
	id := c.UserID()

	fmt.Println(id)

	res, errRes := h.userSrv.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

// RefreshToken is a function that redeem new credentials
// @Summary Redeem new token
// @Description Return the credentials if successfully
// @Param register body dto.RedeemNewToken true "refresh token dto"
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} proto.Credential
// @Failure 400 {object} dto.ResponseErr "Invalid request body"
// @Failure 401 {object} dto.ResponseErr "Invalid refresh token"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /auth/token [post]
func (h *AuthHandler) RefreshToken(c AuthContext) {
	redeemNewToken := dto.RedeemNewToken{}
	err := c.Bind(&redeemNewToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse refresh token dto",
		})
		return
	}

	if errors := h.validate.Validate(redeemNewToken); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	res, errRes := h.service.RefreshToken(redeemNewToken.RefreshToken)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
