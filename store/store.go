package store

import "github.com/boomstarternetwork/bestore"

type Store interface {
	GetProjects() ([]bestore.Project, error)
}
