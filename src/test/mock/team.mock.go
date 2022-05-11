package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type TeamContextMock struct {
	mock.Mock
	V interface{}
}

func (c *TeamContextMock) Bind(v interface{}) error {
	*v.(*dto.TeamDto) = dto.TeamDto{
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	args := c.Called()

	return args.Error(0)
}

func (c *TeamContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *TeamContextMock) ID(id *int32) error {
	*id = 1

	args := c.Called()

	return args.Error(0)
}

func (c *TeamContextMock) PaginationQueryParam(*dto.PaginationQueryParams) error {
	args := c.Called()

	return args.Error(0)
}

type TeamServiceMock struct {
	mock.Mock
}

func (s *TeamServiceMock) FindAll(*dto.PaginationQueryParams) (*proto.TeamPagination, *dto.ResponseErr) {
	args := s.Called()

	return args.Get(0).(*proto.TeamPagination), args.Get(1).(*dto.ResponseErr)
}

func (s *TeamServiceMock) FindOne(id int32) (*proto.Team, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Team), args.Get(1).(*dto.ResponseErr)
}

func (s *TeamServiceMock) Create(*dto.TeamDto) (*proto.Team, *dto.ResponseErr) {
	args := s.Called()

	return args.Get(0).(*proto.Team), args.Get(1).(*dto.ResponseErr)
}

func (s *TeamServiceMock) Update(id int32, _ *dto.TeamDto) (*proto.Team, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Team), args.Get(1).(*dto.ResponseErr)
}

func (s *TeamServiceMock) Delete(id int32) (*proto.Team, *dto.ResponseErr) {
	args := s.Called(id)

	return args.Get(0).(*proto.Team), args.Get(1).(*dto.ResponseErr)
}

type TeamClientMock struct {
	mock.Mock
}

func (c TeamClientMock) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.TeamPaginationResponse), args.Error(1)
}

func (c TeamClientMock) FindOne(ctx context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.TeamResponse), args.Error(1)
}

func (c TeamClientMock) FindMulti(ctx context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (c TeamClientMock) Create(ctx context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.TeamResponse), args.Error(1)
}

func (c TeamClientMock) Update(ctx context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.TeamResponse), args.Error(1)
}

func (c TeamClientMock) Delete(ctx context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	args := c.Called()

	return args.Get(0).(*proto.TeamResponse), args.Error(1)
}
