package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var htmx_ext_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/* img/*
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
				"htmx_version":      htmx_version,
				"htmx_ext_version":  htmx_ext_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})

	var image_fs, _ = fs.Sub(embed_fs, "img")
	router.StaticFS("/img", http.FS(image_fs))

	router.GET("/graph", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"title": "Tokyo Climate", "path": "/img/tokyo.png"})
	})

	router.Run(":8080")
}
