package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID uint64
	Username string
}
func main() {
	users := []User{{ID:1,Username:"zhangsan"},{ID:2,Username:"lisi"}}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var sum = 0
		for i:=0;i<=100000;i++{
			sum +=i
		}
		fmt.Println(sum)
		c.JSON(200,users)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		queryParams := c.Query("name")
		fmt.Println("queryParams",queryParams)
		message := "name is "+name+",action is "+action
		c.String(200,message)
	})
	r.Run(":8080")
}