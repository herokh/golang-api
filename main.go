package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	c "github.com/herokh/go-playground/configs"
	r "github.com/herokh/go-playground/routers"
)

func main() {
	config := c.GetConfiguration()
	logger := c.ConfigureLogger(config)
	logger.Info().Msg("Start application")
	db := c.NewMysql(config)
	db.AutoMigrate()

	gin := gin.Default()
	r.MapEndpoints(gin, logger, db, config)

	err := gin.Run(fmt.Sprintf("localhost:%s", strconv.Itoa(config.Port)))
	if err != nil {
		panic(err)
	}
}
