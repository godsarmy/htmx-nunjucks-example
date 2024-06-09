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
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

type Contact struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required"`
}

var contacts []Contact

func init_db() {

	contacts = []Contact{
		Contact{
			Name:  "foo bar",
			Email: "foo@gmail.com",
		},
		Contact{
			Name:  "Joe Blow",
			Email: "joe@blow.com",
		},
	}
}

func main() {

	init_db()
	fmt.Println(contacts)

	router := gin.Default()
	templ := template.Must(
		template.New("").Delims("{[{", "}]}").ParseFS(embed_fs, "templates/*.tmpl"),
	)
	router.SetHTMLTemplate(templ)
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html.tmpl",
			gin.H{},
		)
	})
	router.GET("/expand-target", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"expand-target.tmpl",
			gin.H{
				"htmx_version":      htmx_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})
	router.GET("/oob-response", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"oob-response.tmpl",
			gin.H{
				"htmx_version":      htmx_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})
	router.GET("/trigger-events", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"trigger-events.tmpl",
			gin.H{
				"htmx_version":      htmx_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})
	router.GET("/path-deps", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"path-deps.tmpl",
			gin.H{
				"htmx_version":      htmx_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})

	router.POST("/contacts", func(c *gin.Context) {
		var newContact Contact

		if err := c.Bind(&newContact); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		contacts = append(contacts, newContact)
		c.Header("HX-Trigger", "newContact")
		c.JSON(http.StatusOK, contacts)
	})

	router.GET("/contacts", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	router.Run(":8080")
}
