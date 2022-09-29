package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/herokh/golang-api/configs"
	"github.com/herokh/golang-api/controllers"
	"github.com/herokh/golang-api/middlewares"
)

func MapEndpoints(g *gin.Engine, logger *configs.Logger, db *configs.Database, config *configs.AppConfiguration) {

	authController := controllers.NewAuthController(logger, db, config)
	g.POST("/login", authController.Login)
	protected := g.Group("/", func(ctx *gin.Context) {
		middlewares.AuthorizationMiddleware(ctx, config)
	})

	// Albums
	albumController := controllers.NewAlbumController(logger, db)
	protected.GET("/albums", albumController.GetAlbums)
	protected.GET("/albums/:id", albumController.GetAlbumByID)
	protected.POST("/albums", albumController.PostAlbums)

}
