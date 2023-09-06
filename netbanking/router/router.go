package router

import (
	"netbanking/controller"
	"netbanking/database"
	middleware "netbanking/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	//initializing database here
	var db database.DatabaseRepository
	db.InitDatabaseConnection()
	//unauthorized access
	router.GET("/ping", controller.Ping)
	router.POST("/signup", controller.Signup(db))
	router.POST("/login", controller.Login(db))

	//authorized access to apis
	public := router.Group("/api")
	public.Use(middleware.JwtAuthMiddleware())
	// public.GET("/albums", controller.GetAlbum)
	// public.GET("/albums/:id", controller.GetAlbumsById)
	// public.POST("/albums", controller.PostAlbum)

	return router
}
