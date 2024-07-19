package main

import (
	"embed"
	"fmt"
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

	router.GET("/head", func(c *gin.Context) {
		head := `
        <head>
          <script src="https://cdn.jsdelivr.net/npm/bootstrap@%s/dist/js/bootstrap.bundle.min.js"></script>
          <link href="https://cdn.jsdelivr.net/npm/bootstrap@%s/dist/css/bootstrap.min.css" rel="stylesheet">
        </head>
        `
		c.Data(
			http.StatusOK,
			"text/html; charset=utf-8",
			[]byte(fmt.Sprintf(head, bootstrap_version, bootstrap_version)),
		)
	})

	router.Run(":8080")
}
