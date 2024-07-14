package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var htmx_ext_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

type Data struct {
	Age string `json:"age" binding:"required"`
}

func main() {

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

	router.POST("/guess", func(c *gin.Context) {
		var data Data
		if err := c.Bind(&data); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		age, err := strconv.Atoi(data.Age)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"message": "age is not integer"},
			)
			return
		}

		if age < 0 {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"message": "age must be positive"},
			)
			return
		}

		year, _, _ := time.Now().Date()
		birth_year := year - age
		c.JSON(
			http.StatusOK,
			gin.H{"message": fmt.Sprintf("You are born in %d", birth_year)},
		)
	})

	router.Run(":8080")
}
