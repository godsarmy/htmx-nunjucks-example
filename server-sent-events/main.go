package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS
var message chan string

func main() {
	message = make(chan string)

	go func() {
		for {
			time.Sleep(time.Second)
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			// Send current time to clients message channel
			message <- currentTime
		}
	}()

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

	router.GET("/message", func(c *gin.Context) {
		c.Stream(func(w io.Writer) bool {
			// Stream message to client from message channel
			if currentTime, ok := <-message; ok {
			    data := map[string]string{"time": currentTime}
			    msg, _ := json.Marshal(data)
				c.SSEvent("message", string(msg))
				return true
			}
			return false
		})
	})
	router.Run(":8080")
}
