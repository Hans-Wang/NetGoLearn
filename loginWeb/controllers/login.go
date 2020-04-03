package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"netgo/loginWeb/models"
	"netgo/loginWeb/utils"
)

func LoginPost(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username, password)

	id := models.QuerUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)

	if id > 0 {
		c.Redirect(http.StatusFound, "/index")
		//c.JSON(http.StatusOK, gin.H{"code":0, "message":"登录成功"})
	}else {
		c.JSON(http.StatusOK, gin.H{"code":0, "message":"登录失败"})
	}

}