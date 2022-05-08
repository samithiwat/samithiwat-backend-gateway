package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Team1 proto.Team
var Team2 proto.Team
var Team3 proto.Team
var Team4 proto.Team
var Teams []*proto.Team

type TeamMockClient struct {
}

func (u TeamMockClient) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	return &proto.TeamPaginationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data: &proto.TeamPagination{
			Items: Teams,
			Meta: &proto.PaginationMetadata{
				TotalItem:    4,
				ItemCount:    4,
				ItemsPerPage: 10,
				TotalPage:    1,
				CurrentPage:  1,
			},
		},
	}, nil
}

func (u TeamMockClient) FindOne(ctx context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Team1,
	}, nil
}

func (u TeamMockClient) FindMulti(ctx context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (u TeamMockClient) Create(ctx context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       &Team1,
	}, nil
}

func (u TeamMockClient) Update(ctx context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Team1,
	}, nil
}

func (u TeamMockClient) Delete(ctx context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Team1,
	}, nil
}

type TeamMockErrClient struct {
}

func (u TeamMockErrClient) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	return nil, nil
}

func (u TeamMockErrClient) FindOne(ctx context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

func (u TeamMockErrClient) FindMulti(ctx context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (u TeamMockErrClient) Create(ctx context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated username or email"},
		Data:       nil,
	}, nil
}

func (u TeamMockErrClient) Update(ctx context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

func (u TeamMockErrClient) Delete(ctx context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return &proto.TeamResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found user"},
		Data:       nil,
	}, nil
}

type TeamMockErrGrpcClient struct {
}

func (u TeamMockErrGrpcClient) FindAll(ctx context.Context, in *proto.FindAllTeamRequest, opts ...grpc.CallOption) (*proto.TeamPaginationResponse, error) {
	return nil, errors.New("Service is down")
}

func (u TeamMockErrGrpcClient) FindOne(ctx context.Context, in *proto.FindOneTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, errors.New("Service is down")
}

func (u TeamMockErrGrpcClient) FindMulti(ctx context.Context, in *proto.FindMultiTeamRequest, opts ...grpc.CallOption) (*proto.TeamListResponse, error) {
	return nil, nil
}

func (u TeamMockErrGrpcClient) Create(ctx context.Context, in *proto.CreateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, errors.New("Service is down")
}

func (u TeamMockErrGrpcClient) Update(ctx context.Context, in *proto.UpdateTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, errors.New("Service is down")
}

func (u TeamMockErrGrpcClient) Delete(ctx context.Context, in *proto.DeleteTeamRequest, opts ...grpc.CallOption) (*proto.TeamResponse, error) {
	return nil, errors.New("Service is down")
}

type TeamMockContext struct {
	V interface{}
}

func (TeamMockContext) Bind(v interface{}) error {
	*v.(*proto.Team) = Team1
	return nil
}

func (c *TeamMockContext) JSON(_ int, v interface{}) {
	c.V = v
}

func (TeamMockContext) TeamID() uint {
	return 1
}

func (TeamMockContext) QueryParam() *proto.FindAllTeamRequest {
	return &proto.FindAllTeamRequest{
		Page:  1,
		Limit: 10,
	}
}

func InitializeMockTeam() {
	Team1 = proto.Team{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team2 = proto.Team{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team3 = proto.Team{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Team4 = proto.Team{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Teams = append(Teams, &Team1, &Team2, &Team3, &Team4)
}
