package main

import (
	"fmt"
	"net/http"

	"wrapdbr"

	"github.com/codemodus/kace"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""
	// ミドルウェアを使用
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		c.Next()
	})

	s := "this is a test sql."
	fmt.Println(kace.Camel(s))
	conn, _ := wrapdbr.Open("postgres", "...", nil)

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":    kace.Camel(s),
			"User-Agent": ua,
		})
	})
	engine.Run(":3000")
}
