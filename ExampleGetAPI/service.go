package main

import "encoding/json"

// MemberService is service for member
type MemberService struct{}

// NewMemberService return new MemberService
func NewMemberService() *MemberService {
	return &MemberService{}
}

func (*MemberService) getMembers(requester IRequester) ([]*Actors, error) {
	body, err := requester.Get("http://member-service/api/people")
	if err != nil {
		return nil, err
	}

	actors := make([]*Member, 0)
	err = json.Unmarshal([]byte(body), &actors)
	if err != nil {
		return nil, err
	}

	return actors, nil
}
