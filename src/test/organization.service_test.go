package test

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/service"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFindAllOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := &proto.OrganizationPagination{
		Items: mock.Organizations,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindAllInvalidQueryParamOrganization(t *testing.T) {
	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid query param",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockErrContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindAllGrpcErrOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrGrpcClient{})

	c := &mock.OrganizationMockContext{}

	srv.FindAll(c)

	assert.Equal(want, c.V)
}

func TestFindOneOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := &mock.Organization1

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneInvalidRequestParamIDOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneErrorNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found organization"},
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrClient{})

	c := &mock.OrganizationMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestFindOneGrpcErrOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrGrpcClient{})

	c := &mock.OrganizationMockContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestCreateOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := &mock.Organization1

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateErrorDuplicatedOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusUnprocessableEntity),
		"Message":    []string{"Duplicated organization name"},
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrClient{})

	c := &mock.OrganizationMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestCreateGrpcErrOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrGrpcClient{})

	c := &mock.OrganizationMockContext{}

	srv.Create(c)

	assert.Equal(want, c.V)
}

func TestUpdateOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := &mock.Organization1

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateInvalidRequestParamIDOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestUpdateErrorNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found organization"},
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrClient{})

	c := &mock.OrganizationMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestUpdateGrpcErrOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrGrpcClient{})

	c := &mock.OrganizationMockContext{}

	srv.Update(c)

	assert.Equal(want, c.V)
}

func TestDeleteOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := &mock.Organization1

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteInvalidRequestParamIDOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadRequest,
		"Message":    "Invalid id",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockClient{})

	c := &mock.OrganizationMockErrContext{}

	srv.FindOne(c)

	assert.Equal(want, c.V)
}

func TestDeleteErrorNotFoundOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": int32(http.StatusNotFound),
		"Message":    []string{"Not found organization"},
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrClient{})

	c := &mock.OrganizationMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}

func TestDeleteGrpcErrOrganization(t *testing.T) {
	mock.InitializeMockOrganization()

	assert := assert.New(t)
	want := map[string]interface{}{
		"StatusCode": http.StatusBadGateway,
		"Message":    "Service is down",
	}

	srv := service.NewOrganizationService(&mock.OrganizationMockErrGrpcClient{})

	c := &mock.OrganizationMockContext{}

	srv.Delete(c)

	assert.Equal(want, c.V)
}
