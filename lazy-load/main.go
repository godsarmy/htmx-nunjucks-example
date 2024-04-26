package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.Static("/img", "./img")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.GET("/graph", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"title": "Tokyo Climate", "path": "/img/tokyo.png"})
	})

	router.Run(":8080")
}
