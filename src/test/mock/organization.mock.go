package mock

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type OrganizationContextMock struct {
	mock.Mock
	V               interface{}
	Organization    *proto.Organization
	Organizations   []*proto.Organization
	OrganizationDto *dto.OrganizationDto
	Query           *dto.PaginationQueryParams
}

func (c *OrganizationContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.OrganizationDto) = *c.OrganizationDto

	return args.Error(0)
}

func (c *OrganizationContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *OrganizationContextMock) ID(id *int32) error {
	args := c.Called(*id)

	*id = 1

	return args.Error(0)
}

func (c *OrganizationContextMock) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	args := c.Called(query)

	*query = *c.Query

	return args.Error(0)
}

type OrganizationServiceMock struct {
	mock.Mock
}

func (s *OrganizationServiceMock) FindAll(query *dto.PaginationQueryParams) (res *proto.OrganizationPagination, err *dto.ResponseErr) {
	args := s.Called(query)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationPagination)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *OrganizationServiceMock) FindOne(id int32) (res *proto.Organization, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Organization)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *OrganizationServiceMock) Create(org *dto.OrganizationDto) (res *proto.Organization, err *dto.ResponseErr) {
	args := s.Called(org)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Organization)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *OrganizationServiceMock) Update(id int32, org *dto.OrganizationDto) (res *proto.Organization, err *dto.ResponseErr) {
	args := s.Called(id, org)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Organization)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *OrganizationServiceMock) Delete(id int32) (res *proto.Organization, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Organization)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type OrganizationClientMock struct {
	mock.Mock
}

func (c *OrganizationClientMock) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (res *proto.OrganizationPaginationResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationPaginationResponse)
	}

	return res, args.Error(1)
}

func (c *OrganizationClientMock) FindOne(ctx context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (res *proto.OrganizationResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationResponse)
	}

	return res, args.Error(1)
}

func (c *OrganizationClientMock) FindMulti(ctx context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (c *OrganizationClientMock) Create(ctx context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (res *proto.OrganizationResponse, err error) {
	args := c.Called(*in.Organization)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationResponse)
	}

	return res, args.Error(1)
}

func (c *OrganizationClientMock) Update(ctx context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (res *proto.OrganizationResponse, err error) {
	args := c.Called(*in.Organization)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationResponse)
	}

	return res, args.Error(1)
}

func (c *OrganizationClientMock) Delete(ctx context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (res *proto.OrganizationResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.OrganizationResponse)
	}

	return res, args.Error(1)
}
