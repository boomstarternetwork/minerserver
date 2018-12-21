package dbstore

import (
	"testing"

	"github.com/boomstarternetwork/bestore"
	"github.com/stretchr/testify/assert"
)

func Test_DBStore_ProjectsList_success(t *testing.T) {
	createTestingTables()
	defer dropTestingTables()

	err := s.gdb.Model(&bestore.Project{}).
		Create(&bestore.Project{ID: "id-2", Name: "name-2"}).
		Create(&bestore.Project{ID: "id-1", Name: "name-1"}).Error
	if !assert.NoError(t, err) {
		return
	}

	wantProjects := []bestore.Project{
		{ID: "id-1", Name: "name-1"},
		{ID: "id-2", Name: "name-2"},
	}

	projects, err := s.GetProjects()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, wantProjects, projects)
}
