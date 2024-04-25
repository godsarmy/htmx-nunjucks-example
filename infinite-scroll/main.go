package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Contact struct {
	FirstName string `json:"firstName"    binding:"required"`
	LastName  string `json:"lastName"     binding:"required"`
	Email     string `json:"email"        binding:"required"`
	ID        string `json:"id"           binding:"required"`
}

func sha256hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

    router.Static("/img", "./img")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.GET("/contacts/", func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		page_num, err := strconv.Atoi(page)

		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		var contacts [10]Contact

		for index, _ := range contacts {
			email := fmt.Sprintf("void%d@null.org", (page_num-1)*10+index)
			contacts[index] = Contact{
				FirstName: "Agent",
				LastName:  "Smith",
				Email:     email,
				ID:        sha256hash(email),
			}
		}

		c.JSON(http.StatusOK, gin.H{"page": page_num, "data": contacts})
	})

	router.Run(":8080")
}
