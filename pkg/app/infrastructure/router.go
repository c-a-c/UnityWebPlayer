package infrastructure

import (
	"UnityWebPlayer/interfaces/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.LoadHTMLGlob("HTML/*")

	gameController := controllers.NewGameController(NewSqlHandler())

	// file
	router.StaticFS("/play", http.Dir("Unity"))

	// top
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "top.html", gin.H{})
	})

	// gameAPI
	router.GET("/games", func(c *gin.Context) { gameController.Index(c) })
	router.GET("/games/:id", func(c *gin.Context) { gameController.Show(c) })

	// redirect
	router.NoRoute(func(c *gin.Context) {
		c.HTML(200, "error.html", gin.H{})
	})

	Router = router
}
