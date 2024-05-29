package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

type Server struct {
	Name string `json:"name" binding:"required"`
	Id   int    `json:"id"   binding:"required"`
}

type ServerDetail struct {
	Id     int `json:"id"     binding:"required"`
	Cpu    int `json:"cpu"    binding:"required"`
	Memory int `json:"memory" binding:"required"`
	Disk   int `json:"disk"   binding:"required"`
}

var servers []Server
var serverDetails []ServerDetail

func init_db() {

	servers = []Server{
		Server{
			Name: "server-zero",
			Id:   0,
		},
		Server{
			Name: "server-one",
			Id:   1,
		},
		Server{
			Name: "server-two",
			Id:   2,
		},
	}

	serverDetails = []ServerDetail{
		ServerDetail{
			Cpu:    4,
			Memory: 16384,
			Disk:   524288,
			Id:     0,
		},
		ServerDetail{
			Cpu:    8,
			Memory: 16384,
			Disk:   524288,
			Id:     1,
		},
		ServerDetail{
			Cpu:    8,
			Memory: 32768,
			Disk:   524288,
			Id:     2,
		},
	}
}

func main() {

	init_db()
	fmt.Println(servers)
	fmt.Println(serverDetails)

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

	router.GET("/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, servers)
	})

	router.GET("/server/:server_id", func(c *gin.Context) {
		server_id := c.Params.ByName("server_id")
		id, err := strconv.Atoi(server_id)

		if err != nil {
			c.AbortWithError(http.StatusNotFound, err)
		}
		if len(serverDetails) <= id {
			c.AbortWithError(http.StatusNotFound, err)
		}
		c.JSON(http.StatusOK, serverDetails[id])
	})

	router.Run(":8080")
}
