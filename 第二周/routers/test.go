package routers

import (
	"err/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Test struct{}

func NewTest() Test {
	return Test{}
}

func (t Test) Get(c *gin.Context) {
	v, _ := strconv.Atoi(c.Param("id"))
	param := service.Test{ID: v}

	svc := service.New(c.Request.Context())
	test, err := svc.GetTest(&param)
	if err != nil {

		return
	}
	c.JSON(200, test)

	return
}
