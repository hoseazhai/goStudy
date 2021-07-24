package main

import (
	"err/global"
	"err/model"
	"err/routers"
	"github.com/gin-gonic/gin"
	"log"
)

//var (
//	DBEngine *sql.DB
//)

func init() {
	err := setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}
func main() {
	r := gin.Default()
	test := routers.NewTest()
	r.GET("/:id", test.Get)
	r.Run()
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine()
	if err != nil {
		return err
	}
	return nil
}
