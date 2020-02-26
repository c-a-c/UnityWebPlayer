package infrastructure

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	router.LoadHTMLGlob("../../UnityHTML/*.html")
	router.LoadHTMLGlob("../../UnityHTML/admin/*.html")

	// userController := controllers.NewUserController(NewSqlHandler())

	// router.POST("/users", func(c *gin.Context) { userController.Create(c) })
	// router.GET("/users", func(c *gin.Context) { userController.Index(c) })
	// router.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.GET("/hello", func(ctx *gin.Context) {
        ctx.HTML(200, "helloworld.html", gin.H{})
    })

    // admin
    router.GET("/admin/upload", func(ctx *gin.Context) {
        ctx.HTML(200, "gameUpload.html", gin.H{})
    })

	Router = router
}
