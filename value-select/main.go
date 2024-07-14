package main

import (
	"embed"
	"fmt"
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

type Make struct {
	Name   string
	Models []string
}

var makes []Make

func init_db() {

	makes = []Make{
		Make{
			Name:   "Audi",
			Models: []string{"A1", "A4", "A6"},
		},
		Make{
			Name:   "BMW",
			Models: []string{"325i", "325ix", "X5"},
		},
		Make{
			Name:   "Toyota",
			Models: []string{"Landcruiser", "Tacoma", "Yaris"},
		},
	}
}

func find_make(name string) *Make {
	var found *Make
	for _, maker := range makes {
		if maker.Name == name {
			found = &maker
			break
		}
	}
	return found
}

func main() {

	init_db()
	fmt.Println(makes)

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

	router.GET("/makes", func(c *gin.Context) {
		var make_names []string

		for _, maker := range makes {
			make_names = append(make_names, maker.Name)
		}
		c.JSON(http.StatusOK, make_names)
	})

	router.GET("/models", func(c *gin.Context) {
		make_name := c.Query("make")

		found := find_make(make_name)
		if found == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"reason": "wrong make"},
			)
			return
		}

		c.JSON(http.StatusOK, found.Models)
	})

	router.Run(":8080")
}
