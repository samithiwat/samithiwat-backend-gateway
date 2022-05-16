package middleware

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"net/http"
	"strconv"
)

type AuthGuard struct {
	service handler.AuthService
}

type AuthContext interface {
	GetToken() string
	SetHeader(string, string)
	JSON(int, interface{})
	Next()
}

func NewAuthGuard(s handler.AuthService) *AuthGuard {
	return &AuthGuard{
		service: s,
	}
}

func (m *AuthGuard) Validate(ctx AuthContext) {
	token := ctx.GetToken()
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, &dto.ResponseErr{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token",
		})
		return
	}

	userId, errRes := m.service.Validate(token)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.SetHeader("UserId", strconv.Itoa(int(userId)))
	ctx.Next()
}
