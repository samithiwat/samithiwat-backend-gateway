package mock

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) FindAll(params *dto.PaginationQueryParams) (res *proto.UserPagination, err *dto.ResponseErr) {
	args := m.Called(params)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserPagination)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *UserServiceMock) FindOne(id int32) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *UserServiceMock) Create(user *dto.UserDto) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(user)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *UserServiceMock) Update(id int32, user *dto.UserDto) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id, user)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (m *UserServiceMock) Delete(id int32) (res *proto.User, err *dto.ResponseErr) {
	args := m.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type UserClientMock struct {
	mock.Mock
}

func (m *UserClientMock) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (res *proto.UserPaginationResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserPaginationResponse)
	}

	return args.Get(0).(*proto.UserPaginationResponse), args.Error(1)
}

func (m *UserClientMock) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *UserClientMock) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (m *UserClientMock) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(*in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *UserClientMock) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(*in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

func (m *UserClientMock) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (res *proto.UserResponse, err error) {
	args := m.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UserResponse)
	}

	return res, args.Error(1)
}

type UserContextMock struct {
	mock.Mock
	V       interface{}
	User    *proto.User
	Users   []*proto.User
	UserDto *dto.UserDto
	Query   *dto.PaginationQueryParams
}

func (c *UserContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.UserDto) = *c.UserDto

	return args.Error(0)
}

func (c *UserContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *UserContextMock) ID(id *int32) error {
	args := c.Called(*id)

	*id = 1

	return args.Error(0)
}

func (c *UserContextMock) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	args := c.Called(query)

	*query = *c.Query

	return args.Error(0)
}
