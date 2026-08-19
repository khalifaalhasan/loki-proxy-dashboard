package config

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB{
	dbPath := GetEnv("DB_PATH", "./dashboard.db")

	db, err := sql.Open("sqliite3", dbPath)
	if err != nil {
		log.Fatal("Gaga buka database :", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Database ga respon: ", err)
	}

	return db
}