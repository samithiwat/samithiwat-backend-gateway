package middleware

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/common"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"net/http"
	"strconv"
)

type AuthGuard struct {
	service  handler.AuthService
	excludes map[string]struct{}
}

type AuthContext interface {
	Token() string
	Path() string
	SetHeader(string, string)
	JSON(int, interface{})
	Next()
}

func NewAuthGuard(s handler.AuthService, e map[string]struct{}) AuthGuard {
	return AuthGuard{
		service:  s,
		excludes: e,
	}
}

func (m *AuthGuard) Validate(ctx AuthContext) {
	path := ctx.Path()

	var id int32
	ids := common.FindIntFromStr(path)
	if len(ids) > 0 {
		id = ids[0]
	}

	path = common.FormatPathID(path, id)
	if common.IsExisted(m.excludes, path) {
		ctx.Next()
		return
	}

	token := ctx.Token()
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
