package services

import (
	c "github.com/herokh/golang-api/configs"
	m "github.com/herokh/golang-api/models"
)

type AlbumService interface {
	GetAlbums() []m.Album
	GetAlbumByID(id uint) m.Album
	PostAlbum(model m.Album) uint
}

type albumService struct {
	Logger *c.Logger
	Db     *c.Database
}

func NewAlbumService(logger *c.Logger, db *c.Database) AlbumService {
	return &albumService{
		Logger: logger,
		Db:     db,
	}
}

func (service *albumService) GetAlbums() []m.Album {
	var albums []m.Album
	service.Db.DbContext.Find(&albums)
	return albums
}

func (service *albumService) GetAlbumByID(id uint) m.Album {
	var album m.Album
	service.Db.DbContext.Find(&album, id)
	return album
}

func (service *albumService) PostAlbum(model m.Album) uint {
	service.Db.DbContext.Create(&model)
	return model.ID
}
