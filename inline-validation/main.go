package main

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	FirstName string `json:"firstName"    binding:"required"`
	LastName  string `json:"lastName"     binding:"required"`
	Email     string `json:"email"        binding:"required"`
}

type Email struct {
	Email string `json:"email"        binding:"required"`
}

var contacts []Contact

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.Static("/img", "./img")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.POST("/contact/email", func(c *gin.Context) {
		var email Email
		if err := c.Bind(&email); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		if _, err := mail.ParseAddress(email.Email); err != nil {
			c.JSON(
				http.StatusOK,
				gin.H{
					"value":   email.Email,
					"message": "Please enter a valid email address.",
				},
			)
			return
		}

		if email.Email != "test@test.com" {
			c.JSON(
				http.StatusOK,
				gin.H{
					"value":   email.Email,
					"message": "That email is already taken. Please enter another email.",
				},
			)
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": email.Email})
	})

	router.POST("/contact/", func(c *gin.Context) {
		var contact Contact
		if err := c.Bind(&contact); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		contacts = append(contacts, contact)
		c.JSON(http.StatusOK, contact)
	})

	router.Run(":8080")
}
