package store

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Store interface {
	ListProjects() ([]Project, error)
}
