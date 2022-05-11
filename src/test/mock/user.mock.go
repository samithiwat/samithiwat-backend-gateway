package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type UserServiceMock struct {
	mock.Mock
}

func (m *UserServiceMock) FindAll(params dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr) {
	args := m.Called()

	return args.Get(0).(*proto.UserPagination), args.Get(1).(*dto.ResponseErr)
}

func (m *UserServiceMock) FindOne(id int32) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m *UserServiceMock) Create(user dto.UserDto) (*proto.User, *dto.ResponseErr) {
	args := m.Called()

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m *UserServiceMock) Update(id int32, user dto.UserDto) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m *UserServiceMock) Delete(id int32) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

type UserClientMock struct {
	mock.Mock
}

func (m *UserClientMock) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserPaginationResponse), args.Error(1)
}

func (m *UserClientMock) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m *UserClientMock) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (m *UserClientMock) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m *UserClientMock) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m *UserClientMock) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

type UserContextMock struct {
	V interface{}
	mock.Mock
}

func (c *UserContextMock) Bind(v interface{}) error {
	*v.(*dto.UserDto) = dto.UserDto{
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	args := c.Called()

	return args.Error(0)
}

func (c *UserContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *UserContextMock) ID(id *int32) error {
	*id = 1

	args := c.Called()

	return args.Error(0)
}

func (c *UserContextMock) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	args := c.Called()

	return args.Error(0)
}
