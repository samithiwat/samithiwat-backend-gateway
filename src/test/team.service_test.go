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

type TeamServiceTest struct {
	suite.Suite
	Team           *proto.Team
	Teams          []*proto.Team
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestTeamService(t *testing.T) {
	suite.Run(t, new(TeamServiceTest))
}

func (s *TeamServiceTest) SetupTest() {
	s.Team = &proto.Team{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team2 := &proto.Team{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team3 := &proto.Team{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team4 := &proto.Team{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	s.Teams = append(s.Teams, s.Team, Team2, Team3, Team4)

	s.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	s.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found team",
		Data:       nil,
	}
}

func (s *TeamServiceTest) TestFindAllTeamService() {
	want := &proto.TeamPagination{
		Items: s.Teams,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	client := new(mock.TeamClientMock)

	client.On("FindAll").Return(&proto.TeamPaginationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       want,
	}, nil)

	srv := service.NewTeamService(client)

	teams, err := srv.FindAll(&dto.PaginationQueryParams{Limit: 10, Page: 1})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, teams)
}

func (s *TeamServiceTest) TestFindAllGrpcErrTeamService() {
	want := s.ServiceDownErr

	client := new(mock.TeamClientMock)

	client.On("FindAll").Return(&proto.TeamPaginationResponse{}, errors.New("Service is down"))

	srv := service.NewTeamService(client)

	_, err := srv.FindAll(&dto.PaginationQueryParams{Limit: 10, Page: 1})

	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestFindOneTeamService() {
	want := s.Team

	client := new(mock.TeamClientMock)

	client.On("FindOne").Return(&proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Team,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.FindOne(1)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, team)
}

func (s *TeamServiceTest) TestFindOneNotFoundTeamService() {
	want := s.NotFoundErr

	client := new(mock.TeamClientMock)

	client.On("FindOne").Return(&proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found team"},
		Data:       nil,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.FindOne(1)

	assert.Nil(s.T(), team)
	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestFindOneGrpcErrTeamService() {
	want := s.ServiceDownErr

	client := new(mock.TeamClientMock)

	client.On("FindOne").Return(&proto.TeamResponse{}, errors.New("Service is down"))

	srv := service.NewTeamService(client)

	_, err := srv.FindOne(1)

	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestCreateTeamService() {
	want := s.Team

	client := new(mock.TeamClientMock)

	client.On("Create").Return(&proto.TeamResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       s.Team,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Create(&dto.TeamDto{})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, team)
}

func (s *TeamServiceTest) TestCreateDuplicatedTeamService() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated email or teamname",
		Data:       nil,
	}

	client := new(mock.TeamClientMock)

	client.On("Create").Return(&proto.TeamResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated email or teamname"},
		Data:       nil,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Create(&dto.TeamDto{})

	assert.Nil(s.T(), team)
	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestCreateGrpcErrTeamService() {
	want := s.ServiceDownErr

	client := new(mock.TeamClientMock)

	client.On("Create").Return(&proto.TeamResponse{}, errors.New("Service is down"))

	srv := service.NewTeamService(client)

	_, err := srv.Create(&dto.TeamDto{})

	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestUpdateTeamService() {
	want := s.Team

	client := new(mock.TeamClientMock)

	client.On("Update").Return(&proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Team,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Update(1, &dto.TeamDto{})

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, team)
}

func (s *TeamServiceTest) TestUpdateNotFoundTeamService() {
	want := s.NotFoundErr

	client := new(mock.TeamClientMock)

	client.On("Update").Return(&proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found team"},
		Data:       nil,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Update(1, &dto.TeamDto{})

	assert.Nil(s.T(), team)
	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestUpdateGrpcErrTeamService() {
	want := s.ServiceDownErr

	client := new(mock.TeamClientMock)

	client.On("Update").Return(&proto.TeamResponse{}, errors.New("Service is down"))

	srv := service.NewTeamService(client)

	_, err := srv.Update(1, &dto.TeamDto{})

	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestDeleteTeamService() {
	want := s.Team

	client := new(mock.TeamClientMock)

	client.On("Delete").Return(&proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       s.Team,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Delete(1)

	assert.Nil(s.T(), err, "Must not got any error")
	assert.Equal(s.T(), want, team)
}

func (s *TeamServiceTest) TestDeleteNotFoundTeamService() {
	want := s.NotFoundErr

	client := new(mock.TeamClientMock)

	client.On("Delete").Return(&proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found team"},
		Data:       nil,
	}, nil)

	srv := service.NewTeamService(client)

	team, err := srv.Delete(1)

	assert.Nil(s.T(), team)
	assert.Equal(s.T(), want, err)
}

func (s *TeamServiceTest) TestDeleteGrpcErrTeamService() {
	want := s.ServiceDownErr

	client := new(mock.TeamClientMock)

	client.On("Delete").Return(&proto.TeamResponse{}, errors.New("Service is down"))

	srv := service.NewTeamService(client)

	_, err := srv.Delete(1)

	assert.Equal(s.T(), want, err)
}
