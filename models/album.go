package models

import (
	v "github.com/herokh/go-playground/views"
	"gorm.io/gorm"
)

type Album struct {
	ID     uint    `gorm:"primaryKey;column:id"`
	Title  string  `gorm:"column:title"`
	Artist string  `gorm:"column:artist"`
	Price  float64 `gorm:"column:price"`
	gorm.Model
}

func (album Album) ToEntity() v.Album {
	return v.Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}

func (entity Album) FromEntity(album v.Album) Album {
	return Album{
		ID:     album.ID,
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}
}
