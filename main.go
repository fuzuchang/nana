package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	ID   uint64
	Name string
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		users := []User{{ID: 123, Name: "张三"}, {ID: 456, Name: "李四"}}

		c.JSON(200, users)

	})
	r.Run(":7999")
}
