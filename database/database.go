package database

import (
	"context"
	"fmt"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DbConnectionPool *pgxpool.Pool
var Pgsq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func NewDbConnectionPool() {
	connectionUrl := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), connectionUrl)
	if err != nil {
		panic(err)
	}
	DbConnectionPool = pool
	fmt.Println("Connection Pool Created!!!")
}
