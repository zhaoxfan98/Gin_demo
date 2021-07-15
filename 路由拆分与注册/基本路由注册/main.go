package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.GET("/topgoer", helloHandler)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n",err)
	}

}

func helloHandler(c *gin.Context) {
	// H is a shortcut for map[string]interface{}
	//type H map[string]interface{}
	c.JSON(http.StatusOK, gin.H{
		"message":"Hello www.topgoer.com",
	})
}