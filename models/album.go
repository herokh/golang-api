package models

import (
	"github.com/herokh/golang-api/dtos"
	"github.com/herokh/golang-api/views"
	"gorm.io/gorm"
)

type Album struct {
	ID     uint    `gorm:"primaryKey;column:id"`
	Title  string  `gorm:"column:title"`
	Artist string  `gorm:"column:artist"`
	Price  float64 `gorm:"column:price"`
	gorm.Model
}

func (album Album) ToEntity() views.Album {
	return views.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}

func (entity Album) FromDto(dto dtos.AlbumDto) Album {
	return Album{
		Title:  dto.Title,
		Artist: dto.Artist,
		Price:  dto.Price,
	}
}
