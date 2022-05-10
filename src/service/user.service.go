package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"time"
)

type UserService struct {
	client proto.UserServiceClient
}

func NewUserService(client proto.UserServiceClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (s *UserService) FindAll(query dto.PaginationQueryParams) (res *proto.UserPaginationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.FindAllUserRequest{
		Page:  query.Page,
		Limit: query.Limit,
	}

	res, err = s.client.FindAll(ctx, req)

	return
}

func (s *UserService) FindOne(id int32) (res *proto.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.FindOne(ctx, &proto.FindOneUserRequest{Id: id})

	return
}

func (s *UserService) Create(userDto dto.UserDto) (res *proto.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := s.DtoToRaw(&userDto)

	res, err = s.client.Create(ctx, &proto.CreateUserRequest{User: user})

	return
}

func (s *UserService) Update(id int32, userDto dto.UserDto) (res *proto.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := s.DtoToRaw(&userDto)
	user.Id = uint32(id)

	res, err = s.client.Update(ctx, &proto.UpdateUserRequest{User: user})

	return
}

func (s *UserService) Delete(id int32) (res *proto.UserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err = s.client.Delete(ctx, &proto.DeleteUserRequest{Id: id})

	return
}

func (UserService) DtoToRaw(userDto *dto.UserDto) *proto.User {
	return &proto.User{
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		ImageUrl:  userDto.ImageUrl,
	}
}
