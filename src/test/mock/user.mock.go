package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var User1 proto.User
var User2 proto.User
var User3 proto.User
var User4 proto.User
var Users []*proto.User

type UserMockClient struct {
}

func (u UserMockClient) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	return nil, nil
}

func (u UserMockClient) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &User1,
	}, nil
}

func (u UserMockClient) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (u UserMockClient) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       &User1,
	}, nil
}

func (u UserMockClient) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &User1,
	}, nil
}

func (u UserMockClient) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &User1,
	}, nil
}

type UserMockErrClient struct {
}

func (u UserMockErrClient) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	return nil, nil
}

func (u UserMockErrClient) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

func (u UserMockErrClient) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (u UserMockErrClient) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated username or email"},
		Data:       nil,
	}, nil
}

func (u UserMockErrClient) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

func (u UserMockErrClient) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

type UserMockErrGrpcClient struct {
}

func (u UserMockErrGrpcClient) FindAll(ctx context.Context, in *proto.FindAllUserRequest, opts ...grpc.CallOption) (*proto.UserPaginationResponse, error) {
	return nil, nil
}

func (u UserMockErrGrpcClient) FindOne(ctx context.Context, in *proto.FindOneUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, errors.New("Service is down")
}

func (u UserMockErrGrpcClient) FindMulti(ctx context.Context, in *proto.FindMultiUserRequest, opts ...grpc.CallOption) (*proto.UserListResponse, error) {
	return nil, nil
}

func (u UserMockErrGrpcClient) Create(ctx context.Context, in *proto.CreateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, errors.New("Service is down")
}

func (u UserMockErrGrpcClient) Update(ctx context.Context, in *proto.UpdateUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, errors.New("Service is down")
}

func (u UserMockErrGrpcClient) Delete(ctx context.Context, in *proto.DeleteUserRequest, opts ...grpc.CallOption) (*proto.UserResponse, error) {
	return nil, errors.New("Service is down")
}

type UserMockContext struct {
	V map[string]interface{}
}

func (UserMockContext) Bind(v interface{}) error {
	*v.(*proto.User) = User1
	return nil
}
func (c *UserMockContext) JSON(_ int, v interface{}) {
	c.V = v.(map[string]interface{})
}
func (UserMockContext) UserId() uint {
	return 1
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
