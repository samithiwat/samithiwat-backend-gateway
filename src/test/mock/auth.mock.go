package mock

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
)

type AuthContextMock struct {
	mock.Mock
	User           *proto.User
	RegisterDto    *dto.Register
	LoginDto       *dto.Login
	ChangePassword *dto.ChangePassword
	V              interface{}
}

func (c *AuthContextMock) Bind(v interface{}) error {
	args := c.Called(v)
	switch v.(type) {
	case *dto.Register:
		*v.(*dto.Register) = *c.RegisterDto
	case *dto.Login:
		*v.(*dto.Login) = *c.LoginDto
	case *dto.ChangePassword:
		*v.(*dto.ChangePassword) = *c.ChangePassword
	}

	return args.Error(0)
}

func (c *AuthContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

type AuthServiceMock struct {
	mock.Mock
}

func (s *AuthServiceMock) Register(register *dto.Register) (*proto.User, *dto.ResponseErr) {
	args := s.Called(register)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) Login(login *dto.Login) (*proto.Credential, *dto.ResponseErr) {
	args := s.Called(login)

	return args.Get(0).(*proto.Credential), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) Logout(token string) (bool, *dto.ResponseErr) {
	args := s.Called(token)

	return args.Bool(0), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) ChangePassword(chPwd *dto.ChangePassword) (bool, *dto.ResponseErr) {
	args := s.Called(chPwd)

	return args.Bool(0), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) Validate(token string) (uint32, *dto.ResponseErr) {
	args := s.Called(token)

	return uint32(args.Int(0)), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) RefreshToken(token string) (*proto.Credential, *dto.ResponseErr) {
	args := s.Called(token)

	return args.Get(0).(*proto.Credential), args.Get(1).(*dto.ResponseErr)
}
