package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"net/http"
)

var User1 proto.User
var User2 proto.User
var User3 proto.User
var User4 proto.User
var Users []*proto.User

type UserMockService struct {
	mock.Mock
}

func (m UserMockService) FindAll(dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr) {
	return &proto.UserPagination{
		Items: Users,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}, nil
}

func (m UserMockService) FindOne(int32) (*proto.User, *dto.ResponseErr) {
	return &User1, nil
}

func (m UserMockService) Create(dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return &User1, nil
}

func (m UserMockService) Update(int32, dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return &User1, nil
}

func (m UserMockService) Delete(int32) (*proto.User, *dto.ResponseErr) {
	return &User1, nil
}

type UserMockErrService struct {
}

func (UserMockErrService) FindAll(dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr) {
	return nil, nil
}

func (UserMockErrService) FindOne(int32) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found user",
		Data:       nil,
	}
}

func (UserMockErrService) Create(dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated email or username",
		Data:       nil,
	}
}

func (UserMockErrService) Update(int32, dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found user",
		Data:       nil,
	}
}

func (UserMockErrService) Delete(int32) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found user",
		Data:       nil,
	}
}

type UserMockErrGrpcService struct {
}

func (UserMockErrGrpcService) FindAll(dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (UserMockErrGrpcService) FindOne(int32) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (UserMockErrGrpcService) Create(dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (UserMockErrGrpcService) Update(int32, dto.UserDto) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (UserMockErrGrpcService) Delete(int32) (*proto.User, *dto.ResponseErr) {
	return nil, &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

type UserMockContext struct {
	V interface{}
}

func (UserMockContext) Bind(v interface{}) error {
	*v.(*proto.User) = User1
	return nil
}

func (c *UserMockContext) JSON(_ int, v interface{}) {
	c.V = v
}

func (UserMockContext) ID(id *int32) error {
	*id = 1
	return nil
}

func (UserMockContext) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	*query = dto.PaginationQueryParams{
		Page:  1,
		Limit: 10,
	}

	return nil
}

type UserMockErrContext struct {
	V interface{}
}

func (UserMockErrContext) Bind(v interface{}) error {
	*v.(*proto.User) = User1
	return nil
}

func (c *UserMockErrContext) JSON(_ int, v interface{}) {
	c.V = v
}

func (UserMockErrContext) ID(*int32) error {
	return errors.New("Invalid ID")
}

func (UserMockErrContext) PaginationQueryParam(*dto.PaginationQueryParams) error {
	return errors.New("Invalid Query Param")
}

func InitializeMockUser() {
	User1 = proto.User{
		Id:        1,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User2 = proto.User{
		Id:        2,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User3 = proto.User{
		Id:        3,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User4 = proto.User{
		Id:        4,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	Users = append(Users, &User1, &User2, &User3, &User4)
}

type UserServiceMock struct {
	mock.Mock
}

func (m UserServiceMock) FindAll(params dto.PaginationQueryParams) (*proto.UserPagination, *dto.ResponseErr) {
	args := m.Called(params)

	return args.Get(0).(*proto.UserPagination), nil
}

func (m UserServiceMock) FindOne(id int32) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m UserServiceMock) Create(user dto.UserDto) (*proto.User, *dto.ResponseErr) {
	args := m.Called(user)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m UserServiceMock) Update(id int32, user dto.UserDto) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id, user)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

func (m UserServiceMock) Delete(id int32) (*proto.User, *dto.ResponseErr) {
	args := m.Called(id)

	return args.Get(0).(*proto.User), args.Get(1).(*dto.ResponseErr)
}

type UserClientMock struct {
	mock.Mock
}

func (m UserClientMock) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserPaginationResponse), args.Error(1)
}

func (m UserClientMock) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m UserClientMock) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (m UserClientMock) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m UserClientMock) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}

func (m UserClientMock) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	args := m.Called()

	return args.Get(0).(*proto.UserResponse), args.Error(1)
}
