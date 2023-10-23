package app

import (
	"database/sql"
	"github.com/jajangratis/money-manager-clone-api/helper"
	_ "github.com/lib/pq"
	"os"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://"+os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_SELECT")+"?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
