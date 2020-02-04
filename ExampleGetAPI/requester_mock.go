package main

import "github.com/stretchr/testify/mock"

// MockRequester is mock for requester
type MockRequester struct {
	mock.Mock
}

// NewMockRequester return new mock requester
func NewMockRequester() *MockRequester {
	return &MockRequester{}
}

// Get is mock function
func (r *MockRequester) Get(url string) (string, error) {
	args := r.Called(url)
	return args.String(0), args.Error(1)
}
