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

//go:embed templates/* script/*
var embed_fs embed.FS

type Rating struct {
	Value string `json:"value" binding:"required"`
}

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

	var script_fs, _ = fs.Sub(embed_fs, "script")
	router.StaticFS("/script", http.FS(script_fs))

	router.GET("/get_example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "shoelace example"})
	})

	var my_rating Rating = Rating{Value: "3"}
	router.GET("/rating", func(c *gin.Context) {
		c.JSON(http.StatusOK, my_rating)
	})
	router.POST("/rating", func(c *gin.Context) {
		var rating Rating
		if err := c.Bind(&rating); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		my_rating.Value = rating.Value
		c.JSON(http.StatusOK, my_rating)
	})

	router.Run(":8080")
}
