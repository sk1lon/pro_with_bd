package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (conn *Connection) GetTasks(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
		SELECT title, description, completed, created_at
		FROM newtasks
	`
	rows, err := conn.Conn.Query(r.Context(), sqlQuery)
	if err != nil {
		fmt.Println("QUERY211")
	}
	defer rows.Close()
	for rows.Next() {
		var task Tasks
		err := rows.Scan(&task.Title, &task.Description, &task.Completed, &task.Created_at)
		if err != nil {
			fmt.Println("GOIDA")
		}
		fmt.Println(task.Title, task.Description, task.Completed, task.Created_at)
		json.NewEncoder(w).Encode(task)
	}
}
