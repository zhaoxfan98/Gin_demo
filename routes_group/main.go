/**
  @Time    : 2021/6/14 20:41
  @Author  : zhaoxfan
*/
//routes group是为了管理一些相同的URL
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	//1. 创建路由
	//默认使用了两个中间件 Logger() Recovery()
	r := gin.Default()
	//路由组1  处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	r.Run(":8000")
}

func login(c *gin.Context){
	name := c.DefaultQuery("name","jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context){
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}