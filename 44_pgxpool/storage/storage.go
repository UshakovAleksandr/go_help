package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"44_pgxpool/model"
)

// UserStorage - interface.
type UserStorage interface {
	Add(user *model.User) (int, error)
	GetById(userID int) (model.User, error)
	GetAll() ([]model.UserResp, error)
	Update(user model.User) error
	DeleteById(userID int) error
}

// NoteStorage - interface.
type NoteStorage interface {
	Add(note *model.Note) (int, error)
	GetById(noteID int, userID int) (model.Note, error)
	GetAll(userID int) ([]model.Note, error)
	Update(note model.Note) error
	DeleteById(noteID, userID int) error
}

type Storages struct {
	User UserStorage
	Note NoteStorage
}

func NewStorages(db *pgxpool.Pool) *Storages {
	return &Storages{
		User: NewUserStorage(db),
		Note: NewNoteStorage(db),
	}
}
