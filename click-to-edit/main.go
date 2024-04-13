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
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	router.GET("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"reason": "NOT_FOUND"})
		}
		if len(contacts) <= id {
			c.JSON(http.StatusNotFound, gin.H{"reason": "NOT_FOUND"})
		}
		c.JSON(http.StatusOK, contacts[id])
	})

	router.PUT("/contact/:contact_id", func(c *gin.Context) {
		contact_id := c.Params.ByName("contact_id")
		id, err := strconv.Atoi(contact_id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"reason": "NOT_FOUND"})
		}
		if len(contacts) <= id {
			c.JSON(http.StatusNotFound, gin.H{"reason": "NOT_FOUND"})
		}

		if err := c.Bind(&contacts[id]); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"reason": err})
		}
		c.JSON(http.StatusOK, contacts[id])
	})

	router.Run(":8080")
}
