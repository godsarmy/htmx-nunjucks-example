package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const RESULT = "Did it!"

func main() {

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.POST("/doit", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{"result": RESULT},
		)
	})

	router.Run(":8080")
}
