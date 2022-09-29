package repositories

import (
	"github.com/herokh/golang-api/configs"
	"github.com/herokh/golang-api/models"
)

type AlbumRepository interface {
	GetAll() []models.Album
	GetByID(id uint) models.Album
	Post(album models.Album) uint
}

type albumRepository struct {
	Logger *configs.Logger
	Db     *configs.Database
}

func NewAlbumRepository(logger *configs.Logger, db *configs.Database) AlbumRepository {
	return &albumRepository{
		Logger: logger,
		Db:     db,
	}
}

func (repo *albumRepository) GetByID(id uint) models.Album {
	var album models.Album
	repo.Db.DbContext.Find(&album, id)
	return album
}

func (repo *albumRepository) GetAll() []models.Album {
	var albums []models.Album
	repo.Db.DbContext.Find(&albums)
	return albums
}

func (repo *albumRepository) Post(album models.Album) uint {
	repo.Db.DbContext.Create(&album)
	return album.ID
}
