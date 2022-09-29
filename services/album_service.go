package services

import (
	c "github.com/herokh/golang-api/configs"
	m "github.com/herokh/golang-api/models"
	"github.com/herokh/golang-api/repositories"
)

type AlbumService interface {
	GetAlbums() []m.Album
	GetAlbumByID(id uint) m.Album
	PostAlbum(model m.Album) uint
}

type albumService struct {
	Logger          *c.Logger
	AlbumRepository repositories.AlbumRepository
}

func NewAlbumService(logger *c.Logger, albumRepository repositories.AlbumRepository) AlbumService {
	return &albumService{
		Logger:          logger,
		AlbumRepository: albumRepository,
	}
}

func (service *albumService) GetAlbums() []m.Album {
	albums := service.AlbumRepository.GetAll()
	return albums
}

func (service *albumService) GetAlbumByID(id uint) m.Album {
	album := service.AlbumRepository.GetByID(id)
	return album
}

func (service *albumService) PostAlbum(model m.Album) uint {
	id := service.AlbumRepository.Post(model)
	return id
}
