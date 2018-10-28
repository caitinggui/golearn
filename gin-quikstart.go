package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 程序的入口
// 这边可以用来建立路由
func main() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path" + c.Request.URL.Path)

	})

	r.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + cCp.Request.URL.Path)

		}()

	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
