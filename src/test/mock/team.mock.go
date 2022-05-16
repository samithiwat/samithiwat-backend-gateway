package mock

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type TeamContextMock struct {
	mock.Mock
	V       interface{}
	Team    *proto.Team
	Teams   []*proto.Team
	TeamDto *dto.TeamDto
	Query   *dto.PaginationQueryParams
}

func (c *TeamContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.TeamDto) = *c.TeamDto

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

func (c *TeamContextMock) PaginationQueryParam(query *dto.PaginationQueryParams) error {
	args := c.Called(query)

	*query = *c.Query

	return args.Error(0)
}

type TeamServiceMock struct {
	mock.Mock
}

func (s *TeamServiceMock) FindAll(query *dto.PaginationQueryParams) (res *proto.TeamPagination, err *dto.ResponseErr) {
	args := s.Called(query)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamPagination)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *TeamServiceMock) FindOne(id int32) (res *proto.Team, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Team)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *TeamServiceMock) Create(team *dto.TeamDto) (res *proto.Team, err *dto.ResponseErr) {
	args := s.Called(team)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Team)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *TeamServiceMock) Update(id int32, team *dto.TeamDto) (res *proto.Team, err *dto.ResponseErr) {
	args := s.Called(id, team)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Team)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *TeamServiceMock) Delete(id int32) (res *proto.Team, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Team)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type TeamClientMock struct {
	mock.Mock
}

func (c TeamClientMock) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (res *proto.TeamPaginationResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamPaginationResponse)
	}

	return res, args.Error(1)
}

func (c TeamClientMock) FindOne(ctx context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (res *proto.TeamResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamResponse)
	}

	return res, args.Error(1)
}

func (c TeamClientMock) FindMulti(ctx context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (c TeamClientMock) Create(ctx context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (res *proto.TeamResponse, err error) {
	args := c.Called(*in.Team)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamResponse)
	}

	return res, args.Error(1)
}

func (c TeamClientMock) Update(ctx context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (res *proto.TeamResponse, err error) {
	args := c.Called(*in.Team)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamResponse)
	}

	return res, args.Error(1)
}

func (c TeamClientMock) Delete(ctx context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (res *proto.TeamResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.TeamResponse)
	}

	return res, args.Error(1)
}
