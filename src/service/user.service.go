package service

import (
	"context"
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
	UserId() uint
}

func (s *UserService) FindOne(c UserContext) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := s.client.FindOne(ctx, &proto.FindOneUserRequest{Id: int32(c.UserId())})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ID":        res.Data.Id,
		"Firstname": res.Data.Firstname,
		"Lastname":  res.Data.Lastname,
		"ImageUrl":  res.Data.ImageUrl,
	})
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
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusCreated {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ID":        res.Data.Id,
		"Firstname": res.Data.Firstname,
		"Lastname":  res.Data.Lastname,
		"ImageUrl":  res.Data.ImageUrl,
	})
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

	res, err := s.client.Update(ctx, &proto.UpdateUserRequest{User: &user})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ID":        res.Data.Id,
		"Firstname": res.Data.Firstname,
		"Lastname":  res.Data.Lastname,
		"ImageUrl":  res.Data.ImageUrl,
	})
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

	res, err := s.client.Delete(ctx, &proto.DeleteUserRequest{Id: int32(c.UserId())})
	if err != nil {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": http.StatusBadGateway,
			"Message":    "Service is down",
		})
		return
	}

	if res.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadGateway, map[string]interface{}{
			"StatusCode": res.StatusCode,
			"Message":    res.Errors,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"ID":        res.Data.Id,
		"Firstname": res.Data.Firstname,
		"Lastname":  res.Data.Lastname,
		"ImageUrl":  res.Data.ImageUrl,
	})
	return
}
