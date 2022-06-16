package organization

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

type OrganizationHandlerTest struct {
	suite.Suite
	Organization    *proto.Organization
	Organizations   []*proto.Organization
	OrganizationDto *dto.OrganizationDto
	Query           *dto.PaginationQueryParams
	NotFoundErr     *dto.ResponseErr
	ServiceDownErr  *dto.ResponseErr
	InvalidIDErr    *dto.ResponseErr
}

func TestOrganizationHandler(t *testing.T) {
	suite.Run(t, new(OrganizationHandlerTest))
}

func (u *OrganizationHandlerTest) SetupTest() {
	u.Organization = &proto.Organization{
		Id:          1,
		Name:        faker.Word(),
		Email:       faker.Email(),
		Description: faker.Sentence(),
	}

	Organization2 := &proto.Organization{
		Id:          2,
		Name:        faker.Word(),
		Email:       faker.Email(),
		Description: faker.Sentence(),
	}

	Organization3 := &proto.Organization{
		Id:          3,
		Name:        faker.Word(),
		Email:       faker.Email(),
		Description: faker.Sentence(),
	}

	Organization4 := &proto.Organization{
		Id:          4,
		Name:        faker.Word(),
		Email:       faker.Email(),
		Description: faker.Sentence(),
	}

	u.Organizations = append(u.Organizations, u.Organization, Organization2, Organization3, Organization4)

	_ = faker.FakeData(&u.OrganizationDto)
	_ = faker.FakeData(&u.Query)

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

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindAll", u.Query).Return(want, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindAllInvalidQueryParamOrganization() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Cannot parse query param",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindAll", u.Query).Return(nil, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(errors.New("Cannot parse query param"))

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindAllGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindAll", u.Query).Return(nil, u.ServiceDownErr)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneOrganization() {
	want := u.Organization

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindOne", int32(1)).Return(u.Organization, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestFindOneGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateOrganization() {
	want := u.Organization

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Create", u.OrganizationDto).Return(u.Organization, nil)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateErrorDuplicatedOrganization() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated organization name",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Create", u.OrganizationDto).Return(&proto.Organization{}, want)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse organization dto",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Create", u.OrganizationDto).Return(&proto.Organization{}, nil)
	c.On("Bind", &dto.OrganizationDto{}).Return(errors.New("Cannot parse body request"))

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestCreateGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Create", u.OrganizationDto).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateOrganization() {
	want := u.Organization

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Update", int32(1), u.OrganizationDto).Return(u.Organization, nil)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Update", int32(1), u.OrganizationDto).Return(&proto.Organization{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse organization dto",
	}

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Update", int32(1), u.OrganizationDto).Return(&proto.Organization{}, nil)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.OrganizationDto{}).Return(errors.New("Cannot parse organization dto"))

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Update", int32(1), u.OrganizationDto).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)
	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestUpdateGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Update", int32(1), u.OrganizationDto).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("Bind", &dto.OrganizationDto{}).Return(nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteOrganization() {
	want := u.Organization

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Delete", int32(1)).Return(u.Organization, nil)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteInvalidRequestParamIDOrganization() {
	want := u.InvalidIDErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, nil)
	c.On("ID").Return(-1, errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteErrorNotFoundOrganization() {
	want := u.NotFoundErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, u.NotFoundErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *OrganizationHandlerTest) TestDeleteGrpcErrOrganization() {
	want := u.ServiceDownErr

	srv := new(ServiceMock)
	c := &ContextMock{
		Organization:    u.Organization,
		Organizations:   u.Organizations,
		OrganizationDto: u.OrganizationDto,
		Query:           u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Organization{}, u.ServiceDownErr)
	c.On("ID").Return(1, nil)

	v, _ := validator.NewValidator()

	h := handler.NewOrganizationHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}
