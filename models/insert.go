package models

import (
	"github.com/swxtz/api-go/db"
)

func Insert(todo Todo) (id int64, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`
	err = conn.QueryRow(query, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
