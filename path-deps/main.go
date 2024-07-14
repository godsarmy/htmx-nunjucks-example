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

	router.POST("/data/:data_type", func(c *gin.Context) {
		data_type := c.Params.ByName("data_type")
		c.JSON(http.StatusOK, gin.H{"message": data_type})
	})

	router.POST("/path1/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "post path1"})
	})

	router.GET("/path2/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "get path2"})
	})

	router.Run(":8080")
}
