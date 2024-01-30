package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"unsleepNight/models"
)

func Index(c *gin.Context) {
	threads, err := models.Threads()
	if err == nil {
		_, err := session(c)
		if err != nil {
			generateHTML(c, threads, "layout", "navbar", "index")
			return
		} else {
			generateHTML(c, threads, "layout", "auth.navbar", "index")
			return
		}
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Threads": threads,
	})

}
