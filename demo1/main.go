package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func testHandle(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
func main(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/test", testHandle)

	//参数传递
	r.GET("/user/search", func(c *gin.Context) {
		//默认参数
		//http://127.0.0.1:8080/user/search?address=aaa
		//username := c.DefaultQuery("username","Hans")
		username := c.Query("username")
		address := c.Query("address")

		c.JSON(http.StatusOK, gin.H{
			"message": "Pong",
			"username": username,
			"address": address,
		})
	})

	//以路径的参数传递
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		//http://127.0.0.1:8080/user/search/aaau/ddd
		username := c.Param("username")
		address := c.Param("address")

		c.JSON(http.StatusOK, gin.H{
			"message": "Pong",
			"username": username,
			"address": address,
		})
	})

	//用post传递参数
	r.POST("/user/ss", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")

		c.JSON(http.StatusOK, gin.H{
			"message": "Pong",
			"username": username,
			"address": address,
		})
	})
	r.Run()
}
