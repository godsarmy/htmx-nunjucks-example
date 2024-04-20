package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	FirstName string `json:"firstName"    binding:"required"`
	LastName  string `json:"lastName"     binding:"required"`
	Email     string `json:"email"        binding:"required"`
	Active    bool   `json:"active"       binding:"required"`
	ID        int    `json:"id"           binding:"required"`
}

var contacts []Contact

func init_db() {

	contacts = []Contact{
		Contact{
			ID:        0,
			FirstName: "Joe",
			LastName:  "Smith",
			Email:     "joe@smith.org",
			Active:    true,
		},
		Contact{
			ID:        1,
			FirstName: "Angie",
			LastName:  "MacDowell",
			Email:     "angie@macdowell.org",
			Active:    true,
		},
		Contact{
			ID:        2,
			FirstName: "Fuqua",
			LastName:  "Tarkenton",
			Email:     "fuqua@tarkenton.org",
			Active:    true,
		},
		Contact{
			ID:        3,
			FirstName: "Kim",
			LastName:  "Yee",
			Email:     "kim@yee.org",
			Active:    false,
		},
	}
}

func remove(l []Contact, id int) []Contact {
	for i, other := range l {
		if other.ID == id {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func main() {

	init_db()
	fmt.Println(contacts)

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, contacts)
	})

	router.GET("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		var found *Contact
		for _, contact := range contacts {
			if contact.ID == id {
				found = &contact
			}
		}

		if found == nil {
			c.AbortWithError(http.StatusNotFound, err)
			return
		}

		c.JSON(http.StatusOK, found)
	})

	router.DELETE("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		contacts = remove(contacts, id)
		c.JSON(http.StatusOK, gin.H{})
	})

	router.Run(":8080")
}
