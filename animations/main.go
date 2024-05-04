package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var COLORS = []string{"red", "blue", "green", "orange"}

func main() {
	current_index := 0

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.GET("/color", func(c *gin.Context) {

		color := COLORS[current_index]
		current_index++
		if current_index >= len(COLORS) {
			current_index = 0
		}
		c.JSON(http.StatusOK, gin.H{"color": color})
	})

	router.Run(":8080")
}
