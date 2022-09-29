package services

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/herokh/golang-api/configs"
	"github.com/herokh/golang-api/models"
)

type AuthService interface {
	Login(email string, password string) string
}

type authService struct {
	Logger *configs.Logger
	Db     *configs.Database
	Config *configs.AppConfiguration
}

func NewAuthService(logger *configs.Logger, db *configs.Database, config *configs.AppConfiguration) AuthService {
	return &authService{
		Logger: logger,
		Db:     db,
		Config: config,
	}
}

func (service *authService) Login(email string, password string) string {
	var model models.User
	service.Db.DbContext.Where("email = ? AND password = ?", email, password).Find(&model)
	if model.ID == 0 {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        strconv.Itoa(int(model.ID)),
		Subject:   email,
		Issuer:    "golang-api",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
	})

	secret := []byte(service.Config.AuthSecret)
	signedToken, _ := token.SignedString(secret)

	return signedToken
}
