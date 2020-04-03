package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"netgo/loginWeb/models"
	"netgo/loginWeb/utils"
	"time"
)

func RegisterGet(c *gin.Context){
	c.HTML(http.StatusOK, "register.html", nil)
}

func RegisterPost(c *gin.Context){
	/*
		1，取出请求的数
	判断注册是否重复--->拿用户名去数据库里查一下
	写入数据库
	 */
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	email := c.PostForm("email")
	fmt.Println(username, password, repassword, email)
	id := models.QuerUserWithUsername(username)
	fmt.Println("id:",id)

	if id >0 {
		c.JSON(http.StatusOK, gin.H{"code:":0,"message":"用户名已存在"})
		return
	}
	
	password = utils.MD5(password)
	fmt.Println("MD5后:", password)
	
	user := models.User{
		Username:   username,
		Password:   password,
		Email: email,
		Status:     0,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err := models.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"注册失败"})
	}else{
		c.JSON(http.StatusOK, gin.H{"code":1,"message":"注册成功"})
	}

}
