package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	c "github.com/herokh/golang-api/configs"
	"github.com/herokh/golang-api/dtos"
	"github.com/herokh/golang-api/services"
)

type AuthController interface {
	Login(c *gin.Context)
}

type authController struct {
	Logger *c.Logger
	Db     *c.Database
	Config *c.AppConfiguration
}

func NewAuthController(logger *c.Logger, db *c.Database, config *c.AppConfiguration) AuthController {
	return &authController{
		Logger: logger,
		Db:     db,
		Config: config,
	}
}

func (controller *authController) Login(c *gin.Context) {
	var dto dtos.LoginCredentials
	err := c.ShouldBind(&dto)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	service := services.NewAuthService(controller.Logger, controller.Db, controller.Config)
	accessToken := service.Login(dto.Email, dto.Password)
	if accessToken == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
