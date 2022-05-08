package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
	"time"
)

type TeamService struct {
	client proto.TeamServiceClient
}

type TeamContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TeamID() uint
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}

func (s *TeamService) FindAll(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &proto.FindAllTeamRequest{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, res.Data)
	return
}

func (s *TeamService) FindOne(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.client.FindOne(ctx, &proto.FindOneTeamRequest{Id: int32(c.TeamID())})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, res.Data)
	return
}

func (s *TeamService) Create(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Team

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Create(ctx, &proto.CreateTeamRequest{Team: &user})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusCreated {
		c.JSON(int(res.StatusCode), map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusCreated, res.Data)
	return
}

func (s *TeamService) Update(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Team

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Update(ctx, &proto.UpdateTeamRequest{Team: &user})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, res.Data)
	return
}

func (s *TeamService) Delete(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Team

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Delete(ctx, &proto.DeleteTeamRequest{Id: int32(c.TeamID())})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(int(res.StatusCode), map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, res.Data)
	return
}
