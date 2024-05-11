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

	router.Static("/img", "./img")
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

	router.GET("/graph", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"title": "Tokyo Climate", "path": "/img/tokyo.png"})
	})

	router.Run(":8080")
}
