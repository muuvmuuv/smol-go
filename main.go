package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	//
	// Routes
	//

	router.GET("/", func(c *gin.Context) {
		r := c.Request

		c.JSON(http.StatusOK, gin.H{
			"Mode":            mode,
			"RemoteAddr":      r.RemoteAddr,
			"User-Agent":      r.Header.Get("User-Agent"),
			"Cookie":          r.Header.Get("Cookie"),
			"Connection":      r.Header.Get("Connection"),
			"Cache-Control":   r.Header.Get("Cache-Control"),
			"Accept-Language": r.Header.Get("Accept-Language"),
			"Accept-Encoding": r.Header.Get("Accept-Encoding"),
			"Accept":          r.Header.Get("Accept"),
		})
	})

	router.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Welcome!",
		})
	})

	//
	// Start
	//

	if devEnv {
		router.Run(":" + port)
	} else {
		router.Run()
	}
}
