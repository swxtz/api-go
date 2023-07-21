package models

import "github.com/swxtz/api-go/db"

func Update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := `UPDATE todos SET title=$2, description=$3, done=$4 WHERE id=$1`
	res, err := conn.Exec(query, id, todo.Title, todo.Description, todo.Done)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
