package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DbConnectionPool *pgxpool.Pool

func NewDbConnectionPool() {
	connectionUrl := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), connectionUrl)
	if err != nil {
		panic(err)
	}
	DbConnectionPool = pool
	fmt.Println("Connection Pool Created!!!")
}
