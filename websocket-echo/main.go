package main

import (
	"embed"
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

//go:embed templates/*
var embed_fs embed.FS

type Input struct {
	Message string `json:"message"`
	Control string `json:"control"`
}

type Output struct {
	Message  string `json:"message"`
	Sequence int    `json:"sequence"`
	Error    string `json:"error"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
				"htmx_version":     htmx_version,
				"nunjucks_version": nunjucks_version,
			},
		)
	})

	router.GET("/chat", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer conn.Close()

		sequence := 0
		for {
			// Read message from client
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				break
			}
			if messageType != websocket.TextMessage {
				c.AbortWithStatusJSON(
					http.StatusBadRequest,
					gin.H{"reason": "unsupported message type"},
				)
				break
			}

			var input Input
			var output Output = Output{}

			if err = json.Unmarshal(p, &input); err != nil {
				output.Error = err.Error()
			} else {
				sequence++
				output.Message = input.Message
				output.Sequence = sequence
			}

			reply_bytes, _ := json.Marshal(output)
			// Echo message back to client
			err = conn.WriteMessage(messageType, reply_bytes)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				break
			}

			if input.Control == "stop" {
				break
			}
		}
	})

	router.Run(":8080")
}
