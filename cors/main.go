package main

import (
	"embed"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

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
				"nunjucks_version":  nunjucks_version,
				"bootstrap_version": bootstrap_version,
			},
		)
	})
	router.GET("/local_example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "local_example"})
	})

	router_other := gin.Default()
	router_other.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "hx-current-url", "hx-request", "hx-target"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))
	router_other.GET("/other_example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "other_example"})
	})

	go func() { router_other.Run(":8090") }() // other service
	router.Run(":8080")                       // default service
}
