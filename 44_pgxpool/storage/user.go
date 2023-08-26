package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"44_pgxpool/model"
)

type userStorage struct {
	db *pgxpool.Pool
}

func NewUserStorage(db *pgxpool.Pool) UserStorage {
	return &userStorage{db: db}
}

func (u *userStorage) Add(user *model.User) (int, error) {
	var (
		query = `
			INSERT INTO users (username, password)
			VALUES ($1, $2)
			RETURNING id
		`
		id int
	)

	err := u.db.QueryRow(context.Background(), query, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("db err: %w", err)
	}

	return id, nil
}

func (u *userStorage) GetById(userID int) (model.User, error) {
	var (
		query = `
			SELECT id, username, password
			FROM users
			WHERE id=$1
		`
		user model.User
	)

	row := u.db.QueryRow(context.Background(), query, userID)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("user not found, err: %w", err)
		}
		return user, fmt.Errorf("db err: %w", err)
	}

	return user, nil
}

func (u *userStorage) GetAll() ([]model.UserResp, error) {
	var (
		query = `
			SELECT id, username
			FROM users
		`
		users []model.UserResp
	)

	rows, err := u.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("db err: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user model.UserResp

		if err = rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	// Если запрос в целом не отработал, то первый же вызов Rows.Next() вернет false, так что внутрь цикла мы не попадем.
	// Отловить такие ошибки поможет следующая проверка на Rows.Err()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("db err: %w", err)
	}

	return users, nil
}

func (u *userStorage) Update(user model.User) error {
	var (
		query = `
			UPDATE users
			SET username=$1, password=$2
			WHERE id=$3
		`
	)

	_, err := u.db.Exec(context.Background(), query, user.Username, user.Password, user.ID)
	if err != nil {
		return fmt.Errorf("db err: %w", err)
	}

	return nil
}

func (u *userStorage) DeleteById(userID int) error {
	query := `
		DELETE FROM users
		WHERE id=$1
	`

	res, err := u.db.Exec(context.Background(), query, userID)
	if err != nil {
		return fmt.Errorf("db err: %w", err)
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
