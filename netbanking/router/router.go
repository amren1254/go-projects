package router

import (
	"netbanking/controller"
	"netbanking/database"
	middleware "netbanking/middlewares"

	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	//initializing database here
	var db database.DatabaseRepository
	db.InitDatabaseConnection()

	//unauthorized access
	router.GET("/ping", controller.Ping)
	router.POST("/signup", controller.Signup(db))
	router.POST("/login", controller.Login(db))

	//authorized access to apis
	public := router.Group("/api/v1")
	public.Use(middleware.JwtAuthMiddleware())
	// username := middleware.ExtractUsernameFromTokenClaims(c *gin.Context)
	public.GET("/profile", controller.GetProfile(db))
	public.POST("profile", controller.UpdateProfile(db))
	// public.GET("/albums", controller.GetAlbum)
	// public.GET("/albums/:id", controller.GetAlbumsById)
	// public.POST("/albums", controller.PostAlbum)

	return router
}
