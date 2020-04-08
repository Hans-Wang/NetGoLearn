package routers

import (
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"netgo/loginWeb/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")

	//设置ssessions 中间件

	//store := cookie.NewStore([]byte("loginuser"))
	//设置redis存储cookie
	store, _ := redis.NewStore(10, "tcp", "47.104.241.166:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/login", controllers.LoginFunc)
	r.POST("/login", controllers.LoginPost)
	//r.GET("/register", controllers.Register)
	r.GET("/register", controllers.RegisterGet)
	r.POST("/register", controllers.RegisterPost)
	r.GET("/index", controllers.IndexFun)
	r.GET("/", controllers.HomeGet)

	article := r.Group("/article")
	{
		article.GET("/add", controllers.AddArticleGet)
		article.POST("/add", controllers.AddArticlePost)
	}
	return r
}
