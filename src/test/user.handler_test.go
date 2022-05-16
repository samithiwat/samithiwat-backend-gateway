package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type UserHandlerTest struct {
	suite.Suite
	User           *proto.User
	Users          []*proto.User
	UserDto        *dto.UserDto
	Query          *dto.PaginationQueryParams
	InvalidIDErr   *dto.ResponseErr
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestUserHandler(t *testing.T) {
	suite.Run(t, new(UserHandlerTest))
}

func (u *UserHandlerTest) SetupTest() {
	u.User = &proto.User{
		Id:        1,
		Firstname: faker.FirstName(),
		Lastname:  faker.LastName(),
		ImageUrl:  faker.URL(),
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

	u.UserDto = &dto.UserDto{
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
	}

	u.Users = append(u.Users, u.User, User2, User3, User4)

	_ = faker.FakeData(&u.Query)

	u.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	u.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found user",
		Data:       nil,
	}

	u.InvalidIDErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
	}
}

func (u *UserHandlerTest) TestFindAllUser() {
	want := &proto.UserPagination{
		Items: u.Users,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(want, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)
	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindAllInvalidQueryParamUser() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Cannot parse query param",
	}

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(nil, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(errors.New("Cannot parse query param"))

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindAllGrpcErrUser() {
	want := u.ServiceDownErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(nil, u.ServiceDownErr)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindOneUser() {
	want := u.User

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(u.User, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindOneInvalidRequestParamIDUser() {
	want := u.InvalidIDErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(nil, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindOneErrorNotFoundUser() {
	want := u.NotFoundErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(nil, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestFindOneGrpcErrUser() {
	want := u.ServiceDownErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(nil, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestCreateUser() {
	want := u.User

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Create", u.UserDto).Return(u.User, nil)
	c.On("Bind", &dto.UserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestCreateErrorDuplicatedUser() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated username or email",
	}

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Create", u.UserDto).Return(nil, want)
	c.On("Bind", &dto.UserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestCreateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse user dto",
	}

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Create", u.UserDto).Return(nil, nil)
	c.On("Bind", &dto.UserDto{}).Return(errors.New("Cannot parse body request"))

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestCreateGrpcErrUser() {
	want := u.ServiceDownErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Create", u.UserDto).Return(nil, u.ServiceDownErr)
	c.On("Bind", &dto.UserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestUpdateUser() {
	want := u.User

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.UserDto).Return(u.User, nil)
	c.On("Bind", &dto.UserDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestUpdateInvalidRequestParamIDUser() {
	want := u.InvalidIDErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.UserDto).Return(nil, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("Bind", &dto.UserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestUpdateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse user dto",
	}

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.UserDto).Return(nil, nil)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.UserDto{}).Return(errors.New("Cannot parse user dto"))

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestUpdateErrorNotFoundUser() {
	want := u.NotFoundErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.UserDto).Return(nil, u.NotFoundErr)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.UserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)
	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestUpdateGrpcErrUser() {
	want := u.ServiceDownErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.UserDto).Return(nil, u.ServiceDownErr)
	c.On("Bind", &dto.UserDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestDeleteUser() {
	want := u.User

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(u.User, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestDeleteInvalidRequestParamIDUser() {
	want := u.InvalidIDErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(nil, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestDeleteErrorNotFoundUser() {
	want := u.NotFoundErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(nil, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *UserHandlerTest) TestDeleteGrpcErrUser() {
	want := u.ServiceDownErr

	srv := new(mock.UserServiceMock)
	c := &mock.UserContextMock{
		User:    u.User,
		Users:   u.Users,
		UserDto: u.UserDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(nil, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}
