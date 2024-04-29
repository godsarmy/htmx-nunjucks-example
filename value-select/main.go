package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		}
	}
	return found
}

func main() {

	init_db()
	fmt.Println(makes)

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.Static("/img", "./img")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

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
