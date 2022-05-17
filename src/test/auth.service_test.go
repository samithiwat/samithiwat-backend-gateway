package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type AuthServiceTest struct {
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

func TestAuthService(t *testing.T) {
	suite.Run(t, new(AuthServiceTest))
}

func (s *AuthServiceTest) SetupTest() {
	s.User = &proto.User{
		Id:          1,
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
	}

	s.RegisterDto = &dto.Register{
		Email:       faker.Email(),
		Password:    faker.Password(),
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
	}

	s.LoginDto = &dto.Login{
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	s.ChangePassword = &dto.ChangePassword{
		UserId:      1,
		OldPassword: faker.Password(),
		NewPassword: faker.Password(),
	}

	s.RefreshToken = &dto.RedeemNewToken{
		RefreshToken: faker.Word(),
	}

	s.Credential = &proto.Credential{
		AccessToken:  faker.Word(),
		RefreshToken: faker.Word(),
		ExpiresIn:    3600,
	}

	s.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	s.UnauthorizedErr = &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Something wrong :(",
		Data:       nil,
	}

	s.BadRequestErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid request body",
		Data:       nil,
	}
}

func (s *AuthServiceTest) TestRegisterSuccess() {
	want := s.User

	client := new(mock.AuthClientMock)

	client.On("Register", proto.Register{
		Email:       s.RegisterDto.Email,
		Password:    s.RegisterDto.Password,
		Firstname:   s.RegisterDto.Firstname,
		Lastname:    s.RegisterDto.Lastname,
		DisplayName: s.RegisterDto.DisplayName,
		ImageUrl:    s.RegisterDto.ImageUrl,
	}).Return(&proto.RegisterResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       s.User,
	}, nil)

	srv := service.NewAuthService(client)

	user, err := srv.Register(s.RegisterDto)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), want, user)
}

