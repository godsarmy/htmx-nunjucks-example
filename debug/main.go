package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

//go:embed templates/*
var embed_fs embed.FS

func main() {

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
				"htmx_version":     htmx_version,
				"nunjucks_version": nunjucks_version,
			},
		)
	})

	router.GET("/get_example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "debug example"})
	})

	router.Run(":8080")
}
