package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

//go:embed templates/*
var embed_fs embed.FS

type Animal struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
	Size string `json:"size" binding:"required"`
}

var animals []Animal

func init_db() {

	animals = []Animal{
		Animal{
			Name: "pigeon",
			Type: "bird",
			Size: "small",
		},
		Animal{
			Name: "chicken",
			Type: "bird",
			Size: "middle",
		},
		Animal{
			Name: "duck",
			Type: "bird",
			Size: "middle",
		},
		Animal{
			Name: "mouse",
			Type: "mammal",
			Size: "small",
		},
		Animal{
			Name: "cat",
			Type: "mammal",
			Size: "middle",
		},
		Animal{
			Name: "tiger",
			Type: "mammal",
			Size: "big",
		},
	}
}

func main() {

	init_db()
	fmt.Println(animals)

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

	router.GET("/:animal_type/", func(c *gin.Context) {
		animal_type := c.Params.ByName("animal_type")

		var result []Animal
		for _, animal := range animals {
			if animal.Type != animal_type {
				continue
			}
			size := c.Query("size")
			if size != "" && animal.Size != size {
				continue
			}
			result = append(result, animal)
		}
		c.JSON(http.StatusOK, result)
	})

	router.Run(":8080")
}
