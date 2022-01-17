package main

import (
	"github.com/gin-gonic/gin"
	"github.com/williampiv/textweb/internal/textify"
	"html/template"
	"log"
	"strings"

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
	title, content, err := textify.ByURL(urlQuery)
	if err != nil {
		log.Println(err)
	}
	// TODO: maybe cleanup how these string replacements occur
	// All links should stay in textWeb if possible
	content = strings.Replace(content, "href=\"", "href=\"/text?url=", -1)
	// No giant images TODO: image compression
	content = strings.Replace(content, "<img src=", "<img width=\"25%\" src=", -1)
	// Let's keep it all in one window
	content = strings.Replace(content, " target=\"_blank\"", "", -1)
	c.HTML(http.StatusOK, "content.html", gin.H{
		"title":   title,
		"payload": template.HTML(content),
	})
}
