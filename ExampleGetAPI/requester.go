package main

import "github.com/parnurzeal/gorequest"

// IRequester is interface for requester
type IRequester interface {
	Get(url string) (string, error)
}

// Get make get request
func (*Requester) Get(url string) (string, error) {
	_, body, errs := gorequest.New().Get(url).End()
	if len(errs) > 0 {
		return "", errs[0]
	}
	return body, nil
}

// Requester is struct for requester
type Requester struct{}

// NewRequester return new requester
func NewRequester() *Requester {
	return &Requester{}
}
