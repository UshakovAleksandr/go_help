package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"44_pgxpool/model"
)

type noteStorage struct {
	db *pgxpool.Pool
}

func NewNoteStorage(db *pgxpool.Pool) NoteStorage {
	return &noteStorage{db: db}
}

func (n *noteStorage) Add(note *model.Note) (int, error) {
	var (
		id    int
		query = `
			INSERT INTO notes (title, info, user_id)
			VALUES ($1, $2, $3)
			RETURNING id
		`
	)

	err := n.db.QueryRow(context.Background(), query, note.Title, note.Info, note.UserID).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("db err: %w", err)
	}

	return id, nil
}

func (n *noteStorage) GetById(noteID, userID int) (model.Note, error) {
	var (
		query = `
			SELECT id, title, info, user_id
			FROM notes
			WHERE id=$1
			AND user_id=$2
		`
		note model.Note
	)

	row := n.db.QueryRow(context.Background(), query, noteID, userID)
	if err := row.Scan(&note.ID, &note.Title, &note.Info, &note.UserID); err != nil {
		if err == pgx.ErrNoRows {
			return note, fmt.Errorf("note not found, err: %w", err)
		}
		return note, fmt.Errorf("db err: %w", err)
	}

	return note, nil
}

func (n *noteStorage) GetAll(userID int) ([]model.Note, error) {
	var (
		query = `
			SELECT id, title, info, user_id
			FROM notes
			WHERE user_id=$1
		`
		notes []model.Note
	)

	rows, err := n.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("db err: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var note model.Note

		if err = rows.Scan(&note.ID, &note.Title, &note.Info, &note.UserID); err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}
	// Если запрос в целом не отработал, то первый же вызов Rows.Next() вернет false, так что внутрь цикла мы не попадем.
	// Отловить такие ошибки поможет следующая проверка на Rows.Err()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db err: %w", err)
	}

	return notes, nil
}

func (n *noteStorage) Update(note model.Note) error {
	var (
		query = `
			UPDATE notes
			SET title=$1, info=$2
			WHERE id=$3
		`
	)

	_, err := n.db.Exec(context.Background(), query, note.Title, note.Info, note.ID)
	if err != nil {
		return fmt.Errorf("db err: %w", err)
	}

	return nil
}

func (n *noteStorage) DeleteById(noteID, userID int) error {
	query := `
		DELETE FROM notes
		WHERE id=$1
		AND user_id=$2
	`

	res, err := n.db.Exec(context.Background(), query, noteID, userID)
	if err != nil {
		return fmt.Errorf("db err: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("note not found")
	}

	return nil
}
