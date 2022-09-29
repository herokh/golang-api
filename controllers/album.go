package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	c "github.com/herokh/golang-api/configs"
	d "github.com/herokh/golang-api/dtos"
	m "github.com/herokh/golang-api/models"
	s "github.com/herokh/golang-api/services"
	u "github.com/herokh/golang-api/utils"
	v "github.com/herokh/golang-api/views"
)

type AlbumController interface {
	GetAlbums(c *gin.Context)
	GetAlbumByID(c *gin.Context)
	PostAlbums(c *gin.Context)
}

type albumController struct {
	Logger *c.Logger
	Db     *c.Database
}

func NewAlbumController(logger *c.Logger, db *c.Database) AlbumController {
	return &albumController{
		Logger: logger,
		Db:     db,
	}
}

func (controller *albumController) GetAlbums(c *gin.Context) {
	controller.Logger.Info().Msg("Running GetAlbums")
	service := s.NewAlbumService(controller.Logger, controller.Db)
	albums := service.GetAlbums()
	views := u.Map(albums, func(m m.Album) v.Album {
		return m.ToEntity()
	})
	c.JSON(http.StatusOK, views)
}

func (controller *albumController) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	iduint, _ := strconv.ParseUint(id, 10, 32)
	controller.Logger.Info().Msg(fmt.Sprintf("Running GetAlbumByID %s", id))
	service := s.NewAlbumService(controller.Logger, controller.Db)
	album := service.GetAlbumByID(uint(iduint))
	if album.ID == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, album.ToEntity())

}

func (controller *albumController) PostAlbums(c *gin.Context) {
	controller.Logger.Info().Msg("Running PostAlbums")
	var dto d.AlbumDto
	err := c.ShouldBind(&dto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	model := m.Album{}.FromDto(dto)
	service := s.NewAlbumService(controller.Logger, controller.Db)
	id := service.PostAlbum(model)
	c.JSON(http.StatusCreated, id)
}
