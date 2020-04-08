package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"netgo/loginWeb/models"
	"netgo/loginWeb/utils"
)

func LoginPost(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username, password)

	id := models.QuerUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)

	if id > 0 {
		//id > 0说明登录成功，给响应种上Cookie
		session := sessions.Default(c)
		session.Set("login_user", username)
		session.Save()
		c.Redirect(http.StatusFound, "/")
		//c.JSON(http.StatusOK, gin.H{"code":0, "message":"登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}

}
