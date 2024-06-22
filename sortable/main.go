package main

import (
	"embed"
	"encoding/json"
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

type Item struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type ItemList struct {
	Item []string `json:"item" binding:"required"`
}

var items []Item

func init_db() {

	items = []Item{
		Item{
			Name: "Item 1",
			ID:   1,
		},
		Item{
			Name: "Item 2",
			ID:   2,
		},
		Item{
			Name: "Item 3",
			ID:   3,
		},
		Item{
			Name: "Item 4",
			ID:   4,
		},
		Item{
			Name: "Item 5",
			ID:   5,
		},
	}
}

func find_item(id int) *Item {
	var found *Item
	for _, item := range items {
		if item.ID == id {
			found = &item
			break
		}
	}
	return found
}

func main() {

	init_db()
	fmt.Println(items)

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

	router.POST("/items", func(c *gin.Context) {
		var order_list ItemList

		err := json.NewDecoder(c.Request.Body).Decode(&order_list)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
		}

		if len(order_list.Item) == 0 {
			c.JSON(http.StatusOK, items)
			return
		}

		var sorted []Item
		for _, id := range order_list.Item {
			id_num, err := strconv.Atoi(id)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
			item := find_item(id_num)

			if item != nil {
				sorted = append(sorted, *item)
			}
		}
		c.JSON(http.StatusOK, sorted)
	})

	router.Run(":8080")
}
