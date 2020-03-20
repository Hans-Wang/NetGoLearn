package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//gin 中间件

func StatCost() gin.HandlerFunc  {
	return func(c *gin.Context) {
		t := time.Now()
		c.Set("example", "12345")
		c.Next()
		latency := time.Since(t)
		log.Println(latency)
	}
}

func main()  {
	r := gin.Default()
	r.Use(StatCost())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println(example)

		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})

	r.Run()
}