func (s *AuthServiceTest) TestRegisterEmailDuplicated() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated Email",
		Data:       nil,
	}

	client := new(mock.AuthClientMock)

	client.On("Register", proto.Register{
		Email:       s.RegisterDto.Email,
		Password:    s.RegisterDto.Password,
		Firstname:   s.RegisterDto.Firstname,
		Lastname:    s.RegisterDto.Lastname,
		DisplayName: s.RegisterDto.DisplayName,
		ImageUrl:    s.RegisterDto.ImageUrl,
	}).Return(&proto.RegisterResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated Email"},
		Data:       nil,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Register(s.RegisterDto)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestRegisterGrpcError() {
	want := s.ServiceDownErr

	client := new(mock.AuthClientMock)

	client.On("Register", proto.Register{
		Email:       s.RegisterDto.Email,
		Password:    s.RegisterDto.Password,
		Firstname:   s.RegisterDto.Firstname,
		Lastname:    s.RegisterDto.Lastname,
		DisplayName: s.RegisterDto.DisplayName,
		ImageUrl:    s.RegisterDto.ImageUrl,
	}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.Register(s.RegisterDto)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestLoginSuccess() {
	want := s.Credential

	client := new(mock.AuthClientMock)

	client.On("Login", proto.Login{
		Email:    s.LoginDto.Email,
		Password: s.LoginDto.Password,
	}).Return(&proto.LoginResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Credential,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Login(s.LoginDto)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), want, res)
}

func (s *AuthServiceTest) TestLoginUnauthorized() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
		Data:       nil,
	}

	client := new(mock.AuthClientMock)

	client.On("Login", proto.Login{
		Email:    s.LoginDto.Email,
		Password: s.LoginDto.Password,
	}).Return(&proto.LoginResponse{
		StatusCode: http.StatusUnauthorized,
		Errors:     []string{"Invalid Token"},
		Data:       nil,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Login(s.LoginDto)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestLoginGrpcError() {
	want := s.ServiceDownErr

	client := new(mock.AuthClientMock)

	client.On("Login", proto.Login{
		Email:    s.LoginDto.Email,
		Password: s.LoginDto.Password,
	}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.Login(s.LoginDto)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestLogoutSuccess() {
	client := new(mock.AuthClientMock)

	client.On("Logout", &proto.LogoutRequest{UserId: s.User.Id}).Return(&proto.LogoutResponse{
		StatusCode: http.StatusNoContent,
		Errors:     nil,
		Data:       true,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Logout(s.User.Id)

	assert.Nil(s.T(), err)
	assert.True(s.T(), res)
}

func (s *AuthServiceTest) TestLogoutUnauthorized() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
		Data:       nil,
	}

	client := new(mock.AuthClientMock)

	client.On("Logout", &proto.LogoutRequest{UserId: s.User.Id}).Return(&proto.LogoutResponse{
		StatusCode: http.StatusUnauthorized,
		Errors:     []string{"Invalid Token"},
		Data:       false,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Logout(s.User.Id)

	assert.False(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestLogoutGrpcError() {
	want := s.ServiceDownErr

	client := new(mock.AuthClientMock)

	client.On("Logout", &proto.LogoutRequest{
		UserId: s.User.Id,
	}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.Logout(s.User.Id)

	assert.False(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestChangePasswordSuccess() {
	client := new(mock.AuthClientMock)

	client.On("ChangePassword", proto.ChangePassword{
		UserId:      s.ChangePassword.UserId,
		OldPassword: s.ChangePassword.OldPassword,
		NewPassword: s.ChangePassword.NewPassword,
	}).Return(&proto.ChangePasswordResponse{
		StatusCode: http.StatusNoContent,
		Errors:     nil,
		Data:       true,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.ChangePassword(s.ChangePassword)

	assert.Nil(s.T(), err)
	assert.True(s.T(), res)
}

func (s *AuthServiceTest) TestChangePasswordUnauthorized() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
		Data:       nil,
	}

	client := new(mock.AuthClientMock)

	client.On("ChangePassword", proto.ChangePassword{
		UserId:      s.ChangePassword.UserId,
		OldPassword: s.ChangePassword.OldPassword,
		NewPassword: s.ChangePassword.NewPassword,
	}).Return(&proto.ChangePasswordResponse{
		StatusCode: http.StatusUnauthorized,
		Errors:     []string{"Invalid Token"},
		Data:       false,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.ChangePassword(s.ChangePassword)

	assert.False(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestChangePasswordGrpcError() {
	want := s.ServiceDownErr

	client := new(mock.AuthClientMock)

	client.On("ChangePassword", proto.ChangePassword{
		UserId:      s.ChangePassword.UserId,
		OldPassword: s.ChangePassword.OldPassword,
		NewPassword: s.ChangePassword.NewPassword,
	}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.ChangePassword(s.ChangePassword)

	assert.False(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestValidateSuccess() {
	want := s.User.Id
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("Validate", &proto.ValidateRequest{Token: token}).Return(&proto.ValidateResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.User.Id,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Validate(token)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), want, res)
}

func (s *AuthServiceTest) TestValidateUnauthorized() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
		Data:       nil,
	}
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("Validate", &proto.ValidateRequest{Token: token}).Return(&proto.ValidateResponse{
		StatusCode: http.StatusUnauthorized,
		Errors:     []string{"Invalid Token"},
		Data:       0,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.Validate(token)

	assert.Equal(s.T(), uint32(0), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestValidateGrpcError() {
	want := s.ServiceDownErr
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("Validate", &proto.ValidateRequest{Token: token}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.Validate(token)

	assert.Equal(s.T(), uint32(0), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestRefreshTokenSuccess() {
	want := s.Credential
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("RefreshToken", &proto.RefreshTokenRequest{RefreshToken: token}).Return(&proto.RefreshTokenResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Credential,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.RefreshToken(token)

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), want, res)
}

func (s *AuthServiceTest) TestRefreshTokenUnauthorized() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnauthorized,
		Message:    "Invalid Token",
		Data:       nil,
	}
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("RefreshToken", &proto.RefreshTokenRequest{RefreshToken: token}).Return(&proto.RefreshTokenResponse{
		StatusCode: http.StatusUnauthorized,
		Errors:     []string{"Invalid Token"},
		Data:       nil,
	}, nil)

	srv := service.NewAuthService(client)

	res, err := srv.RefreshToken(token)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}

func (s *AuthServiceTest) TestRefreshTokenGrpcError() {
	want := s.ServiceDownErr
	token := faker.Word()

	client := new(mock.AuthClientMock)

	client.On("RefreshToken", &proto.RefreshTokenRequest{RefreshToken: token}).Return(nil, errors.New("Service is down"))

	srv := service.NewAuthService(client)

	res, err := srv.RefreshToken(token)

	assert.Nil(s.T(), res)
	assert.Equal(s.T(), want, err)
}
