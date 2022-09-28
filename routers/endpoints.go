package routers

import (
	"github.com/gin-gonic/gin"
	cf "github.com/herokh/go-playground/configs"
	c "github.com/herokh/go-playground/controllers"
)

func MapEndpoints(g *gin.Engine, logger *cf.Logger, db *cf.Database, config *cf.AppConfiguration) {
	// Albums
	albumController := c.NewAlbumController(logger, db)
	g.GET("/albums", albumController.GetAlbums)
	g.GET("/albums/:id", albumController.GetAlbumByID)
	g.POST("/albums", albumController.PostAlbums)

}
