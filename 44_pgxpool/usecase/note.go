package usecase

import (
	"fmt"

	ruauka "github.com/ruauka/attrs-go"

	"44_pgxpool/model"
)

func (u *Usecases) AddNotes() {
	notes := []model.Note{
		{
			Title:  "title1",
			Info:   "info1",
			UserID: 1,
		},
		{
			Title:  "title2",
			Info:   "info2",
			UserID: 2,
		},
	}
	for _, note := range notes {
		id, err := u.usecases.Note.Add(&note)
		if err != nil {
			fmt.Println("ERROR by adding new note: ", err)
			return
		}
		fmt.Printf("added new note: id=%d\n", id)
	}
}

func (u *Usecases) GetNote() {
	note, err := u.usecases.Note.GetById(1, 1)
	if err != nil {
		fmt.Println("func - GetNoteById. ERROR by getting note:", err)
		return
	}

	fmt.Printf("get note: %#v\n", note)
}

func (u *Usecases) GetAllNotesByUser() {
	notes, err := u.usecases.Note.GetAll(1)
	if err != nil {
		fmt.Println("func - GetAllNotesByUser. ERROR by getting notes:", err)
		return
	}

	fmt.Printf("get notes: %#v\n", notes)
}

func (u *Usecases) UpdateNote() {
	var (
		userID    = 1
		noteID    = 1
		noteTitle = "new_title1"
		noteInfo  = "new_info1"
	)

	newNote := model.NoteUpdate{
		Title: &noteTitle,
		Info:  &noteInfo,
	}

	if IsEmpty(newNote, model.NoteUpdate{}) {
		fmt.Println("no fields to update")
		return
	}

	note, err := u.usecases.Note.GetById(noteID, userID)
	if err != nil {
		fmt.Println("func - UpdateNote. ERROR by updating note:", err)
	}

	if err := ruauka.SetStructAttrs(&note, newNote); err != nil {
		fmt.Println("func - UpdateNote. ERROR by updating note:", err)
		return
	}

	if err := u.usecases.Note.Update(note); err != nil {
		fmt.Println("func - UpdateUser. ERROR by updating user:", err)
		return
	}

	fmt.Printf("update note: %#v\n", note)
}

func (u *Usecases) DeleteNote() {
	var (
		userID = 2
		noteID = 2
	)

	err := u.usecases.Note.DeleteById(userID, noteID)
	if err != nil {
		fmt.Println("func - DeleteNote. ERROR by deleting note:", err)
		return
	}

	fmt.Printf("deleted note: id=%d\n", userID)
}
