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

func (s *OrganizationService) Create(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Organization

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Create(ctx, &proto.CreateOrganizationRequest{Organization: &user})
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

func (s *OrganizationService) Update(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Organization

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Update(ctx, &proto.UpdateOrganizationRequest{Organization: &user})
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

func (s *OrganizationService) Delete(c OrganizationContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.Organization

	err := c.Bind(&user)
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
