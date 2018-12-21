package dbstore

import (
	"github.com/boomstarternetwork/minerserver/store"
)

func (s *DBStore) ListProjects() (ps []store.Project, err error) {
	err = s.gdb.Order("name").Find(&ps).Error
	return
}
