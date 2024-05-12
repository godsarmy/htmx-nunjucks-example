package main

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var htmx_version = "latest"
var nunjucks_version = "3.2.4"

//go:embed templates/*
var embed_fs embed.FS

type Progress struct {
	Percent int `json:"percent"`
	Status  int `json:"status"`
}

const UNKNOWN = -1
const RUNNING = 0
const STOPPED = 1

func main() {
	var jobset map[string]Progress = make(map[string]Progress)

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

	router.GET("/job/:job_id", func(c *gin.Context) {
		var status int
		var percent int

		job_id := c.Params.ByName("job_id")
		progress, ok := jobset["job_id"]
		if !ok {
			status = UNKNOWN
		} else {
			status = progress.Status
			percent = progress.Percent
		}

		c.JSON(http.StatusOK, gin.H{"job_id": job_id, "status": status, "percent": percent})
	})

	router.POST("/job/:job_id", func(c *gin.Context) {
		job_id := c.Params.ByName("job_id")
		var progress Progress
		if err := c.Bind(&progress); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		old_progress, ok := jobset["job_id"]
		if !ok {
			jobset["job_id"] = progress
		} else {
			if old_progress.Percent < 100 {
				progress.Status = old_progress.Percent
			}
			jobset["job_id"] = progress
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"job_id":  job_id,
				"status":  jobset["job_id"].Status,
				"percent": jobset["job_id"].Percent,
			},
		)
	})

	router.GET("/job/:job_id/progress", func(c *gin.Context) {
		job_id := c.Params.ByName("job_id")

		progress, ok := jobset["job_id"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"reason": "wrong job_id"})
		}

		if progress.Status == RUNNING {
			progress.Percent += 11
			if progress.Percent >= 100 {
				progress.Status = STOPPED
			}
			jobset["job_id"] = progress
		}

		if progress.Status == STOPPED {
			c.Header("HX-Trigger", "done")
		} else {
			c.Header("HX-Trigger", "")
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"job_id":  job_id,
				"status":  jobset["job_id"].Status,
				"percent": jobset["job_id"].Percent,
			},
		)
	})

	router.Run(":8080")
}
