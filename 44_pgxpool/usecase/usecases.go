package usecase

import (
	"44_pgxpool/storage"
)

type Usecases struct {
	usecases *storage.Storages
}

func New(store *storage.Storages) *Usecases {
	return &Usecases{usecases: store}
}
