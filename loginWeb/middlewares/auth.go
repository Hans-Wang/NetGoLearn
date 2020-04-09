package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUser := session.Get("login_user")
		if loginUser == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Set("is_login", true)
		c.Set("login_user", loginUser)
		c.Next()
	}
}
