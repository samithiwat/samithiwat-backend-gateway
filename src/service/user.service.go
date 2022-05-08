package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/model"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
	"time"
)

type UserService struct {
	client proto.UserServiceClient
}

func NewUserService(client proto.UserServiceClient) *UserService {
	return &UserService{
		client: client,
	}
}

type UserContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
	ID(*int32) error
	PaginationQueryParam(*model.PaginationQueryParams) error
}

// FindAllUser is a function that get all users in database
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Book}
// @Failure 503 {object} ResponseHTTP{}
// @Router /user [get]
func (s *UserService) FindAll(c UserContext) {
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

	req := &proto.FindAllUserRequest{
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

func (s *UserService) FindOne(c UserContext) {
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

	res, err := s.client.FindOne(ctx, &proto.FindOneUserRequest{Id: id})
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

func (s *UserService) Create(c UserContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.User

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"StatusCode": http.StatusBadRequest,
			"Message":    "Invalid request body",
		})
		return
	}

	res, err := s.client.Create(ctx, &proto.CreateUserRequest{User: &user})
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

func (s *UserService) Update(c UserContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.User

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

	user.Id = uint32(id)

	res, err := s.client.Update(ctx, &proto.UpdateUserRequest{User: &user})
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

func (s *UserService) Delete(c UserContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user proto.User

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

	res, err := s.client.Delete(ctx, &proto.DeleteUserRequest{Id: id})
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
