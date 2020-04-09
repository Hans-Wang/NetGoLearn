package routers

import (
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"netgo/loginWeb/controllers"
	"netgo/loginWeb/logger"
	"netgo/loginWeb/middlewares"
	"time"
)

func SetupRouter() *gin.Engine {
	//r := gin.Default()
	r := gin.New()
	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	r.SetFuncMap(template.FuncMap{
		"timeStr": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})

	r.Static("/static", "static")

	r.LoadHTMLGlob("templates/*")

	//设置ssessions 中间件

	//store := cookie.NewStore([]byte("loginuser"))
	//设置redis存储cookie
	store, _ := redis.NewStore(10, "tcp", "47.104.241.166:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	{
		r.GET("/login", controllers.LoginFunc)
		r.POST("/login", controllers.LoginPost)
		//r.GET("/register", controllers.Register)
		r.GET("/register", controllers.RegisterGet)
		r.POST("/register", controllers.RegisterPost)
	}

	{
		basicAuthGroup := r.Group("/", middlewares.BasicAuth())
		basicAuthGroup.GET("/home", controllers.HomeGet)
		basicAuthGroup.GET("/", controllers.HomeGet)
		basicAuthGroup.GET("/logout", controllers.LogoutHandler)

		article := basicAuthGroup.Group("/article")
		{
			article.GET("/add", controllers.AddArticleGet)
			article.POST("/add", controllers.AddArticlePost)
			//文章详情
			article.GET("/show/:id", controllers.ShowArticleGet)
			article.GET("/update", controllers.UpdateArticleGet)
			article.POST("/update", controllers.UpdateArticlePost)
			//删除文章
			article.GET("/delete", controllers.DeleteArticle)
		}
		//相册
		basicAuthGroup.GET("/album", controllers.AlbumGet)
		//文件上传
		basicAuthGroup.POST("/upload", controllers.UploadPost)
	}

	r.GET("/index", controllers.IndexFun)

	return r
}
