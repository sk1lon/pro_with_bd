package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (conn *Connection) CreateTask(w http.ResponseWriter, r *http.Request) {
	var jsForm Tasks
	err := json.NewDecoder(r.Body).Decode(&jsForm)
	if err != nil {
		fmt.Println("Ошибка на декодировании")
	}
	sqlQuery := `
		INSERT INTO newtasks (title,description,completed,created_at)
		VALUES ($1,$2,$3,$4)
	`
	_, err = conn.Conn.Exec(r.Context(), sqlQuery, jsForm.Title, jsForm.Description, jsForm.Completed, jsForm.Created_at)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Все данные отправлены в бд!")
	w.Write([]byte("Все данные отправлены в бд!"))
}
