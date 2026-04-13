package tasks

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateTable(conn *pgx.Conn, ctx context.Context) {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS newtasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(200),
			description VARCHAR(1000),
			completed BOOLEAN,
			created_at TIMESTAMP
		);`

	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		fmt.Println("GG")
	}
}
