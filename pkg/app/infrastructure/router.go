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

	// router.POST("/games", func(c *gin.Context) { userController.Create(c) })
	router.GET("/games", func(c *gin.Context) { gameController.Index(c) })
	router.GET("/games/:id", func(c *gin.Context) { gameController.Show(c) })

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "top.html", gin.H{})
	})

	router.NoRoute(func (c *gin.Context) {
		c.HTML(http.StatusOK, "error.html", nil)
	})


	Router = router
}
