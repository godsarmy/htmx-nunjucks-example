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

const TITLE = "Modal title"
const MESSAGE = "Modal body text goes here."

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

	router.GET("/modal", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"title": TITLE, "message": MESSAGE},
		)
	})

	router.Run(":8080")
}
