package store

type Store interface {
	GetUser(id int) string
	CreateUser(name string) string
}

type store struct{}

func NewStore() Store {
	return &store{}
}

func (s *store) GetUser(id int) string {
	if id == 1 {
		return "Alex"
	}
	return ""
}

func (s *store) CreateUser(name string) string {
	if name == "1" {
		return "Ushakov"
	}
	return ""
}
