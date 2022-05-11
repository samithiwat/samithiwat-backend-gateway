package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
	"time"
)

type TeamService struct {
	client proto.TeamServiceClient
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}

func (s *TeamService) FindAll(query dto.PaginationQueryParams) (result *proto.TeamPagination, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.FindAllTeamRequest{
		Page:  query.Page,
		Limit: query.Limit,
	}

	res, errRes := s.client.FindAll(ctx, req)
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *TeamService) FindOne(id int32) (result *proto.Team, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.FindOne(ctx, &proto.FindOneTeamRequest{Id: id})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *TeamService) Create(teamDto *dto.TeamDto) (result *proto.Team, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	team := s.DtoToRaw(teamDto)

	res, errRes := s.client.Create(ctx, &proto.CreateTeamRequest{Team: team})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusCreated {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *TeamService) Update(id int32, teamDto *dto.TeamDto) (result *proto.Team, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	team := s.DtoToRaw(teamDto)
	team.Id = uint32(id)

	res, errRes := s.client.Update(ctx, &proto.UpdateTeamRequest{Team: team})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (s *TeamService) Delete(id int32) (result *proto.Team, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.Delete(ctx, &proto.DeleteTeamRequest{Id: id})
	if errRes != nil {
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data

	return
}

func (TeamService) DtoToRaw(teamDto *dto.TeamDto) *proto.Team {
	return &proto.Team{
		Name:        teamDto.Name,
		Description: teamDto.Description,
	}
}
