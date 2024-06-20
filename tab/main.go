package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"
var bootstrap_version = "latest"

//go:embed templates/*
var embed_fs embed.FS

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
            break
		}
	}
	return found
}

func main() {

	init_db()
	fmt.Println(tabs)

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
