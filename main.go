package main

import (
	"context"
	"db/tasks"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:skj09@localhost:5432/postgres")
	if err != nil {
		fmt.Println("ERROR")
	}
	db := &tasks.Connection{
		Conn: conn,
	}
	tasks.CreateTable(conn, ctx)
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(db.CreateTask)
	router.Path("/get").Methods("GET").HandlerFunc(db.GetTasks)
	if errors := http.ListenAndServe(":9091", router); errors != nil {
		fmt.Println("BAAAAd")
	}
}
