package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/herokh/go-playground/configs"
	"github.com/herokh/go-playground/controllers"
	"github.com/herokh/go-playground/middlewares"
)

func MapEndpoints(g *gin.Engine, logger *configs.Logger, db *configs.Database, config *configs.AppConfiguration) {

	authController := controllers.NewAuthController(logger, db)
	g.POST("/login", authController.Login)
	protected := g.Group("/", middlewares.AuthorizationMiddleware)

	// Albums
	albumController := controllers.NewAlbumController(logger, db)
	protected.GET("/albums", albumController.GetAlbums)
	protected.GET("/albums/:id", albumController.GetAlbumByID)
	protected.POST("/albums", albumController.PostAlbums)

}
