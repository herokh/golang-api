package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	c "github.com/herokh/go-playground/configs"
	m "github.com/herokh/go-playground/models"
	s "github.com/herokh/go-playground/services"
	u "github.com/herokh/go-playground/utils"
	v "github.com/herokh/go-playground/views"
)

type AlbumController struct {
	Logger *c.Logger
	Db     *c.Database
}

func NewAlbumController(logger *c.Logger, db *c.Database) *AlbumController {
	return &AlbumController{
		Logger: logger,
		Db:     db,
	}
}

func (controller *AlbumController) GetAlbums(c *gin.Context) {
	controller.Logger.Info().Msg("Running GetAlbums")
	service := s.NewAlbumService(controller.Logger, controller.Db)
	albums := service.GetAlbums()
	views := u.Map(albums, func(m m.Album) v.Album {
		return m.ToEntity()
	})
	c.IndentedJSON(http.StatusOK, views)
}

func (controller *AlbumController) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	iduint, _ := strconv.ParseUint(id, 10, 32)
	controller.Logger.Info().Msg(fmt.Sprintf("Running GetAlbumByID %s", id))
	service := s.NewAlbumService(controller.Logger, controller.Db)
	album := service.GetAlbumByID(uint(iduint))
	if album.ID != 0 {
		c.IndentedJSON(http.StatusOK, album.ToEntity())
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (controller *AlbumController) PostAlbums(c *gin.Context) {
	controller.Logger.Info().Msg("Running PostAlbums")
	var view v.Album
	json.NewDecoder(c.Request.Body).Decode(&view)
	model := m.Album{}.FromEntity(view)
	service := s.NewAlbumService(controller.Logger, controller.Db)
	id := service.PostAlbum(model)
	view.ID = id
	c.IndentedJSON(http.StatusCreated, view)
}
