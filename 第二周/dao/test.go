package dao

import (
	"err/model"
)

type Test struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (d *Dao) GetTest(id int) (model.Test, error) {
	test := model.Test{ID: id}
	return test.Get(d.engine)
}
