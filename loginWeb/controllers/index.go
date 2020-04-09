package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexFun(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "用户已退出",
	})

}
