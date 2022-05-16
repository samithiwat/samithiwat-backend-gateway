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

type TeamHandlerTest struct {
	suite.Suite
	Team           *proto.Team
	Teams          []*proto.Team
	TeamDto        *dto.TeamDto
	Query          *dto.PaginationQueryParams
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
	InvalidIDErr   *dto.ResponseErr
}

func TestTeamHandler(t *testing.T) {
	suite.Run(t, new(TeamHandlerTest))
}

func (u *TeamHandlerTest) SetupTest() {
	u.Team = &proto.Team{
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

	_ = faker.FakeData(&u.TeamDto)
	_ = faker.FakeData(&u.Query)

	u.Teams = append(u.Teams, u.Team, Team2, Team3, Team4)

	u.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	u.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Not found team",
		Data:       nil,
	}

	u.InvalidIDErr = &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid ID",
	}
}

func (u *TeamHandlerTest) TestFindAllTeam() {
	want := &proto.TeamPagination{
		Items: u.Teams,
		Meta: &proto.PaginationMetadata{
			TotalItem:    4,
			ItemCount:    4,
			ItemsPerPage: 10,
			TotalPage:    1,
			CurrentPage:  1,
		},
	}

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(want, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindAllInvalidQueryParamTeam() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Cannot parse query param",
	}

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(nil, nil)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(errors.New("Cannot parse query param"))

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindAllGrpcErrTeam() {
	want := u.ServiceDownErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindAll", u.Query).Return(&proto.TeamPagination{}, u.ServiceDownErr)
	c.On("PaginationQueryParam", &dto.PaginationQueryParams{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.FindAll(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindOneTeam() {
	want := u.Team

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(u.Team, nil)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindOneInvalidRequestParamIDTeam() {
	want := u.InvalidIDErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Team{}, nil)
	c.On("ID").Return(errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindOneErrorNotFoundTeam() {
	want := u.NotFoundErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Team{}, u.NotFoundErr)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestFindOneGrpcErrTeam() {
	want := u.ServiceDownErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("FindOne", int32(1)).Return(&proto.Team{}, u.ServiceDownErr)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.FindOne(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestCreateTeam() {
	want := u.Team

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Create", u.TeamDto).Return(u.Team, nil)
	c.On("Bind", &dto.TeamDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestCreateErrorDuplicatedTeam() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    "Duplicated team name",
	}

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Create", u.TeamDto).Return(&proto.Team{}, want)
	c.On("Bind", &dto.TeamDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestCreateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse team dto",
	}

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Create", u.TeamDto).Return(&proto.Team{}, nil)
	c.On("Bind", &dto.TeamDto{}).Return(errors.New("Cannot parse body request"))

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestCreateGrpcErrTeam() {
	want := u.ServiceDownErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Create", u.TeamDto).Return(&proto.Team{}, u.ServiceDownErr)
	c.On("Bind", &dto.TeamDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestUpdateTeam() {
	want := u.Team

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.TeamDto).Return(u.Team, nil)
	c.On("Bind", &dto.TeamDto{}).Return(nil)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestUpdateInvalidRequestParamIDTeam() {
	want := u.InvalidIDErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.TeamDto).Return(&proto.Team{}, nil)
	c.On("ID").Return(errors.New("Invalid ID"))
	c.On("Bind", &dto.TeamDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestUpdateInvalidBodyRequest() {
	want := &dto.ResponseErr{
		StatusCode: http.StatusBadRequest,
		Message:    "Cannot parse team dto",
	}

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.TeamDto).Return(&proto.Team{}, nil)
	c.On("ID").Return(nil)
	c.On("Bind", &dto.TeamDto{}).Return(errors.New("Cannot parse team dto"))

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.Create(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestUpdateErrorNotFoundTeam() {
	want := u.NotFoundErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.TeamDto).Return(&proto.Team{}, u.NotFoundErr)
	c.On("ID").Return(nil)
	c.On("Bind", &dto.TeamDto{}).Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)
	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestUpdateGrpcErrTeam() {
	want := u.ServiceDownErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Update", int32(1), u.TeamDto).Return(&proto.Team{}, u.ServiceDownErr)
	c.On("Bind", &dto.TeamDto{}).Return(nil)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Update(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestDeleteTeam() {
	want := u.Team

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(u.Team, nil)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestDeleteInvalidRequestParamIDTeam() {
	want := u.InvalidIDErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Team{}, nil)
	c.On("ID").Return(errors.New("Invalid ID"))

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestDeleteErrorNotFoundTeam() {
	want := u.NotFoundErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Team{}, u.NotFoundErr)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}

func (u *TeamHandlerTest) TestDeleteGrpcErrTeam() {
	want := u.ServiceDownErr

	srv := new(mock.TeamServiceMock)
	c := &mock.TeamContextMock{
		Team:    u.Team,
		Teams:   u.Teams,
		TeamDto: u.TeamDto,
		Query:   u.Query,
	}

	srv.On("Delete", int32(1)).Return(&proto.Team{}, u.ServiceDownErr)
	c.On("ID").Return(nil)

	v, _ := validator.NewValidator()

	h := handler.NewTeamHandler(srv, v)

	h.Delete(c)

	assert.Equal(u.T(), want, c.V)
}
