package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tab struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

var tabs []Tab

func init_db() {

	tabs = []Tab{
		Tab{
			Name:    "tab1",
			Content: "Commodo normcore truffaut VHS duis gluten-free keffiyeh iPhone taxidermy godard ramps anim pour-over.",
		},
		Tab{
			Name:    "tab2",
			Content: "Kitsch fanny pack yr, farm-to-table cardigan cillum commodo reprehenderit plaid dolore cronut meditation.",
		},
		Tab{
			Name:    "tab3",
			Content: "Aute chia marfa echo park tote bag hammock mollit artisan listicle direct trade. Raw denim flexitarian eu godard etsy.",
		},
	}
}

func find_tab(tab_name string) *Tab {
	var found *Tab
	for _, tab := range tabs {
		if tab.Name == tab_name {
			found = &tab
		}
	}
	return found
}

func main() {

	init_db()
	fmt.Println(tabs)

	router := gin.Default()
	router.Delims("{[{", "}]}")
	router.LoadHTMLGlob("./templates/*.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{})
	})

	router.GET("/tab", func(c *gin.Context) {
		var tab_names []string
		for _, tab := range tabs {
			tab_names = append(tab_names, tab.Name)
		}
		c.JSON(http.StatusOK, tab_names)
	})

	router.GET("/tab/:tab_id", func(c *gin.Context) {
		tab_id := c.Params.ByName("tab_id")

		found := find_tab(tab_id)

		if found == nil {
			c.AbortWithStatusJSON(
				http.StatusNotFound,
				gin.H{"reason": "wrong tab_id"},
			)
			return
		}

		c.JSON(http.StatusOK, found)
	})

	router.Run(":8080")
}