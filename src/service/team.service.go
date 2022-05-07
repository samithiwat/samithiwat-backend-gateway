package service

import "github.com/samithiwat/samithiwat-backend-gateway/src/proto"

type TeamService struct {
	client proto.TeamServiceClient
}

type TeamContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}
