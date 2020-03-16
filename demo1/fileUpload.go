package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	r := gin.Default()

	//文件上传设置内存大小，默认32M,超过这个才会写磁盘
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		//单个文件上传
		file, err := c.FormFile("file")
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("/tmp/%s", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	r.POST("/mult_upload", func(c *gin.Context) {
		//多个文件上传
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files{
			log.Println(file.Filename)
			dst := fmt.Sprintf("/tmp/%s_%d",file.Filename,index)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d  files uploaded!", len(files)),
		})
	})



	r.Run()
}
