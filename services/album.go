package services

import (
	c "github.com/herokh/go-playground/configs"
	m "github.com/herokh/go-playground/models"
)

type AlbumService struct {
	Logger *c.Logger
	Db     *c.Database
}

func NewAlbumService(logger *c.Logger, db *c.Database) *AlbumService {
	return &AlbumService{
		Logger: logger,
		Db:     db,
	}
}

func (service *AlbumService) GetAlbums() []m.Album {
	var albums []m.Album
	service.Db.DbContext.Find(&albums)
	return albums
}

func (service *AlbumService) GetAlbumByID(id uint) m.Album {
	var album m.Album
	service.Db.DbContext.Find(&album, id)
	return album
}

func (service *AlbumService) PostAlbum(model m.Album) uint {
	service.Db.DbContext.Create(&model)
	return model.ID
}
