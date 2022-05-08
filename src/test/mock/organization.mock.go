package mock

import (
	"context"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/samithiwat/samithiwat-backend-gateway/src/model"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"google.golang.org/grpc"
	"net/http"
)

var Organization1 proto.Organization
var Organization2 proto.Organization
var Organization3 proto.Organization
var Organization4 proto.Organization
var Organizations []*proto.Organization

type OrganizationMockClient struct {
}

func (u OrganizationMockClient) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	return &proto.OrganizationPaginationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data: &proto.OrganizationPagination{
			Items: Organizations,
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

func (u OrganizationMockClient) FindOne(ctx context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Organization1,
	}, nil
}

func (u OrganizationMockClient) FindMulti(ctx context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (u OrganizationMockClient) Create(ctx context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusCreated,
		Errors:     nil,
		Data:       &Organization1,
	}, nil
}

func (u OrganizationMockClient) Update(ctx context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Organization1,
	}, nil
}

func (u OrganizationMockClient) Delete(ctx context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusOK,
		Errors:     nil,
		Data:       &Organization1,
	}, nil
}

type OrganizationMockErrClient struct {
}

func (u OrganizationMockErrClient) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	return nil, nil
}

func (u OrganizationMockErrClient) FindOne(ctx context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil
}

func (u OrganizationMockErrClient) FindMulti(ctx context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (u OrganizationMockErrClient) Create(ctx context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     []string{"Duplicated organization name"},
		Data:       nil,
	}, nil
}

func (u OrganizationMockErrClient) Update(ctx context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil
}

func (u OrganizationMockErrClient) Delete(ctx context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return &proto.OrganizationResponse{
		StatusCode: http.StatusNotFound,
		Errors:     []string{"Not found organization"},
		Data:       nil,
	}, nil
}

type OrganizationMockErrGrpcClient struct {
}

func (u OrganizationMockErrGrpcClient) FindAll(ctx context.Context, in *proto.FindAllOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationPaginationResponse, error) {
	return nil, errors.New("Service is down")
}

func (u OrganizationMockErrGrpcClient) FindOne(ctx context.Context, in *proto.FindOneOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, errors.New("Service is down")
}

func (u OrganizationMockErrGrpcClient) FindMulti(ctx context.Context, in *proto.FindMultiOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationListResponse, error) {
	return nil, nil
}

func (u OrganizationMockErrGrpcClient) Create(ctx context.Context, in *proto.CreateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, errors.New("Service is down")
}

func (u OrganizationMockErrGrpcClient) Update(ctx context.Context, in *proto.UpdateOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, errors.New("Service is down")
}

func (u OrganizationMockErrGrpcClient) Delete(ctx context.Context, in *proto.DeleteOrganizationRequest, opts ...grpc.CallOption) (*proto.OrganizationResponse, error) {
	return nil, errors.New("Service is down")
}

type OrganizationMockContext struct {
	V interface{}
}

func (OrganizationMockContext) Bind(v interface{}) error {
	*v.(*proto.Organization) = Organization1
	return nil
}

func (c *OrganizationMockContext) JSON(_ int, v interface{}) {
	c.V = v
}

func (OrganizationMockContext) ID(id *int32) error {
	*id = 1
	return nil
}

func (OrganizationMockContext) PaginationQueryParam(query *model.PaginationQueryParams) error {
	*query = model.PaginationQueryParams{
		Page:  1,
		Limit: 10,
	}
	return nil
}

type OrganizationMockErrContext struct {
	V interface{}
}

func (OrganizationMockErrContext) Bind(v interface{}) error {
	*v.(*proto.Organization) = Organization1
	return nil
}

func (c *OrganizationMockErrContext) JSON(_ int, v interface{}) {
	c.V = v
}

func (OrganizationMockErrContext) ID(*int32) error {
	return errors.New("Invalid ID")
}

func (OrganizationMockErrContext) PaginationQueryParam(*model.PaginationQueryParams) error {
	return errors.New("Invalid Query Param")
}

func InitializeMockOrganization() {
	Organization1 = proto.Organization{
		Id:          1,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization2 = proto.Organization{
		Id:          2,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization3 = proto.Organization{
		Id:          3,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organization4 = proto.Organization{
		Id:          4,
		Name:        faker.Word(),
		Description: faker.Sentence(),
	}

	Organizations = append(Organizations, &Organization1, &Organization2, &Organization3, &Organization4)
}
