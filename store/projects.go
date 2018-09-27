package store

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProjectsStore interface {
	List() ([]Project, error)
}

type dbProjectsStore struct {
	db *sql.DB
}

func NewDBProjectsStore(db *sql.DB) ProjectsStore {
	return dbProjectsStore{db: db}
}

func (ps dbProjectsStore) List() ([]Project, error) {
	var projects []Project

	rows, err := ps.db.Query(
		"SELECT id, name FROM projects ORDER BY name ASC")
	if err != nil {
		return projects, err
	}
	defer rows.Close()

	for rows.Next() {
		var project Project

		err = rows.Scan(&project.ID, &project.Name)
		if err != nil {
			return projects, err
		}

		projects = append(projects, project)
	}

	return projects, rows.Err()
}

type MockProjectsStore struct {
	mock.Mock
}

func NewMockProjectsStore() *MockProjectsStore {
	return &MockProjectsStore{}
}

func (ps *MockProjectsStore) List() ([]Project, error) {
	args := ps.Called()
	projects := args.Get(0)
	if projects == nil {
		return nil, args.Error(1)
	}
	return projects.([]Project), args.Error(1)
}
