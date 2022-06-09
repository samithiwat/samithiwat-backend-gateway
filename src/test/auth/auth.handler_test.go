package auth

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/user"
	"github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type AuthHandlerTest struct {
	suite.Suite
	User            *proto.User
	RegisterDto     *dto.Register
	LoginDto        *dto.Login
	ChangePassword  *dto.ChangePassword
	RefreshToken    *dto.RedeemNewToken
	Credential      *proto.Credential
	BadRequestErr   *dto.ResponseErr
	UnauthorizedErr *dto.ResponseErr
	ServiceDownErr  *dto.ResponseErr
}

func TestAuthHandler(t *testing.T) {
	suite.Run(t, new(AuthHandlerTest))
}

func (u *AuthHandlerTest) SetupTest() {
	u.User = &proto.User{
		Id:          1,
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
	}

	u.RegisterDto = &dto.Register{
		Email:       faker.Email(),
		Password:    faker.Password(),
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
	}

	u.LoginDto = &dto.Login{
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	u.ChangePassword = &dto.ChangePassword{
		UserId:      1,
		OldPassword: faker.Password(),
		NewPassword: faker.Password(),
	}

	u.RefreshToken = &dto.RedeemNewToken{
		RefreshToken: faker.Word(),
	}

	u.Credential = &proto.Credential{
		AccessToken:  faker.Word(),
		RefreshToken: faker.Word(),
		ExpiresIn:    3600,
	}

	u.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	u.UnauthorizedErr = &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Something wrong :(",
		Data:       nil,
	}

	u.BadRequestErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid request body",
		Data:       nil,
	}
}

func (u *AuthHandlerTest) TestRegisterSuccess() {
	want := u.User

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Register", c.RegisterDto).Return(u.User, nil)
	c.On("Bind", &dto.Register{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)
	h.Register(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRegisterErrEmailDuplicated() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Email is already existed",
		Data:       nil,
	}

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Register", c.RegisterDto).Return(nil, want)
	c.On("Bind", &dto.Register{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)
	h.Register(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRegisterInvalidDTO() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse register dto",
		Data:       nil,
	}

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Register", c.RegisterDto).Return(nil, want)
	c.On("Bind", &dto.Register{}).Return(errors.New("Malformed data"))

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)
	h.Register(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRegisterGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Register", c.RegisterDto).Return(nil, u.ServiceDownErr)
	c.On("Bind", &dto.Register{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Register(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLoginSuccess() {
	want := u.Credential

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Login", c.LoginDto).Return(u.Credential, nil)
	c.On("Bind", &dto.Login{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Login(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLoginUnAuthorizeErr() {
	want := u.UnauthorizedErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Login", c.LoginDto).Return(nil, u.UnauthorizedErr)
	c.On("Bind", &dto.Login{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Login(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLoginInvalidDTO() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse login dto",
		Data:       nil,
	}

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Login", c.LoginDto).Return(nil, want)
	c.On("Bind", &dto.Login{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Login(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLoginGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	srv.On("Login", c.LoginDto).Return(nil, u.ServiceDownErr)
	c.On("Bind", &dto.Login{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Login(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLogoutSuccess() {
	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	srv.On("Logout", u.User.Id).Return(true, nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Logout(c)

	assert.True(u.T(), c.V.(bool))
}

func (u *AuthHandlerTest) TestLogoutBadRequest() {
	want := u.BadRequestErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	srv.On("Logout", u.User.Id).Return(nil, u.BadRequestErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Logout(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestLogoutGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	srv.On("Logout", u.User.Id).Return(nil, u.ServiceDownErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Logout(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestChangePasswordSuccess() {
	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	c.On("Bind", &dto.ChangePassword{}).Return(nil)
	srv.On("ChangePassword", u.ChangePassword).Return(true, nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.ChangePassword(c)

	assert.True(u.T(), c.V.(bool))
}

func (u *AuthHandlerTest) TestChangePasswordInvalidUserID() {
	want := u.BadRequestErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	c.On("Bind", &dto.ChangePassword{}).Return(nil)
	srv.On("Logout", u.User.Id).Return(nil, u.BadRequestErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Logout(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestChangePasswordGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	c.On("Bind", &dto.ChangePassword{}).Return(nil)
	srv.On("ChangePassword", u.ChangePassword).Return(nil, u.ServiceDownErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.ChangePassword(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestValidateSuccess() {
	want := u.User

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	userSrv.On("FindOne", int32(u.User.Id)).Return(u.User, nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestValidateFail() {
	want := u.UnauthorizedErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	userSrv.On("FindOne", int32(u.User.Id)).Return(nil, u.UnauthorizedErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestValidateGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
	}

	c.On("UserID").Return(int(u.User.Id))
	userSrv.On("FindOne", int32(u.User.Id)).Return(nil, u.ServiceDownErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.Validate(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRefreshTokenSuccess() {
	want := u.Credential

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
		RefreshToken:   u.RefreshToken,
	}

	c.On("Bind", &dto.RedeemNewToken{}).Return(nil)
	srv.On("RefreshToken", u.RefreshToken.RefreshToken).Return(u.Credential, nil)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.RefreshToken(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRefreshTokenUnauthorized() {
	want := u.UnauthorizedErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
		RefreshToken:   u.RefreshToken,
	}

	c.On("Bind", &dto.RedeemNewToken{}).Return(nil)
	srv.On("RefreshToken", u.RefreshToken.RefreshToken).Return(nil, u.UnauthorizedErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.RefreshToken(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *AuthHandlerTest) TestRefreshTokenGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	userSrv := new(user.ServiceMock)
	c := &ContextMock{
		User:           u.User,
		RegisterDto:    u.RegisterDto,
		LoginDto:       u.LoginDto,
		ChangePassword: u.ChangePassword,
		RefreshToken:   u.RefreshToken,
	}

	c.On("Bind", &dto.RedeemNewToken{}).Return(nil)
	srv.On("RefreshToken", u.RefreshToken.RefreshToken).Return(nil, u.ServiceDownErr)

	v, _ := validator.NewValidator()

	h := handler.NewAuthHandler(srv, userSrv, v)

	h.RefreshToken(c)

	assert.Equal(u.T(), want, c.V)
}
