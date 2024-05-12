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
var nunjucks_version = "3.2.4"

//go:embed templates/*
var embed_fs embed.FS

type Contact struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"  binding:"required"`
	Email     string `json:"email"     binding:"required"`
	ID        int    `json:"id"`
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

func find_contact(id int) *Contact {
	var found *Contact
	for _, contact := range contacts {
		if contact.ID == id {
			found = &contact
		}
	}
	return found
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

	router.GET("/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	router.GET("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		found := find_contact(id)
		if found == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"reason": "wrong contact_id"},
			)
			return
		}

		c.JSON(http.StatusOK, found)
	})

	router.PUT("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		found := find_contact(id)
		if found == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"reason": "wrong contact_id"},
			)
			return
		}

		if err := c.Bind(found); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, found)
	})

	router.Run(":8080")
}
