// Package model Package model
package model

// Note model.
type Note struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Info   string `json:"info"`
	UserID int    `json:"user_id"`
}

// NoteUpdate model.
type NoteUpdate struct {
	Title *string `json:"title"`
	Info  *string `json:"info"`
}
