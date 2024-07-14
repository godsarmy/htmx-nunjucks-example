package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var htmx_ext_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

var COLORS = []string{"red", "blue", "green", "orange"}

func main() {
	current_index := 0
	fade_in := 0

	router := gin.Default()
	templ := template.Must(
		template.New("").Delims("{[{", "}]}").ParseFS(embed_fs, "templates/*.tmpl"),
	)
	router.SetHTMLTemplate(templ)
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html.tmpl",
			gin.H{
				"htmx_version":      htmx_version,
				"htmx_ext_version":  htmx_ext_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})

	router.GET("/color", func(c *gin.Context) {

		color := COLORS[current_index]
		current_index++
		if current_index >= len(COLORS) {
			current_index = 0
		}
		c.JSON(http.StatusOK, gin.H{"color": color})
	})

	router.DELETE("/fade-out", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	router.POST("/fade-in", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"count": fade_in})
		fade_in++
	})

	router.POST("/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"caption": "Submitted"})
	})

	router.Run(":8080")
}
