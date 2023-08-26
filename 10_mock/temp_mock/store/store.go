package store

type Store interface {
	GetUser(id int) string
}

type store struct{}

func NewStore() Store {
	return &store{}
}

func (s store) GetUser(id int) string {
	if id == 1 {
		return "Ivanova"
	}
	return "Ushakova"
}
