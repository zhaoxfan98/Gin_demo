/**
  @Time    : 2021/6/14 18:26
  @Author  : zhaoxfan
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	r.POST("/form", func(c *gin.Context) {
		//表单参数可以通过PostForm()方法获取，该方法默认解析的是 x-www-form-urlencoded 或 from-data 格式的参数
		types := c.DefaultPostForm("type","post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	r.Run()
}
