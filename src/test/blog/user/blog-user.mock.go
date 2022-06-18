package blog_user

import (
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"github.com/stretchr/testify/mock"
)

type ContextMock struct {
	mock.Mock
	V           interface{}
	BlogUser    *proto.BlogUser
	BlogUserDto *dto.BlogUserDto
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	*v.(*dto.BlogUserDto) = *c.BlogUserDto

	return args.Error(0)
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) ID() (int32, error) {
	args := c.Called()

	return int32(args.Int(0)), args.Error(1)
}

func (c *ContextMock) UserID() int32 {
	args := c.Called()

	return int32(args.Int(0))
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindOne(id int32) (res *proto.BlogUser, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.BlogUser)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Create(in *dto.BlogUserDto) (res *proto.BlogUser, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.BlogUser)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Update(id int32, in *dto.BlogUserDto) (res *proto.BlogUser, err *dto.ResponseErr) {
	args := s.Called(id, in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.BlogUser)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Delete(id int32) (res *proto.BlogUser, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.BlogUser)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) AddBookmark(id int32, userId int32) (res []*uint32, err *dto.ResponseErr) {
	args := s.Called(id, userId)

	if args.Get(0) != nil {
		res = args.Get(0).([]*uint32)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) FindAllBookmark(id int32) (res []*uint32, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		res = args.Get(0).([]*uint32)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) DeleteBookmark(id int32, userId int32) (res bool, err *dto.ResponseErr) {
	args := s.Called(id, userId)

	if args.Get(0) != nil {
		res = args.Get(0).(bool)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Read(id int32, userId int32) (res bool, err *dto.ResponseErr) {
	args := s.Called(id, userId)

	if args.Get(0) != nil {
		res = args.Get(0).(bool)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}
