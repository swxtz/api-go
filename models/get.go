package models

import "github.com/swxtz/api-go/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT * FROM todos WHERE id = $1`
	row := conn.QueryRow(query, id)

	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
