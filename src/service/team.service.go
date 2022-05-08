package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/model"
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
	ID(*int32) error
	PaginationQueryParam(*model.PaginationQueryParams) error
}

func NewTeamService(client proto.TeamServiceClient) *TeamService {
	return &TeamService{
		client: client,
	}
}

// FindAll is a function that get all teams in database
// @Summary Get all teams
// @Description Return the arrays of team dto if successfully
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} proto.Team
// @Failure 400 {object} model.ResponseErr "Invalid query param"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /team [get]
func (s *TeamService) FindAll(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := &model.PaginationQueryParams{}

	err := c.PaginationQueryParam(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid query param",
		})
		return
	}

	req := &proto.FindAllTeamRequest{
		Page:  query.Page,
		Limit: query.Limit,
	}

	res, err := s.client.FindAll(ctx, req)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, map[string]interface{}{
			"StatusCode": http.StatusServiceUnavailable,
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

// FindOne is a function that get the specific teams with id
// @Summary Get specific team with id
// @Description Return the team dto if successfully
// @Param id path int true "id"
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} proto.Team
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found team"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /team/{id} [get]
func (s *TeamService) FindOne(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var id int32
	err := c.ID(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid id",
		})
		return
	}

	res, err := s.client.FindOne(ctx, &proto.FindOneTeamRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, map[string]interface{}{
			"StatusCode": http.StatusServiceUnavailable,
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

// Create is a function that create the team
// @Summary Create the team
// @Description Return the team dto if successfully
// @Param team body proto.Team true "team dto"
// @Tags team
// @Accept json
// @Produce json
// @Success 201 {object} proto.Team
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found team"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /team [post]
func (s *TeamService) Create(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var team proto.Team

	err := c.Bind(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Create(ctx, &proto.CreateTeamRequest{Team: &team})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, map[string]interface{}{
			"StatusCode": http.StatusServiceUnavailable,
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

// Update is a function that update the team
// @Summary Update the existing team
// @Description Return the team dto if successfully
// @Param id path int true "id"
// @Param team body proto.Team true "team dto"
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} proto.Team
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found team"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /team/{id} [patch]
func (s *TeamService) Update(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var team proto.Team

	err := c.Bind(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Update(ctx, &proto.UpdateTeamRequest{Team: &team})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, map[string]interface{}{
			"StatusCode": http.StatusServiceUnavailable,
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

// Delete is a function that delete the team
// @Summary Delete the team
// @Description Return the team dto if successfully
// @Param id path int true "id"
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} proto.Team
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found team"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /team/{id} [delete]
func (s *TeamService) Delete(c TeamContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var team proto.Team

	err := c.Bind(&team)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	var id int32
	err = c.ID(&id)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid id",
		})
		return
	}

	res, err := s.client.Delete(ctx, &proto.DeleteTeamRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, map[string]interface{}{
			"StatusCode": http.StatusServiceUnavailable,
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
