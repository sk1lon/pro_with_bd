package tasks

import "github.com/jackc/pgx/v5"

type Connection struct {
	Conn *pgx.Conn
}
