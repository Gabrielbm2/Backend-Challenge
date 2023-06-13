package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"

	_ "github.com/lib/pq"
)

var db *pgx.Conn

func GetDatabase() *pgx.Conn {
	if db == nil {
		Connect()
	}
	return db
}

func Connect() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db = conn
}
