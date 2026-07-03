package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// InitDB устанавливает соединение с БД и возвращает *sql.DB.
// Вызывающий обязан закрыть его через defer.
func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected successfully")
	return db, nil
}
