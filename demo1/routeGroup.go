package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func login(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func submit(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit",
	})
}

func read(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"message": "read",
	})
}


//绑定参数
type Login struct{
	User string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main(){
	r := gin.Default()
	//路由分组
	v1 := r.Group("v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
		v1.GET("/read", read)

	}


	//绑定参数
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
	})

	r.GET("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.ShouldBindJSON(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user": login.User,
				"password": login.Password,
			})
		}else {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		}
	})


	//gin 渲染
	//json渲染
	r.GET("/someJSON", func(c *gin.Context) {
		//拼接JSON
		c.JSON(http.StatusOK, gin.H{
			"message" : "Hello",
			"status" : http.StatusOK,
		})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		//把结构体渲染成json
		var msg struct{
			Name string `json:"user"`
			Message string
			Number int
		}

		msg.Name = "Hans"
		msg.Message = "Hello JSON"
		msg.Number = 1
		c.JSON(http.StatusOK, msg)
	})

	//XML渲染
	r.GET("/moreXML", func(c *gin.Context) {
		//把结构体渲染成json
		type messagea struct{
			Name string `json:"user"`
			Message string
			Number int
		}

		var msg messagea
		msg.Name = "Hans"
		msg.Message = "Hello XML"
		msg.Number = 1
		c.XML(http.StatusOK, msg)
	})

	r.Static("/static", "./static")
	r.Run()
}
