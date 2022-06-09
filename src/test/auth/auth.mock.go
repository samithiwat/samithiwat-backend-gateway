package auth

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ContextMock struct {
	mock.Mock
	User           *proto.User
	RegisterDto    *dto.Register
	LoginDto       *dto.Login
	ChangePassword *dto.ChangePassword
	RefreshToken   *dto.RedeemNewToken
	V              interface{}
	Header         map[string]string
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)
	switch v.(type) {
	case *dto.Register:
		*v.(*dto.Register) = *c.RegisterDto
	case *dto.Login:
		*v.(*dto.Login) = *c.LoginDto
	case *dto.ChangePassword:
		*v.(*dto.ChangePassword) = *c.ChangePassword
	case *dto.RedeemNewToken:
		*v.(*dto.RedeemNewToken) = *c.RefreshToken
	}

	return args.Error(0)
}

func (c *ContextMock) UserID() int32 {
	args := c.Called()

	return int32(args.Int(0))
}

func (c *ContextMock) ID() (int32, error) {
	args := c.Called()

	return int32(args.Int(0)), args.Error(1)
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) Token() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) StoreValue(key string, val string) {
	_ = c.Called(key, val)

	c.Header = map[string]string{key: val}
}

func (c *ContextMock) Method() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Path() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Next() {
	_ = c.Called()

	return
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Register(register *dto.Register) (res *proto.User, err *dto.ResponseErr) {
	args := s.Called(register)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Login(login *dto.Login) (res *proto.Credential, err *dto.ResponseErr) {
	args := s.Called(login)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Credential)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Logout(userId uint32) (res bool, err *dto.ResponseErr) {
	args := s.Called(userId)

	if args.Get(0) != nil {
		res = args.Bool(0)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) ChangePassword(chPwd *dto.ChangePassword) (res bool, err *dto.ResponseErr) {
	args := s.Called(chPwd)

	if args.Get(0) != nil {
		res = args.Bool(0)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Validate(token string) (userId uint32, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return uint32(args.Int(0)), err
}

func (s *ServiceMock) RefreshToken(token string) (res *proto.Credential, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Credential)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Register(ctx context.Context, in *proto.RegisterRequest, opts ...grpc.CallOption) (res *proto.RegisterResponse, err error) {
	args := c.Called(in.Register)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.RegisterResponse)
	}

	return res, args.Error(1)
}
func (c *ClientMock) Login(ctx context.Context, in *proto.LoginRequest, opts ...grpc.CallOption) (res *proto.LoginResponse, err error) {
	args := c.Called(in.Login)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.LoginResponse)
	}

	return res, args.Error(1)
}
func (c *ClientMock) Logout(ctx context.Context, in *proto.LogoutRequest, opts ...grpc.CallOption) (res *proto.LogoutResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.LogoutResponse)
	}

	return res, args.Error(1)
}
func (c *ClientMock) ChangePassword(ctx context.Context, in *proto.ChangePasswordRequest, opts ...grpc.CallOption) (res *proto.ChangePasswordResponse, err error) {
	args := c.Called(in.ChangePassword)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.ChangePasswordResponse)
	}

	return res, args.Error(1)
}
func (c *ClientMock) Validate(ctx context.Context, in *proto.ValidateRequest, opts ...grpc.CallOption) (res *proto.ValidateResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.ValidateResponse)
	}

	return res, args.Error(1)
}
func (c *ClientMock) RefreshToken(ctx context.Context, in *proto.RefreshTokenRequest, opts ...grpc.CallOption) (res *proto.RefreshTokenResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.RefreshTokenResponse)
	}

	return res, args.Error(1)
}
