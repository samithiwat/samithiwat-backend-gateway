package common

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UtilTest struct {
	suite.Suite
}

func TestUtil(t *testing.T) {
	suite.Run(t, new(UtilTest))
}

func (u *UtilTest) SetupTest() {
}

func (u *UtilTest) TestIsExistedTrue() {
	e := map[string]struct{}{
		"A": {},
		"B": {},
	}

	ok := IsExisted(e, "A")

	assert.True(u.T(), ok)
}

func (u *UtilTest) TestIsExistedFalse() {
	e := map[string]struct{}{
		"A": {},
		"B": {},
	}

	ok := IsExisted(e, "C")

	assert.False(u.T(), ok)
}

func (u *UtilTest) TestFormatPathWithID() {
	want := "POST /user/:id"

	path := FormatPath("POST", "/user/1", 1)

	assert.Equal(u.T(), want, path)
}

func (u *UtilTest) TestFormatPathWithoutID() {
	want := "POST /user"

	path := FormatPath("POST", "/user", 0)

	assert.Equal(u.T(), want, path)
}

func (u *UtilTest) TestGetIntFromStrFound() {
	want := []int32{1}

	id := FindIntFromStr("/user/1")

	assert.Equal(u.T(), want, id)
}

func (u *UtilTest) TestGetIntFromStrNotFound() {
	var want []int32

	id := FindIntFromStr("/user")

	assert.Equal(u.T(), want, id)
}
