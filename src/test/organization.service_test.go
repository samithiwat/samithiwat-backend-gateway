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

type OrganizationServiceTest struct {
	suite.Suite
	Organization   *proto.Organization
	Organizations  []*proto.Organization
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestOrganizationService(t *testing.T) {
	suite.Run(t, new(OrganizationServiceTest))
}

func (s *OrganizationServiceTest) SetupTest() {
	s.Organization = &proto.Organization{
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

	s.Organizations = append(s.Organizations, s.Organization, Organization2, Organization3, Organization4)

	s.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	s.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found organization",
		Data:       nil,
	}
}

func (s *OrganizationServiceTest) TestFindAllOrganizationService() {
	want := &proto.OrganizationPagination{
		Items: s.Organizations,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	client := new(mock.OrganizationClientMock)

	client.On("FindAll").Return(&proto.OrganizationPaginationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       want,
	}, nil)

	srv := service.NewOrganizationService(client)

	organizations, err := srv.FindAll(&dto.PaginationQueryParams{Limit: 10, Page: 1})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, organizations)
}

func (s *OrganizationServiceTest) TestFindAllGrpcErrOrganizationService() {
	want := s.ServiceDownErr

	client := new(mock.OrganizationClientMock)

	client.On("FindAll").Return(&proto.OrganizationPaginationResponse{}, errors.New("Service is down"))

	srv := service.NewOrganizationService(client)

	_, err := srv.FindAll(&dto.PaginationQueryParams{Limit: 10, Page: 1})

	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestFindOneOrganizationService() {
	want := s.Organization

	client := new(mock.OrganizationClientMock)

	client.On("FindOne").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Organization,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.FindOne(1)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, organization)
}

func (s *OrganizationServiceTest) TestFindOneNotFoundOrganizationService() {
	want := s.NotFoundErr

	client := new(mock.OrganizationClientMock)

	client.On("FindOne").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.FindOne(1)

	assert.Nil(s.T(), organization)
	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestFindOneGrpcErrOrganizationService() {
	want := s.ServiceDownErr

	client := new(mock.OrganizationClientMock)

	client.On("FindOne").Return(&proto.OrganizationResponse{}, errors.New("Service is down"))

	srv := service.NewOrganizationService(client)

	_, err := srv.FindOne(1)

	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestCreateOrganizationService() {
	want := s.Organization

	client := new(mock.OrganizationClientMock)

	client.On("Create").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       s.Organization,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Create(&dto.OrganizationDto{})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, organization)
}

func (s *OrganizationServiceTest) TestCreateDuplicatedOrganizationService() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated email or organizationname",
		Data:       nil,
	}

	client := new(mock.OrganizationClientMock)

	client.On("Create").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated email or organizationname"},
		Data:       nil,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Create(&dto.OrganizationDto{})

	assert.Nil(s.T(), organization)
	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestCreateGrpcErrOrganizationService() {
	want := s.ServiceDownErr

	client := new(mock.OrganizationClientMock)

	client.On("Create").Return(&proto.OrganizationResponse{}, errors.New("Service is down"))

	srv := service.NewOrganizationService(client)

	_, err := srv.Create(&dto.OrganizationDto{})

	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestUpdateOrganizationService() {
	want := s.Organization

	client := new(mock.OrganizationClientMock)

	client.On("Update").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Organization,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Update(1, &dto.OrganizationDto{})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, organization)
}

func (s *OrganizationServiceTest) TestUpdateNotFoundOrganizationService() {
	want := s.NotFoundErr

	client := new(mock.OrganizationClientMock)

	client.On("Update").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Update(1, &dto.OrganizationDto{})

	assert.Nil(s.T(), organization)
	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestUpdateGrpcErrOrganizationService() {
	want := s.ServiceDownErr

	client := new(mock.OrganizationClientMock)

	client.On("Update").Return(&proto.OrganizationResponse{}, errors.New("Service is down"))

	srv := service.NewOrganizationService(client)

	_, err := srv.Update(1, &dto.OrganizationDto{})

	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestDeleteOrganizationService() {
	want := s.Organization

	client := new(mock.OrganizationClientMock)

	client.On("Delete").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Organization,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Delete(1)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, organization)
}

func (s *OrganizationServiceTest) TestDeleteNotFoundOrganizationService() {
	want := s.NotFoundErr

	client := new(mock.OrganizationClientMock)

	client.On("Delete").Return(&proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil)

	srv := service.NewOrganizationService(client)

	organization, err := srv.Delete(1)

	assert.Nil(s.T(), organization)
	assert.Equal(s.T(), want, err)
}

func (s *OrganizationServiceTest) TestDeleteGrpcErrOrganizationService() {
	want := s.ServiceDownErr

	client := new(mock.OrganizationClientMock)

	client.On("Delete").Return(&proto.OrganizationResponse{}, errors.New("Service is down"))

	srv := service.NewOrganizationService(client)

	_, err := srv.Delete(1)

	assert.Equal(s.T(), want, err)
}
