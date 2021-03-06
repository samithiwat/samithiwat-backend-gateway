package handler

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	validate "github.com/samithiwat/samithiwat-backend-gateway/src/validator"
	"net/http"
)

type TeamHandler struct {
	service  TeamService
	validate *validate.DtoValidator
}

func NewTeamHandler(service TeamService, validate *validate.DtoValidator) *TeamHandler {
	return &TeamHandler{
		service:  service,
		validate: validate,
	}
}

type TeamContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID() (int32, error)
	PaginationQueryParam(*dto.PaginationQueryParams) error
}

type TeamService interface {
	FindAll(*dto.PaginationQueryParams) (*proto.TeamPagination, *dto.ResponseErr)
	FindOne(int32) (*proto.Team, *dto.ResponseErr)
	Create(*dto.TeamDto) (*proto.Team, *dto.ResponseErr)
	Update(int32, *dto.TeamDto) (*proto.Team, *dto.ResponseErr)
	Delete(int32) (*proto.Team, *dto.ResponseErr)
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
// @Failure 400 {object} dto.ResponseErr "Invalid query param"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /team [get]
func (h *TeamHandler) FindAll(c TeamContext) {
	query := dto.PaginationQueryParams{}

	err := c.PaginationQueryParam(&query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Cannot parse query param",
		})
		return
	}

	teams, errRes := h.service.FindAll(&query)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, teams)
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
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found team"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Router /team/{id} [get]
func (h *TeamHandler) FindOne(c TeamContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	team, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, team)
	return
}

// Create is a function that create the team
// @Summary Create the team
// @Description Return the team dto if successfully
// @Param team body dto.TeamDto true "team dto"
// @Tags team
// @Accept json
// @Produce json
// @Success 201 {object} proto.Team
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found team"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /team [post]
func (h *TeamHandler) Create(c TeamContext) {
	teamDto := dto.TeamDto{}
	err := c.Bind(&teamDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse team dto",
		})
		return
	}

	if errors := h.validate.Validate(teamDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	team, errRes := h.service.Create(&teamDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, team)
	return
}

// Update is a function that update the team
// @Summary Update the existing team
// @Description Return the team dto if successfully
// @Param id path int true "id"
// @Param team body dto.TeamDto true "team dto"
// @Tags team
// @Accept json
// @Produce json
// @Success 200 {object} proto.Team
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found team"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /team/{id} [patch]
func (h *TeamHandler) Update(c TeamContext) {
	teamDto := dto.TeamDto{}
	err := c.Bind(&teamDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Cannot parse team dto",
		})
		return
	}

	if errors := h.validate.Validate(teamDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	team, errRes := h.service.Update(id, &teamDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, team)
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
// @Failure 400 {object} dto.ResponseErr "Invalid ID"
// @Failure 404 {object} dto.ResponseErr "Not found team"
// @Failure 503 {object} dto.ResponseErr "Service is down"
// @Security     AuthToken
// @Router /team/{id} [delete]
func (h *TeamHandler) Delete(c TeamContext) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		})
		return
	}

	team, errRes := h.service.Delete(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, team)
	return
}
