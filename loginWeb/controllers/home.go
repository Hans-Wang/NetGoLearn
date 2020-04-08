package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"netgo/loginWeb/models"
)

func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	fmt.Println("loginUser", loginUser)
	if loginUser != nil {
		return true
	} else {
		return false
	}
}

func HomeGet(c *gin.Context) {
	isLogin := GetSession(c)
	fmt.Println("home", isLogin)
	session := sessions.Default(c)
	username := session.Get("login_user").(string)
	page := 1

	articleList, err := models.QueryCurrUserArticleWithPage(username, page)
	if err != nil {
		fmt.Println("homeGet() err:", err)
	}
	html := models.GenHomeBlocks(articleList, isLogin)
	c.HTML(http.StatusOK, "home.html", gin.H{"IsLogin": isLogin, "Content": html})
}
