package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"netgo/loginWeb/logger"
	"netgo/loginWeb/models"
	"time"
)

func AddArticleGet(c *gin.Context) {
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"Islogin": islogin})
}

func AddArticlePost(c *gin.Context) {
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	session := sessions.Default(c)
	loginUser := session.Get("login_user")
	username := loginUser.(string)
	//fmt.Printf("title:%s, tags:%s\n", title, tags)

	//实例化model, 将它输入到数据库中
	art := models.Article{
		Title:      title,
		Tags:       tags,
		Short:      short,
		Content:    content,
		Author:     username,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, err := models.AddArticle(art)
	response := gin.H{}
	if err == nil {
		response = gin.H{"code": 1, "message": "OK"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}
	c.JSON(http.StatusOK, response)
}

func ShowArticleGet(c *gin.Context) {

}

func UpdateArticleGet(c *gin.Context) {
	isLogin := c.MustGet("is_login")
	idStr := c.Query("id")

	article, err := models.QueryArticleWithId(idStr)
	if err != nil {
		return
	}

	if article == nil {
		c.String(http.StatusOK, "bad id")
		return
	}

	c.HTML(http.StatusOK, "write_article.html", gin.H{
		"isLogin": isLogin, "article": article})
}

func UpdateArticlePost(c *gin.Context) {

}

func DeleteArticle(c *gin.Context) {

}

func AlbumGet(c *gin.Context) {

}

func UploadPost(c *gin.Context) {

}
