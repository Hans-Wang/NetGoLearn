package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginFunc(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", nil)
}

//func Register(c *gin.Context){
//	c.HTML(http.StatusOK, "register.html", nil)
//}