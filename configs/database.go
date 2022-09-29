package configs

import (
	m "github.com/herokh/go-playground/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DbContext *gorm.DB
}

func (db *Database) AutoMigrate() {
	db.DbContext.AutoMigrate(
		&m.Album{},
		&m.User{})
}

func NewMysql(config *AppConfiguration) *Database {
	var db *gorm.DB
	var err error
	if db, err = gorm.Open(mysql.Open(config.DSN), &gorm.Config{}); err != nil {
		panic(err)
	}

	return &Database{
		DbContext: db,
	}
}
