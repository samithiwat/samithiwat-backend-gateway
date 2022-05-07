package test

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindOneUser(t *testing.T) {
	mock.InitializeMockUser()

	assert := assert.New(t)
	want := map[string]interface{}{
		"ID":        mock.User1.Id,
		"Firstname": mock.User1.Firstname,
		"Lastname":  mock.User1.Lastname,
		"ImageUrl":  mock.User1.ImageUrl,
	}

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneErrorNotFound(t *testing.T) {
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
	want := map[string]interface{}{
		"ID":        mock.User1.Id,
		"Firstname": mock.User1.Firstname,
		"Lastname":  mock.User1.Lastname,
		"ImageUrl":  mock.User1.ImageUrl,
	}

	srv := service.NewUserService(&mock.UserMockClient{})

	c := &mock.UserMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateErrorDuplicated(t *testing.T) {
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
