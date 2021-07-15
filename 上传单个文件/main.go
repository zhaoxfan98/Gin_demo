/**
  @Time    : 2021/6/14 20:14
  @Author  : zhaoxfan
*/
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main(){
	r := gin.Default()
	//限制上传最大尺寸
	r.POST("/upload", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil{
			log.Printf("Error when try to get file:%v",err)
		}
		//headers.Size获取文件大小
		if headers.Size > 1024*1024*2{
			fmt.Println("too large")
			return
		}
		//headers.Header.Get("Content-type")获取上传文件的类型
		if headers.Header.Get("Content-type") != "image/png" {
			fmt.Println("only png")
			return
		}
		c.SaveUploadedFile(headers, headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})
	r.Run()
}
