package main

import (
	"github.com/boomstarternetwork/bestore"
	"github.com/stretchr/testify/mock"
)

type mockStore struct {
	mock.Mock
}

func newMockStore() *mockStore {
	return &mockStore{}
}

func (ps *mockStore) GetProjects() ([]bestore.Project, error) {
	args := ps.Called()
	projects := args.Get(0)
	if projects == nil {
		return nil, args.Error(1)
	}
	return projects.([]bestore.Project), args.Error(1)
}
