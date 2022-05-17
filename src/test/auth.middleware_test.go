package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/middleware"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
)

type AuthGuardTest struct {
	suite.Suite
	ExcludePath     map[string]struct{}
	UserId          int32
	Token           string
	UnauthorizedErr *dto.ResponseErr
	ServiceDownErr  *dto.ResponseErr
}

func TestAuthGuard(t *testing.T) {
	suite.Run(t, new(AuthGuardTest))
}

func (u *AuthGuardTest) SetupTest() {
	u.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	u.UnauthorizedErr = &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid token",
		Data:       nil,
	}

	u.Token = faker.Word()
	u.UserId = int32(rand.Intn(100))

	u.ExcludePath = map[string]struct{}{
		"POST /exclude":     {},
		"POST /exclude/:id": {},
	}
}

func (u *AuthGuardTest) TestValidateSuccess() {
	want := u.UserId

	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/auth")
	c.On("Token").Return(u.Token)
	srv.On("Validate", u.Token).Return(int(u.UserId), nil)
	c.On("SetHeader", "UserId", strconv.Itoa(int(u.UserId)))
	c.On("Next")

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	id, err := strconv.Atoi(c.Header["UserId"])

	assert.Nil(u.T(), err, "Invalid user id")
	assert.Equal(u.T(), want, int32(id))
	c.AssertNumberOfCalls(u.T(), "Next", 1)
}

func (u *AuthGuardTest) TestValidateSkippedFromExcludePath() {
	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/exclude")
	c.On("Token").Return("")
	c.On("Next")

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	c.AssertNumberOfCalls(u.T(), "Next", 1)
	c.AssertNumberOfCalls(u.T(), "Token", 0)
}

func (u *AuthGuardTest) TestValidateSkippedFromExcludePathWithID() {
	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/exclude/1")
	c.On("Token").Return("")
	c.On("Next")

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	c.AssertNumberOfCalls(u.T(), "Next", 1)
	c.AssertNumberOfCalls(u.T(), "Token", 0)
}

func (u *AuthGuardTest) TestValidateFailed() {
	want := u.UnauthorizedErr

	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/auth")
	c.On("Token").Return(u.Token)
	srv.On("Validate", u.Token).Return(-1, u.UnauthorizedErr)

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthGuardTest) TestValidateTokenNotIncluded() {
	want := u.UnauthorizedErr

	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/auth")
	c.On("Token").Return("")
	srv.On("Validate")

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
	srv.AssertNumberOfCalls(u.T(), "Validate", 0)

}

func (u *AuthGuardTest) TestValidateTokenGrpcErr() {
	want := u.ServiceDownErr

	srv := new(mock.AuthServiceMock)
	c := new(mock.AuthContextMock)

	c.On("Method").Return("POST")
	c.On("Path").Return("/auth")
	c.On("Token").Return(u.Token)
	srv.On("Validate", u.Token).Return(-1, u.ServiceDownErr)

	h := middleware.NewAuthGuard(srv, u.ExcludePath)
	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
}
