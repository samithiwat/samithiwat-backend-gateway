package test

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := &proto.UserPagination{
		Items: mock.Users,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindAllGrpcErrUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewUserService(&mock.UserMockErrGrpcClient{})

	c := &mock.UserMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindOneUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := &mock.User1

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneErrorNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found user"},
	}

	srv := service.NewUserService(&mock.UserMockErrClient{})

	c := &mock.UserMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneGrpcErrUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewUserService(&mock.UserMockErrGrpcClient{})

	c := &mock.UserMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestCreateUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := &mock.User1

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateErrorDuplicatedUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusUnprocessableEntity),
		"Message":    []string{"Duplicated username or email"},
	}

	srv := service.NewUserService(&mock.UserMockErrClient{})

	c := &mock.UserMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateGrpcErrUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewUserService(&mock.UserMockErrGrpcClient{})

	c := &mock.UserMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestUpdateUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := &mock.User1

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateErrorNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found user"},
	}

	srv := service.NewUserService(&mock.UserMockErrClient{})

	c := &mock.UserMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateGrpcErrUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewUserService(&mock.UserMockErrGrpcClient{})

	c := &mock.UserMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestDeleteUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := &mock.User1

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteErrorNotFoundUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found user"},
	}

	srv := service.NewUserService(&mock.UserMockErrClient{})

	c := &mock.UserMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteGrpcErrUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewUserService(&mock.UserMockErrGrpcClient{})

	c := &mock.UserMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}
