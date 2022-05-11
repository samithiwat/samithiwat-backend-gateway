package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"net/http"
	"time"
)

type OrganizationService struct {
	client proto.OrganizationServiceClient
}

func NewOrganizationService(client proto.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}

func (s *OrganizationService) FindAll(query *dto.PaginationQueryParams) (result *proto.OrganizationPagination, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto.FindAllOrganizationRequest{
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

func (s *OrganizationService) FindOne(id int32) (result *proto.Organization, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.FindOne(ctx, &proto.FindOneOrganizationRequest{Id: id})
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

func (s *OrganizationService) Create(organizationDto *dto.OrganizationDto) (result *proto.Organization, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	organization := s.DtoToRaw(organizationDto)

	res, errRes := s.client.Create(ctx, &proto.CreateOrganizationRequest{Organization: organization})
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

func (s *OrganizationService) Update(id int32, organizationDto *dto.OrganizationDto) (result *proto.Organization, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	organization := s.DtoToRaw(organizationDto)
	organization.Id = uint32(id)

	res, errRes := s.client.Update(ctx, &proto.UpdateOrganizationRequest{Organization: organization})
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

func (s *OrganizationService) Delete(id int32) (result *proto.Organization, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.Delete(ctx, &proto.DeleteOrganizationRequest{Id: id})
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

func (OrganizationService) DtoToRaw(organizationDto *dto.OrganizationDto) *proto.Organization {
	return &proto.Organization{
		Name:        organizationDto.Name,
		Description: organizationDto.Description,
	}
}
