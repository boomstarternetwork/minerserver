package dbstore

import (
	"testing"

	"github.com/boomstarternetwork/minerserver/store"
	"github.com/stretchr/testify/assert"
)

func Test_DBStore_ProjectsList_success(t *testing.T) {
	createTestingTables()
	defer dropTestingTables()

	err := s.gdb.Model(&store.Project{}).
		Create(&store.Project{ID: "id-2", Name: "name-2"}).
		Create(&store.Project{ID: "id-1", Name: "name-1"}).Error
	if !assert.NoError(t, err) {
		return
	}

	wantProjects := []store.Project{
		{ID: "id-1", Name: "name-1"},
		{ID: "id-2", Name: "name-2"},
	}

	projects, err := s.ListProjects()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, wantProjects, projects)
}
