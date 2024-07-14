package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var htmx_ext_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

type Contact struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
	Email     string `json:"email"     binding:"required"`
}

var contacts []Contact

func init_db() {

	contacts = []Contact{
		Contact{
			FirstName: "foo",
			LastName:  "bar",
			Email:     "foo@gmail.com",
		},
		Contact{
			FirstName: "Joe",
			LastName:  "Blow",
			Email:     "joe@blow.com",
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
			gin.H{
				"htmx_version":      htmx_version,
				"htmx_ext_version":  htmx_ext_version,
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})

	router.GET("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
		if len(contacts) <= id {
			c.AbortWithError(http.StatusNotFound, err)
		}
		c.JSON(http.StatusOK, contacts[id])
	})

	router.PUT("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
		if len(contacts) <= id {
			c.AbortWithError(http.StatusNotFound, err)
		}

		if err := c.Bind(&contacts[id]); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusOK, contacts[id])
	})

	router.Run(":8080")
}
