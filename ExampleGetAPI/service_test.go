package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type MemberServiceTestSuite struct {
	suite.Suite
}

func TestMemberServiceTypeSuite(t *testing.T) {
	suite.Run(t, &MemberServiceTestSuite{})
}

func (ts *MemberServiceTestSuite) TestGetMembers_ExpectCorrectMembersReturn() {
	mockRequester := NewMockRequester()
	// Usage testify libray
	mockRequester.On("Get", "https://swapi.co/api/people/1").Return(
		`"name": "Luke Skywalker",
		"height": "172"`, nil)
	mockCacher.On("Set", "member_cache", "{}").Return(true)
	
	memberSvc := NewMemberService()
	members, err := memberSvc.getMembers(mockRequester)

	is := assert.New(ts.T())
	if is.NoError(err) {
		is.Equal("Luke Skywalker", members.Name)
		is.Equal("172", members.Height)
	}
	mockRequester.AssertExpectations(ts.T())
}

func (ts *MemberServiceTestSuite) TestGetMembers_GivenError_ExpectErrorReturn() {
	mockRequester := NewMockRequester()
	// Usage testify libray
	mockRequester.On("Get", "http://member-service/api/members").
		Return(``, fmt.Errorf("Error 1"))

	memberSvc := NewMemberService()
	members, err := memberSvc.getMembers(mockRequester)

	is := assert.New(ts.T())
	is.Nil(members)
	is.Error(err)
	is.Equal("Error 1", err.Error())

	mockRequester.AssertExpectations(ts.T())
}
