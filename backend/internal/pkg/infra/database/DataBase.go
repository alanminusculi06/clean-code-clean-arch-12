package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type DataBase struct {
}

func NewDataBase() DataBase {
	return DataBase{}
}

func (database DataBase) Config() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mysqlroot",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "ccca_t12",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}
