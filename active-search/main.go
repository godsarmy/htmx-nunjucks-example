package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

//go:embed templates/* img/*
var embed_fs embed.FS

type Contact struct {
	ID        int    `json:"id"        binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var contacts []Contact

func init_db() {

	contacts = []Contact{
		Contact{
			ID:        0,
			FirstName: "Joe",
			LastName:  "Smith",
			Email:     "joe@smith.org",
		},
		Contact{
			ID:        1,
			FirstName: "Angie",
			LastName:  "MacDowell",
			Email:     "angie@macdowell.org",
		},
		Contact{
			ID:        2,
			FirstName: "Fuqua",
			LastName:  "Tarkenton",
			Email:     "fuqua@tarkenton.org",
		},
		Contact{
			ID:        3,
			FirstName: "Kim",
			LastName:  "Yee",
			Email:     "kim@yee.org",
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
				"htmx_version":     htmx_version,
				"nunjucks_version": nunjucks_version,
			},
		)
	})

	var image_fs, _ = fs.Sub(embed_fs, "img")
	router.StaticFS("/img", http.FS(image_fs))

	router.GET("/contacts/", func(c *gin.Context) {
		search := c.Query("search")
		if search == "" {
			c.JSON(http.StatusOK, contacts)
			return
		}

		var selected_contacts []Contact
		for _, contact := range contacts {
			if strings.Contains(contact.FirstName, search) ||
				strings.Contains(contact.LastName, search) ||
				strings.Contains(contact.Email, search) {
				selected_contacts = append(selected_contacts, contact)
			}
		}
		c.JSON(http.StatusOK, selected_contacts)
	})

	router.Run(":8080")
}
