package services

import (
	"github.com/herokh/go-playground/configs"
	"github.com/herokh/go-playground/models"
)

type AuthService interface {
	Login(email string, password string) string
}

type authService struct {
	Logger *configs.Logger
	Db     *configs.Database
}

func NewAuthService(logger *configs.Logger, db *configs.Database) AuthService {
	return &authService{
		Logger: logger,
		Db:     db,
	}
}

func (service *authService) Login(email string, password string) string {
	var model models.User
	service.Db.DbContext.Where("email = ? AND password = ?", email, password).Find(&model)
	if model.ID == 0 {
		return ""
	}
	return "TEST_TOKEN"
}
