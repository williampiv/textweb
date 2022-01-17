package main

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
)

var commitVersion string

var router *gin.Engine

func main() {
	log.Println("textWeb Version:", commitVersion)
	router = gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	if err := router.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		log.Fatalln(err)
	}
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	if err := router.Run(); err != nil {
		log.Fatalln(err)
	}

}
