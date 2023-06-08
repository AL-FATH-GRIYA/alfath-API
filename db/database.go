package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dbInstance *bun.DB

func InitDB() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/al-fath-dev?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	dbInstance = bun.NewDB(sqlDb, pgdialect.New())

	return dbInstance
}

// GetConn return database connection instance
func GetConn() *bun.DB {
	return dbInstance
}
