package main

import (
	"github.com/gin-gonic/gin"
	"github.com/williampiv/textweb/internal/textify"
	"html/template"

	"net/http"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/text", textifyPage)
	router.GET("/sites", showSitesPage)
}

func showIndexPage(c *gin.Context) {

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}

func showSitesPage(c *gin.Context) {
	c.HTML(http.StatusOK, "sites.html", gin.H{
		"title": "Text Sites",
	})
}

func textifyPage(c *gin.Context) {
	urlQuery := c.Query("url")
	content := textify.ByURL(urlQuery)
	c.HTML(http.StatusOK, "content.html", gin.H{
		"title":   "Something",
		"payload": template.HTML(content),
	})
}
