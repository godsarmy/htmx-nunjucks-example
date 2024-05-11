package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html.tmpl",
			gin.H{
				"htmx_version":     htmx_version,
				"nunjucks_version": nunjucks_version,
			},
		)
	})

	router.POST("/submit", func(c *gin.Context) {
		prompt := c.Request.Header["Hx-Prompt"]

		c.JSON(http.StatusOK, gin.H{"prompt": prompt})
	})

	router.Run(":8080")
}
