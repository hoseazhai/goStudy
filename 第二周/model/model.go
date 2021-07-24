package model

import (
	"database/sql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MyDB struct {
	DB *sql.DB
}

func NewDBEngine() (*sql.DB, error) {
	//db := MyDB{}
	db, err := sql.Open("mysql", "root:bt5@tcp(127.0.0.1:3306)/instant")
	if err != nil {
		return nil, err
	}
	//db.DB = d
	return db, nil
}
