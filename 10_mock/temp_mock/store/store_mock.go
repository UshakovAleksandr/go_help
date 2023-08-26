package store

type storeMock struct{}

func NewStoreMock() Store {
	return &storeMock{}
}

func (s storeMock) GetUser(id int) string {
	return "Ushakov"
}
