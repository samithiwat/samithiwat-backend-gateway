package blog_user

import (
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type BlogUserHandlerTest struct {
	suite.Suite
	BlogUser        *proto.BlogUser
	BlogUsers       []*proto.BlogUser
	BlogUserDto     *dto.BlogUserDto
	Query           *dto.PaginationQueryParams
	NotFoundErr     *dto.ResponseErr
	NotFoundPostErr *dto.ResponseErr
	ServiceDownErr  *dto.ResponseErr
	InvalidIDErr    *dto.ResponseErr
}

func TestBlogUserHandler(t *testing.T) {
	suite.Run(t, new(BlogUserHandlerTest))
}

func (u *BlogUserHandlerTest) SetupTest() {
	u.BlogUser = &proto.BlogUser{
		Id:          1,
		Firstname:   faker.FirstName(),
		Lastname:    faker.LastName(),
		DisplayName: faker.Username(),
		ImageUrl:    faker.URL(),
		Description: faker.Sentence(),
	}

	_ = faker.FakeData(&u.BlogUserDto)
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

	u.NotFoundPostErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found blog post",
		Data:       nil,
	}

	u.InvalidIDErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
	}
}

func (u *BlogUserHandlerTest) TestFindOneBlogUser() {
	want := u.BlogUser

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("FindOne", int32(1)).Return(u.BlogUser, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindOneInvalidRequestParamIDBlogUser() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("FindOne", int32(1)).Return(&proto.BlogUser{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindOneErrorNotFoundBlogUser() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("FindOne", int32(1)).Return(&proto.BlogUser{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindOneGrpcErrBlogUser() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("FindOne", int32(1)).Return(&proto.BlogUser{}, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestCreateBlogUser() {
	want := u.BlogUser

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Create", u.BlogUserDto).Return(u.BlogUser, nil)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestCreateErrorDuplicatedBlogUser() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated blog user name",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Create", u.BlogUserDto).Return(&proto.BlogUser{}, want)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestCreateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse blog user dto",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Create", u.BlogUserDto).Return(&proto.BlogUser{}, nil)
	c.On("Bind", &dto.BlogUserDto{}).Return(errors.New("Cannot parse body request"))

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestCreateGrpcErrBlogUser() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Create", u.BlogUserDto).Return(&proto.BlogUser{}, u.ServiceDownErr)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestUpdateBlogUser() {
	want := u.BlogUser

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Update", int32(1), u.BlogUserDto).Return(u.BlogUser, nil)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestUpdateInvalidRequestParamIDBlogUser() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Update", int32(1), u.BlogUserDto).Return(&proto.BlogUser{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestUpdateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse blog user dto",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Update", int32(1), u.BlogUserDto).Return(&proto.BlogUser{}, nil)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.BlogUserDto{}).Return(errors.New("Cannot parse blog user dto"))

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestUpdateErrorNotFoundBlogUser() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Update", int32(1), u.BlogUserDto).Return(&proto.BlogUser{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)
	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestUpdateGrpcErrBlogUser() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Update", int32(1), u.BlogUserDto).Return(&proto.BlogUser{}, u.ServiceDownErr)
	c.On("Bind", &dto.BlogUserDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteBlogUser() {
	want := u.BlogUser

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Delete", int32(1)).Return(u.BlogUser, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteInvalidRequestParamIDBlogUser() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Delete", int32(1)).Return(&proto.BlogUser{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteErrorNotFoundBlogUser() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Delete", int32(1)).Return(&proto.BlogUser{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteGrpcErrBlogUser() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Delete", int32(1)).Return(&proto.BlogUser{}, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestAddBookmarkSuccess() {
	var want []*uint32
	for i := 0; i < 5; i++ {
		num := uint32(i)
		want = append(want, &num)
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("AddBookmark", int32(1), int32(1)).Return(want, nil)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.AddBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestAddBookmarkNotFound() {
	want := u.NotFoundPostErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []*uint32

	srv.On("AddBookmark", int32(1), int32(1)).Return(out, u.NotFoundPostErr)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.AddBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestAddBookmarkInvalidRequestParamID() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("AddBookmark", int32(1), int32(1)).Return([]uint32{1, 2, 3, 4}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.AddBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestAddBookmarkGrpcErrBlogUser() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []uint32

	srv.On("AddBookmark", int32(1), int32(1)).Return(out, u.ServiceDownErr)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.AddBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindAllBookmarkSuccess() {
	var want []*uint32
	for i := 0; i < 5; i++ {
		num := uint32(i)
		want = append(want, &num)
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("FindAllBookmark", int32(1)).Return(want, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindAllBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindAllBookmarkNotFound() {
	want := u.NotFoundPostErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []*uint32

	srv.On("FindAllBookmark", int32(1)).Return(out, u.NotFoundPostErr)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindAllBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestFindAllBookmarkGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []uint32

	srv.On("FindAllBookmark", int32(1)).Return(out, u.ServiceDownErr)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.FindAllBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteBookmarkSuccess() {
	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("DeleteBookmark", int32(1), int32(1)).Return(true, nil)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.DeleteBookmark(c)

	assert.True(u.T(), c.V.(bool))
}

func (u *BlogUserHandlerTest) TestDeleteBookmarkNotFound() {
	want := u.NotFoundPostErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("DeleteBookmark", int32(1), int32(1)).Return(false, u.NotFoundPostErr)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.DeleteBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteBookmarkInvalidRequestParamID() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("DeleteBookmark", int32(1), int32(1)).Return([]uint32{1, 2, 3, 4}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.DeleteBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestDeleteBookmarkGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []uint32

	srv.On("DeleteBookmark", int32(1), int32(1)).Return(out, u.ServiceDownErr)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.DeleteBookmark(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestReadSuccess() {
	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Read", int32(1), int32(1)).Return(true, nil)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Read(c)

	assert.True(u.T(), c.V.(bool))
}

func (u *BlogUserHandlerTest) TestReadNotFound() {
	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Read", int32(1), int32(1)).Return(false, nil)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Read(c)

	assert.False(u.T(), c.V.(bool))
}

func (u *BlogUserHandlerTest) TestReadInvalidRequestParamID() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	srv.On("Read", int32(1), int32(1)).Return([]uint32{1, 2, 3, 4}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Read(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *BlogUserHandlerTest) TestReadGrpcErr() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		BlogUser:    u.BlogUser,
		BlogUserDto: u.BlogUserDto,
	}

	var out []uint32

	srv.On("Read", int32(1), int32(1)).Return(out, u.ServiceDownErr)
	c.On("ID").Return(1, nil)
	c.On("UserID").Return(1)

	v, _ := validator.NewValidator()

	h := handler.NewBlogUserHandler(srv, v)

	h.Read(c)

	assert.Equal(u.T(), want, c.V)
}
