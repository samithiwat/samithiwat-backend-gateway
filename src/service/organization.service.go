package service

import "github.com/samithiwat/samithiwat-backend-gateway/src/proto"

type OrganizationService struct {
	client proto.OrganizationServiceClient
}

type OrganizationContext interface {
	Bind(interface{}) error
	JSON(int, interface{})
}

func NewOrganizationService(client proto.OrganizationServiceClient) *OrganizationService {
	return &OrganizationService{
		client: client,
	}
}
