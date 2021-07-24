package model

import (
	"database/sql"
)

var DB, _ = sql.Open("mysql", "root:bt5@tcp(127.0.0.1:3306)/instant")

type Test struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (t Test) TableName() string {
	return "re_test"
}

func (t Test) Get(DB *sql.DB) (Test, error) {
	//sSql := fmt.Sprintf("select * from test where id = %d", t.ID)
	//rows, err := DB.Query(sSql)
	err := DB.QueryRow("select * from test where id = ?", t.ID).Scan(&t.Title)
	if err != nil && err != sql.ErrNoRows {
		return t, err
	}
	return t, nil
}
