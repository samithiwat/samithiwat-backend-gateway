package test

import (
	"github.com/bxcodec/faker/v3"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/handler"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/test/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type OrganizationHandlerTest struct {
	suite.Suite
	Organization   *proto.Organization
	Organizations  []*proto.Organization
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
	InvalidIDErr   *dto.ResponseErr
}

func TestOrganizationHandler(t *testing.T) {
	suite.Run(t, new(OrganizationHandlerTest))
}

func (u *OrganizationHandlerTest) SetupTest() {
	u.Organization = &proto.Organization{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization2 := &proto.Organization{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization3 := &proto.Organization{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization4 := &proto.Organization{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	u.Organizations = append(u.Organizations, u.Organization, Organization2, Organization3, Organization4)

	u.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	u.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found organization",
		Data:       nil,
	}

	u.InvalidIDErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
	}
}

func (u *OrganizationHandlerTest) TestFindAllOrganization() {
	want := &proto.OrganizationPagination{
		Items: u.Organizations,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindAll").Return(want, &dto.ResponseErr{})
	c.On("PaginationQueryParam").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindAllInvalidQueryParamOrganization() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Cannot parse query param",
	}

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindAll").Return(nil, nil)
	c.On("PaginationQueryParam").Return(errors.New("Cannot parse query param"))

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindAllGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindAll").Return(&proto.OrganizationPagination{}, u.ServiceDownErr)
	c.On("PaginationQueryParam").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneOrganization() {
	want := u.Organization

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindOne", int32(1)).Return(u.Organization, &dto.ResponseErr{})
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, &dto.ResponseErr{})
	c.On("ID").Return(errors.New("Invalid ID"))

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateOrganization() {
	want := u.Organization

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Create").Return(u.Organization, &dto.ResponseErr{})
	c.On("Bind").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateErrorDuplicatedOrganization() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated organization name",
	}

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Create").Return(&proto.Organization{}, want)
	c.On("Bind").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse organization dto",
	}

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Create").Return(&proto.Organization{}, &dto.ResponseErr{})
	c.On("Bind").Return(errors.New("Cannot parse body request"))

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Create").Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("Bind").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateOrganization() {
	want := u.Organization

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Update", int32(1)).Return(u.Organization, &dto.ResponseErr{})
	c.On("Bind").Return(nil)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Update", int32(1)).Return(&proto.Organization{}, &dto.ResponseErr{})
	c.On("ID").Return(errors.New("Invalid ID"))
	c.On("Bind").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse organization dto",
	}

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Update", int32(1)).Return(&proto.Organization{}, &dto.ResponseErr{})
	c.On("ID").Return(nil)
	c.On("Bind").Return(errors.New("Cannot parse organization dto"))

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Update", int32(1)).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(nil)
	c.On("Bind").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)
	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Update", int32(1)).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("Bind").Return(nil)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteOrganization() {
	want := u.Organization

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Delete", int32(1)).Return(u.Organization, &dto.ResponseErr{})
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, &dto.ResponseErr{})
	c.On("ID").Return(errors.New("Invalid ID"))

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(mock.OrganizationServiceMock)
	c := new(mock.OrganizationContextMock)

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("ID").Return(nil)

	v := validator.New()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}
