/**
  @Time    : 2021/6/14 21:22
  @Author  : zhaoxfan
*/
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message":"Hello www.topgoer.com!",
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	return r
}
