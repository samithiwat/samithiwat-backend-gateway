package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/model"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
	"time"
)

type OrganizationService struct {
	client proto.OrganizationServiceClient
}

type OrganizationContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID(*int32) error
	PaginationQueryParam(*model.PaginationQueryParams) error
}

func NewOrganizationService(client proto.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}

// FindAll is a function that get all organizations in database
// @Summary Get all organizations
// @Description Return the arrays of organization dto if successfully
// @Param limit query int false "Limit"
// @Param page query int false "Page"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} model.ResponseErr "Invalid query param"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /organization [get]
func (s *OrganizationService) FindAll(c OrganizationContext) {
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

	req := &proto.FindAllOrganizationRequest{
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

// FindOne is a function that get the specific organizations with id
// @Summary Get specific organization with id
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found organization"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /organization/{id} [get]
func (s *OrganizationService) FindOne(c OrganizationContext) {
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

	res, err := s.client.FindOne(ctx, &proto.FindOneOrganizationRequest{Id: id})
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

// Create is a function that create the organization
// @Summary Create the organization
// @Description Return the organization dto if successfully
// @Param organization body proto.Organization true "organization dto"
// @Tags organization
// @Accept json
// @Produce json
// @Success 201 {object} proto.Organization
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found organization"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /organization [post]
func (s *OrganizationService) Create(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var organization proto.Organization

	err := c.Bind(&organization)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Create(ctx, &proto.CreateOrganizationRequest{Organization: &organization})
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

// Update is a function that update the organization
// @Summary Update the existing organization
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Param organization body proto.Organization true "organization dto"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found organization"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /organization/{id} [patch]
func (s *OrganizationService) Update(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var organization proto.Organization

	err := c.Bind(&organization)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Update(ctx, &proto.UpdateOrganizationRequest{Organization: &organization})
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

// Delete is a function that delete the organization
// @Summary Delete the organization
// @Description Return the organization dto if successfully
// @Param id path int true "id"
// @Tags organization
// @Accept json
// @Produce json
// @Success 200 {object} proto.Organization
// @Failure 400 {object} model.ResponseErr "Invalid ID"
// @Failure 404 {object} model.ResponseErr "Not found organization"
// @Failure 503 {object} model.ResponseErr "Service is down"
// @Router /organization/{id} [delete]
func (s *OrganizationService) Delete(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var organization proto.Organization

	err := c.Bind(&organization)
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

	res, err := s.client.Delete(ctx, &proto.DeleteOrganizationRequest{Id: id})
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
