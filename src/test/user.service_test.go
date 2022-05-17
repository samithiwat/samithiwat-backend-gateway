package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type UserServiceTest struct {
	suite.Suite
	User           *proto.User
	UserReq        *proto.User
	Users          []*proto.User
	UserDto        *dto.UserDto
	Query          *dto.PaginationQueryParams
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (s *UserServiceTest) SetupTest() {
	s.User = &proto.User{
		Id:        1,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	s.UserReq = &proto.User{
		Firstname:   s.User.Firstname,
		Lastname:    s.User.Lastname,
		DisplayName: s.User.DisplayName,
		ImageUrl:    s.User.ImageUrl,
	}

	User2 := &proto.User{
		Id:        2,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User3 := &proto.User{
		Id:        3,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	User4 := &proto.User{
		Id:        4,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
	}

	s.Users = append(s.Users, s.User, User2, User3, User4)

	s.UserDto = &dto.UserDto{
		Firstname:   s.User.Firstname,
		Lastname:    s.User.Lastname,
		DisplayName: s.User.DisplayName,
		ImageUrl:    s.User.ImageUrl,
	}

	_ = faker.FakeData(&s.Query)

	s.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	s.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found user",
		Data:       nil,
	}
}

func (s *UserServiceTest) TestFindAllUserService() {
	want := &proto.UserPagination{
		Items: s.Users,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	client := new(mock.UserClientMock)

	client.On("FindAll", &proto.FindAllUserRequest{
		Limit: s.Query.Limit,
		Page:  s.Query.Page,
	}).Return(&proto.UserPaginationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       want,
	}, nil)

	srv := service.NewUserService(client)

	users, err := srv.FindAll(s.Query)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, users)
}

func (s *UserServiceTest) TestFindAllGrpcErrUserService() {
	want := s.ServiceDownErr

	client := new(mock.UserClientMock)

	client.On("FindAll", &proto.FindAllUserRequest{
		Limit: s.Query.Limit,
		Page:  s.Query.Page,
	}).Return(&proto.UserPaginationResponse{}, errors.New("Service is down"))

	srv := service.NewUserService(client)

	_, err := srv.FindAll(s.Query)

	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestFindOneUserService() {
	want := s.User

	client := new(mock.UserClientMock)

	var id int32
	_ = faker.FakeData(&id)

	client.On("FindOne", &proto.FindOneUserRequest{Id: id}).Return(&proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.User,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.FindOne(id)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, user)
}

func (s *UserServiceTest) TestFindOneNotFoundUserService() {
	want := s.NotFoundErr

	client := new(mock.UserClientMock)

	var id int32
	_ = faker.FakeData(&id)

	client.On("FindOne", &proto.FindOneUserRequest{Id: id}).Return(&proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.FindOne(id)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestFindOneGrpcErrUserService() {
	want := s.ServiceDownErr

	client := new(mock.UserClientMock)

	var id int32
	_ = faker.FakeData(&id)

	client.On("FindOne", &proto.FindOneUserRequest{Id: id}).Return(&proto.UserResponse{}, errors.New("Service is down"))

	srv := service.NewUserService(client)

	_, err := srv.FindOne(id)

	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestCreateUserService() {
	want := s.User

	client := new(mock.UserClientMock)

	client.On("Create", *s.UserReq).Return(&proto.UserResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       s.User,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Create(s.UserDto)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, user)
}

func (s *UserServiceTest) TestCreateDuplicatedUserService() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated email or username",
		Data:       nil,
	}

	client := new(mock.UserClientMock)

	client.On("Create", *s.UserReq).Return(&proto.UserResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated email or username"},
		Data:       nil,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Create(s.UserDto)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestCreateGrpcErrUserService() {
	want := s.ServiceDownErr

	client := new(mock.UserClientMock)

	client.On("Create", *s.UserReq).Return(&proto.UserResponse{}, errors.New("Service is down"))

	srv := service.NewUserService(client)

	_, err := srv.Create(s.UserDto)

	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestUpdateUserService() {
	want := s.User

	client := new(mock.UserClientMock)

	client.On("Update", *s.User).Return(&proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.User,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Update(1, s.UserDto)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, user)
}

func (s *UserServiceTest) TestUpdateNotFoundUserService() {
	want := s.NotFoundErr

	client := new(mock.UserClientMock)

	client.On("Update", *s.User).Return(&proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Update(1, s.UserDto)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestUpdateGrpcErrUserService() {
	want := s.ServiceDownErr

	client := new(mock.UserClientMock)

	client.On("Update", *s.User).Return(&proto.UserResponse{}, errors.New("Service is down"))

	srv := service.NewUserService(client)

	_, err := srv.Update(1, s.UserDto)

	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestDeleteUserService() {
	want := s.User

	var id int32
	_ = faker.FakeData(&id)

	client := new(mock.UserClientMock)

	client.On("Delete", &proto.DeleteUserRequest{Id: id}).Return(&proto.UserResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.User,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Delete(id)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, user)
}

func (s *UserServiceTest) TestDeleteNotFoundUserService() {
	want := s.NotFoundErr

	var id int32
	_ = faker.FakeData(&id)

	client := new(mock.UserClientMock)

	client.On("Delete", &proto.DeleteUserRequest{Id: id}).Return(&proto.UserResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil)

	srv := service.NewUserService(client)

	user, err := srv.Delete(id)

	assert.Nil(s.T(), user)
	assert.Equal(s.T(), want, err)
}

func (s *UserServiceTest) TestDeleteGrpcErrUserService() {
	want := s.ServiceDownErr

	var id int32
	_ = faker.FakeData(&id)

	client := new(mock.UserClientMock)

	client.On("Delete", &proto.DeleteUserRequest{Id: id}).Return(&proto.UserResponse{}, errors.New("Service is down"))

	srv := service.NewUserService(client)

	_, err := srv.Delete(id)

	assert.Equal(s.T(), want, err)
}
