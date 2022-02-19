package main

import (
	"context"
	"database/sql"
	"os"
	"strconv"
	"time"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

//Context timeout
var Timeout = 3 * time.Second

//DB global variable
var db *sql.DB

func DBConnect() error {
	var err error

	conn := os.Getenv("DATABASE_URL")

	db, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}
	return db.Ping()
}

type User struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
}

type Comments struct {
	ID      int    `json:"id"`
	Comment string `json:"comment"`
	Date    int    `json:"date"`
	UserID  int    `json:"userid"`
}

func getUsersDB(ctx context.Context) ([]User, errAPI) {
	queryCtx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	rows, err := db.QueryContext(queryCtx, "SELECT * FROM users")
	if err == context.DeadlineExceeded {
		return nil, tooSlowErr
	}
	if err != nil {
		ErrLog.Println(err)
		return nil, unexpErr
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Age, &u.Gender); err != nil {
			ErrLog.Println(err)
			return nil, unexpErr
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		ErrLog.Println(err)
		return nil, unexpErr
	}

	return users, nil
}

func getCommentsDB(ctx context.Context, uid string) ([]Comments, errAPI) {
	queryCtx, cancel := context.WithTimeout(ctx, Timeout)
	defer cancel()

	id, err := strconv.Atoi(uid)
	if err != nil || id < 1 {
		return nil, noUserIDErr
	}

	rows, err := db.QueryContext(queryCtx, "SELECT * FROM comments WHERE userid = $1", uid)
	if err == context.DeadlineExceeded {
		return nil, tooSlowErr
	}
	if err != nil {
		ErrLog.Println(err)
		return nil, unexpErr
	}
	defer rows.Close()

	var comms []Comments

	for rows.Next() {
		var c Comments
		err := rows.Scan(&c.ID, &c.Comment, &c.Date, &c.UserID)
		if err != nil {
			ErrLog.Println(err)
			return nil, unexpErr
		}
		comms = append(comms, c)
	}

	if err := rows.Err(); err != nil {
		ErrLog.Println(err)
		return nil, unexpErr
	}

	if len(comms) == 0 {
		return nil, noDataInDBErr
	}

	return comms, nil
}
