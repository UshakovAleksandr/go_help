package main

import (
	"fmt"
	"log"

	"44_pgxpool/storage"
	"44_pgxpool/usecase"
)

func execute(usecase *usecase.Usecases) {
	// insert
	usecase.AddUsers()
	usecase.AddNotes()

	// get one
	usecase.GetUser()
	usecase.GetNote()

	////get all
	usecase.GetAllUsers()
	usecase.GetAllNotesByUser()

	//update
	usecase.UpdateUser()
	usecase.UpdateNote()

	// delete
	usecase.DeleteUser()
	usecase.DeleteNote()
}

func main() {
	// conn to DB create
	db, err := storage.NewDB()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to init db: %s", err.Error()))
	}

	store := storage.NewStorages(db)
	use := usecase.New(store)

	execute(use)
}

//package main
//
//// package storage
//
//import (
//	"context"
//
//	"github.com/jackc/pgx/v4/pgxpool"
//)
//
//// Storage - Хранилище данных.
//type Storage struct {
//	db *pgxpool.Pool
//}
//
//// New - Конструктор, принимает строку подключения к БД.
//func New(constr string) (*Storage, error) {
//	db, err := pgxpool.Connect(context.Background(), constr)
//	if err != nil {
//		return nil, err
//	}
//	s := Storage{
//		db: db,
//	}
//	return &s, nil
//}
//
//// Task - Задача.
//type Task struct {
//	ID         int
//	Opened     int64
//	Closed     int64
//	AuthorID   int
//	AssignedID int
//	Title      string
//	Content    string
//}
//
//// Tasks - возвращает список задач из БД.
//func (s *Storage) Tasks(taskID, authorID int) ([]Task, error) {
//	rows, err := s.db.Query(context.Background(), `
//		SELECT
//			id,
//			opened,
//			closed,
//			author_id,
//			assigned_id,
//			title,
//			content
//		FROM tasks
//		WHERE
//			($1 = 0 OR id = $1) AND
//			($2 = 0 OR author_id = $2)
//		ORDER BY id;
//	`,
//		taskID,
//		authorID,
//	)
//	if err != nil {
//		return nil, err
//	}
//	var tasks []Task
//	// итерирование по результату выполнения запроса
//	// и сканирование каждой строки в переменную
//	for rows.Next() {
//		var t Task
//		err = rows.Scan(
//			&t.ID,
//			&t.Opened,
//			&t.Closed,
//			&t.AuthorID,
//			&t.AssignedID,
//			&t.Title,
//			&t.Content,
//		)
//		if err != nil {
//			return nil, err
//		}
//		// добавление переменной в массив результатов
//		tasks = append(tasks, t)
//
//	}
//	// ВАЖНО не забыть проверить rows.Err()
//	return tasks, rows.Err()
//}
//
//// NewTask создаёт новую задачу и возвращает её id.
//func (s *Storage) NewTask(t Task) (int, error) {
//	var id int
//	err := s.db.QueryRow(context.Background(), `
//		INSERT INTO tasks (title, content)
//		VALUES ($1, $2) RETURNING id;
//		`,
//		t.Title,
//		t.Content,
//	).Scan(&id)
//	return id, err
//}

//////////////////////////////////////////

// Query(). Используется для выборки данных, то есть для выполнения оператора SELECT.
// rows, err := db.Query(`SELECT * FROM users ORDER BY id;`)

// QueryRow(). Аналогичен Query, но используется, когда запрос возвращает ровно одну строку.
// err := db.QueryRow(`SELECT name FROM users WHERE id = 10;`).Scan(u.name)

// Exec(). Используется для всех запросов, которые могут изменять данные.
// res, err := db.Exec(`INSERT INTO users (name) VALUES ('Rob'), ('Ken');`)

/*
   Схема БД для информационной системы
   отслеживания выполнения задач.
*/

//DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;
//
//-- пользователи системы
//CREATE TABLE users (
//id SERIAL PRIMARY KEY,
//name TEXT NOT NULL
//);
//
//-- метки задач
//CREATE TABLE labels (
//id SERIAL PRIMARY KEY,
//name TEXT NOT NULL
//);
//
//-- задачи
//CREATE TABLE tasks (
//id SERIAL PRIMARY KEY,
//opened BIGINT NOT NULL DEFAULT extract(epoch from now()), -- время создания задачи
//closed BIGINT DEFAULT 0 0, -- время выполнения задачи
//author_id INTEGER REFERENCES users(id) DEFAULT 0, - автор задачи
//assigned_id INTEGER REFERENCES users(id) DEFAULT 0, -- ответственный
//title TEXT, -- название задачи
//content TEXT -- задачи
//);
//
//-- связь многие - ко- многим между задачами и метками
//CREATE TABLE tasks_labels (
//task_id INTEGER REFERENCES tasks(id),
//label_id INTEGER REFERENCES labels(id)
//);
//-- наполнение БД начальными данными
//INSERT INTO users (id, name) VALUES (0, 'default');

/////////
//// addBooks добавляет в БД массив книг одной транзакцией.
//func addBooks(ctx context.Context, db *pgxpool.Pool, books []book) error {
//	// начало транзакции
//	tx, err := db.Begin(ctx)
//	if err != nil {
//		return err
//	}
//	// отмена транзакции в случае ошибки
//	defer tx.Rollback(ctx)
//
//	// пакетный запрос
//	batch := new(pgx.Batch)
//	// добавление заданий в пакет
//	for _, book := range books {
//		batch.Queue(`INSERT INTO books(title, year) VALUES ($1, $2)`, book.Title, book.Year)
//	}
//	// отправка пакета в БД (может выполняться для транзакции или соединения)
//	res := tx.SendBatch(ctx, batch)
//	// обязательная операция закрытия соединения
//	err = res.Close()
//	if err != nil {
//		return err
//	}
//	// подтверждение транзакции
//	return tx.Commit(ctx)
//}

//// addUsersTx добавляет пользователей в БД.
//// Используется транзакция.
//func addUsersTx(db *sql.DB, users []user) error {
//	tx, err := db.Begin()
//	if err != nil {
//		return err
//	}
//	// tx - объект транзакции; позволяет управлять ее работой
//	for _, u := range users {
//		// запрос на вставку данных
//		_, err := tx.Exec(`
//		INSERT INTO users (name)
//		VALUES (?);
//		`,
//			u.name,
//		)
//		if err != nil {
//			// откат транзакции в случае ошибки
//			tx.Rollback()
//			return err
//		}
//	}
//	// фиксация (подтверждение) транзакции
//	tx.Commit()
//	return nil
//}

//// addUsersPrepared добавляет пользователей в БД.
//// Используется подготовленный запрос.
//func addUsersPrepared(db *sql.DB, users []user) error {
//	// подготовка запроса
//	stmt, err := db.Prepare(`INSERT INTO users (name)VALUES (?);`)
//	if err != nil {
//		return err
//	}
//	for _, u := range users {
//		// выполнение подготовленного запроса на вставку данных
//		_, err = stmt.Exec(u.name)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

// db, err := sql.Open("sqlite3", "books.db")
//// ...
//
//query := `
//    drop table if exists books;
//    create table if not exists books(
//        id integer primary key,
//        title text,
//        author text,
//        num_pages integer,
//        rating real
//    );
//`
//
//_, err := db.Exec(query)
//if err != nil {
//    panic(err)
//}
//fmt.Println("✓ created books table")
//// ✓ created books table

//// //////////////// Подготовленные выражения///////////
//func prep() {
//	stmt, err := db.Prepare(`
//    insert into books(title, author, num_pages, rating)
//    values (?, ?, ?, ?)
//`)
//	if err != nil {
//		panic(err)
//	}
//	defer stmt.Close()
//
//	data := [][]any{
//		{"The Catcher in the Rye", "J.D. Salinger", 277, 3.8},
//		{"The Fellowship of the Ring", "J.R.R. Tolkin", 398, 4.36},
//		{"The Giver", "Lois Lowry", 208, 4.13},
//	}
//
//	for _, vals := range data {
//		res, err := stmt.Exec(vals...)
//		if err != nil {
//			panic(err)
//		}
//		bookID, _ := res.LastInsertId()
//		fmt.Printf("added new book: id=%d\n", bookID)
//	}
//
//	/*
//	   added new book: id=1
//	   added new book: id=2
//	   added new book: id=3
//	*/
//
//	//Exec(args ...any) (Result, error)
//	//Query(args ...any) (*Rows, error)
//	//QueryRow(args ...any) *Row
//}

