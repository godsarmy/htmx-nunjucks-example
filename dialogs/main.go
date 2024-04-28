package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.POST("/submit", func(c *gin.Context) {
		prompt := c.Request.Header["Hx-Prompt"]

		c.JSON(http.StatusOK, gin.H{"prompt": prompt})
	})

	router.Run(":8080")
}
