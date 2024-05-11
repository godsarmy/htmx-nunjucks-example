package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

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

	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
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