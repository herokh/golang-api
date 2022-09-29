package configs

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

type AppConfiguration struct {
	// application
	Port int `env:"APP_PORT"`

	// database
	DSN string `env:"DB_DSN"`

	// logging
	ConsoleLoggingEnabled bool
	FileLoggingEnabled    bool
	LogDirectory          string
	LogMaxAge             int

	// auth
	AuthSecret string `env:"AUTH_SECRET"`
}

func GetConfiguration() *AppConfiguration {
	var env string
	if env = os.Getenv("ENV"); env == "" {
		env = "dev"
	}

	configuration := &AppConfiguration{}
	filename := fmt.Sprintf("./%s_config.json", env)
	err := gonfig.GetConf(filename, configuration)
	if err != nil {
		panic(err)
	}
	return configuration
}
