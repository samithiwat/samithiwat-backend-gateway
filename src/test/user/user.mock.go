package user

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ServiceMock struct {
	mock.Mock
}

func (m *ServiceMock) FindAll(params *dto.PaginationQueryParams) (res *proto.UserPagination, err *dto.ResponseErr) {
	args := m.Called(params)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserPagination)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *ServiceMock) FindOne(id int32) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *ServiceMock) Create(user *dto.UserDto) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(user)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *ServiceMock) Update(id int32, user *dto.UserDto) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id, user)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *ServiceMock) Delete(id int32) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (res *proto.UserPaginationResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserPaginationResponse)
	}

	return args.Get(0).(*proto.UserPaginationResponse), args.Error(1)
}

func (m *ClientMock) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *ClientMock) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (m *ClientMock) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *ClientMock) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *ClientMock) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

type ContextMock struct {
	mock.Mock
	V       interface{}
	User    *proto.User
	Users   []*proto.User
	UserDto *dto.UserDto
	Query   *dto.PaginationQueryParams
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.UserDto) = *c.UserDto

	return args.Error(0)
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) ID() (int32, error) {
	args := c.Called()

	return int32(args.Int(0)), args.Error(1)
}

func (c *ContextMock) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	args := c.Called(query)

	*query = *c.Query

	return args.Error(0)
}