/////////////////////////////////Транзакции ////////////////
//// ┌────┬───────┬─────────┐
////│ id │ name  │ balance │
////├────┼───────┼─────────┤
////│ 1  │ Alice │ 100     │
////│ 2  │ Bob   │ 100     │
////└────┴───────┴─────────┘
//
//// ┌────┬───────┬─────────┐
//// │ id │ name  │ balance │
//// ├────┼───────┼─────────┤
//// │ 1  │ Alice │ 50      │
//// │ 2  │ Bob   │ 150     │
//// └────┴───────┴─────────┘
//
//// begin transaction;
//// update accounts set balance = balance - 50 where id = 1;
//// update accounts set balance = balance + 50 where id = 2;
//// commit transaction;
//func foo() {
//	db, err := sql.Open("sqlite3", "accounts.db")
//	// ...
//	const (
//		aliceID = 1
//		bobID   = 2
//		amount  = 50
//	)
//	err = transfer(db, aliceID, bobID, amount)
//	if err != nil {
//		panic(err)
//	}
//}
//
//// transfer переносит amount денег с счета fromID на счет toID
//func transfer(db *sql.DB, fromID, toID, amount int) error {
//	const query = "update accounts set balance = balance + ? where id = ?"
//
//	tx, err := db.Begin() // (1)
//	if err != nil {
//		return err
//	}
//	defer tx.Rollback() // (2)
//
//	_, err = tx.Exec(query, -amount, fromID) // (3)
//	if err != nil {
//		return err
//	}
//	_, err = tx.Exec(query, amount, toID) // (4)
//	if err != nil {
//		return err
//	}
//	return tx.Commit() // (5)
//
//	// Exec(query string, args ...any) (Result, error)
//	//Query(query string, args ...any) (*Rows, error)
//	//QueryRow(query string, args ...any) *Row
//	//Prepare(query string) (*Stmt, error)
//}

// Благодаря Tx.Prepare() можно подготовить выражение sql.Stmt
//и использовать его внутри транзакции, аналогично тому, как мы это делали
//с sql.DB на предыдущем шаге. Кроме того,
//Tx умеет взять существующее подготовленное выражение и «перенести»
//его внутрь транзакции:

// // вне транзакции
//stmt, err := db.Prepare(query)
//// ...
//
//// внутри транзакции
//txStmt := tx.Stmt(stmt)
//// запрос выполнится внутри транзакции
//txStmt.Exec(vals...)

/////////////////////////////// Отмена операций //////////

//func ex() {
//	db, err := sql.Open("sqlite3", "books.db")
//	// ...
//
//	// таймаут 3 секунды на выполнение запроса
//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
//	defer cancel()
//
//	query := "select id, title, author, num_pages, rating from books"
//	rows, err := db.QueryContext(ctx, query)
//	if err != nil {
//		panic(err)
//	}
//	defer rows.Close()
//
//	// ...
//}

// Для sql.DB:
//// контекст используется для выполнения запроса
//ExecContext(ctx context.Context, query string, args ...any) (Result, error)
//QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
//QueryRowContext(ctx context.Context, query string, args ...any) *Row
//
//// контекст используется только для подготовки выражения, не для выполнения
//PrepareContext(ctx context.Context, query string) (*Stmt, error)
//
//// контекст используется для транзакции в целом, до фиксации или отката
//BeginTx(ctx context.Context, opts *TxOptions) (*Tx, error)

// Для sql.Tx:
//
//// контекст используется для выполнения запроса
//ExecContext(ctx context.Context, query string, args ...any) (Result, error)
//QueryContext(ctx context.Context, query string, args ...any) (*Rows, error)
//QueryRowContext(ctx context.Context, query string, args ...any) *Row
//
//// контекст используется только для подготовки выражения, не для выполнения
//PrepareContext(ctx context.Context, query string) (*Stmt, error)
//StmtContext(ctx context.Context, stmt *Stmt) *Stmt

// Для sql.Stmt:
//
//ExecContext(ctx context.Context, args ...any) (Result, error)
//QueryContext(ctx context.Context, args ...any) (*Rows, error)
//QueryRowContext(ctx context.Context, args ...any) *Row

//Метод DB.BeginTx() позволяет задать не только контекст, но и уровень изоляции транзакции:
//
//ctx := context.Background()
//opts := sql.TxOptions{Isolation: sql.LevelSerializable}
//tx, err := db.BeginTx(ctx, &opts)
