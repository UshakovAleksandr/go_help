package handler

import (
	"mocking/store"
)

type UseCase struct {
	Store store.Store
}

func NewHandler(s store.Store) *UseCase {
	return &UseCase{Store: s}
}

func (u *UseCase) GetUserHandler(id int) string {
	return u.Store.GetUser(id)
}

func (u *UseCase) CreateUserHandler(name string) string {
	return u.Store.CreateUser(name)
}
