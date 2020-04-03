package routers

import "github.com/gin-gonic/gin"
import "netgo/loginWeb/controllers"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")
	r.GET("/", controllers.LoginFunc)
	r.GET("/login", controllers.LoginFunc)
	r.POST("/login", controllers.LoginPost)
	//r.GET("/register", controllers.Register)
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register",controllers.RegisterPost)
	r.GET("/index", controllers.IndexFun)
	return r
}