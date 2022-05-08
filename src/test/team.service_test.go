package test

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := &proto.TeamPagination{
		Items: mock.Teams,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindAllInvalidQueryParamTeam(t *testing.T) {
	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid query param",
	}

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockErrContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindAllGrpcErrTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewTeamService(&mock.TeamMockErrGrpcClient{})

	c := &mock.TeamMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindOneTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := &mock.Team1

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneInvalidRequestParamIDTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneErrorNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found team"},
	}

	srv := service.NewTeamService(&mock.TeamMockErrClient{})

	c := &mock.TeamMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneGrpcErrTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewTeamService(&mock.TeamMockErrGrpcClient{})

	c := &mock.TeamMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestCreateTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := &mock.Team1

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateErrorDuplicatedTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusUnprocessableEntity),
		"Message":    []string{"Duplicated team name"},
	}

	srv := service.NewTeamService(&mock.TeamMockErrClient{})

	c := &mock.TeamMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateGrpcErrTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewTeamService(&mock.TeamMockErrGrpcClient{})

	c := &mock.TeamMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestUpdateTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := &mock.Team1

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateInvalidRequestParamIDTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestUpdateErrorNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found team"},
	}

	srv := service.NewTeamService(&mock.TeamMockErrClient{})

	c := &mock.TeamMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateGrpcErrTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewTeamService(&mock.TeamMockErrGrpcClient{})

	c := &mock.TeamMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestDeleteTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := &mock.Team1

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteInvalidRequestParamIDTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewTeamService(&mock.TeamMockClient{})

	c := &mock.TeamMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestDeleteErrorNotFoundTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found team"},
	}

	srv := service.NewTeamService(&mock.TeamMockErrClient{})

	c := &mock.TeamMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteGrpcErrTeam(t *testing.T) {
	mock.InitializeMockTeam()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewTeamService(&mock.TeamMockErrGrpcClient{})

	c := &mock.TeamMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}
