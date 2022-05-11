package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type OrganizationContextMock struct {
	mock.Mock
	V interface{}
}

func (c *OrganizationContextMock) Bind(v interface{}) error {
	*v.(*dto.OrganizationDto) = dto.OrganizationDto{
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	args := c.Called()

	return args.Error(0)
}

func (c *OrganizationContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *OrganizationContextMock) ID(id *int32) error {
	*id = 1

	args := c.Called()

	return args.Error(0)
}

func (c *OrganizationContextMock) PaginationQueryParam(*dto.PaginationQueryParams) error {
	args := c.Called()

	return args.Error(0)
}

type OrganizationServiceMock struct {
	mock.Mock
}

func (s *OrganizationServiceMock) FindAll(*dto.PaginationQueryParams) (*proto.OrganizationPagination, *dto.ResponseErr) {
	args := s.Called()

	return args.Get(0).(*proto.OrganizationPagination), args.Get(1).(*dto.ResponseErr)
}

func (s *OrganizationServiceMock) FindOne(id int32) (*proto.Organization, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Organization), args.Get(1).(*dto.ResponseErr)
}

func (s *OrganizationServiceMock) Create(*dto.OrganizationDto) (*proto.Organization, *dto.ResponseErr) {
	args := s.Called()

	return args.Get(0).(*proto.Organization), args.Get(1).(*dto.ResponseErr)
}

func (s *OrganizationServiceMock) Update(id int32, _ *dto.OrganizationDto) (*proto.Organization, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Organization), args.Get(1).(*dto.ResponseErr)
}

func (s *OrganizationServiceMock) Delete(id int32) (*proto.Organization, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Organization), args.Get(1).(*dto.ResponseErr)
}

type OrganizationClientMock struct {
	mock.Mock
}

func (c *OrganizationClientMock) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.OrganizationPaginationResponse), args.Error(1)
}

func (c *OrganizationClientMock) FindOne(ctx context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.OrganizationResponse), args.Error(1)
}

func (c *OrganizationClientMock) FindMulti(ctx context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (c *OrganizationClientMock) Create(ctx context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.OrganizationResponse), args.Error(1)
}

func (c *OrganizationClientMock) Update(ctx context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.OrganizationResponse), args.Error(1)
}

func (c *OrganizationClientMock) Delete(ctx context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.OrganizationResponse), args.Error(1)
}
