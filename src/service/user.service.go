package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
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

func (s *UserService) FindAll(query *dto.PaginationQueryParams) (result *proto.UserPagination, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.FindAllUserRequest{
		Page:  query.Page,
		Limit: query.Limit,
	}

	res, errRes := s.client.FindAll(ctx, req)
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *UserService) FindOne(id int32) (result *proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.FindOne(ctx, &proto.FindOneUserRequest{Id: id})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *UserService) Create(userDto *dto.UserDto) (result *proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := s.DtoToRaw(userDto)

	res, errRes := s.client.Create(ctx, &proto.CreateUserRequest{User: user})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusCreated {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *UserService) Update(id int32, userDto *dto.UserDto) (result *proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := s.DtoToRaw(userDto)
	user.Id = uint32(id)

	res, errRes := s.client.Update(ctx, &proto.UpdateUserRequest{User: user})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *UserService) Delete(id int32) (result *proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.Delete(ctx, &proto.DeleteUserRequest{Id: id})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (UserService) DtoToRaw(userDto *dto.UserDto) *proto.User {
	return &proto.User{
		Firstname: userDto.Firstname,
		Lastname:  userDto.Lastname,
		ImageUrl:  userDto.ImageUrl,
	}
}
