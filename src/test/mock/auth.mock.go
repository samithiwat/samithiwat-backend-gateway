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
	Header         map[string]string
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

func (c *AuthContextMock) ID() (int32, error) {
	args := c.Called()

	return int32(args.Int(0)), args.Error(1)
}

func (c *AuthContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *AuthContextMock) Token() string {
	args := c.Called()

	return args.String(0)
}

func (c *AuthContextMock) SetHeader(key string, val string) {
	_ = c.Called(key, val)

	c.Header = map[string]string{key: val}
}

func (c *AuthContextMock) Path() string {
	args := c.Called()

	return args.String(0)
}

func (c *AuthContextMock) Next() {
	_ = c.Called()

	return
}

type AuthServiceMock struct {
	mock.Mock
}

func (s *AuthServiceMock) Register(register *dto.Register) (res *proto.User, err *dto.ResponseErr) {
	args := s.Called(register)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *AuthServiceMock) Login(login *dto.Login) (res *proto.Credential, err *dto.ResponseErr) {
	args := s.Called(login)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Credential)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *AuthServiceMock) Logout(token string) (res bool, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		res = args.Bool(0)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *AuthServiceMock) ChangePassword(chPwd *dto.ChangePassword) (bool, *dto.ResponseErr) {
	args := s.Called(chPwd)

	return args.Bool(0), args.Get(1).(*dto.ResponseErr)
}

func (s *AuthServiceMock) Validate(token string) (userId uint32, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return uint32(args.Int(0)), err
}

func (s *AuthServiceMock) RefreshToken(token string) (*proto.Credential, *dto.ResponseErr) {
	args := s.Called(token)

	return args.Get(0).(*proto.Credential), args.Get(1).(*dto.ResponseErr)
}
