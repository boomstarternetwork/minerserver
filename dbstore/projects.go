package dbstore

import (
	"github.com/boomstarternetwork/bestore"
)

func (s *DBStore) GetProjects() (ps []bestore.Project, err error) {
	err = s.gdb.Order("name").Find(&ps).Error
	return
}
