package configs

import (
	"fmt"
	"os"

	"github.com/tkanos/gonfig"
)

type AppConfiguration struct {
	DSN                   string `env:"APP_DSN"`
	Port                  int    `env:"APP_PORT"`
	ConsoleLoggingEnabled bool
	FileLoggingEnabled    bool
	LogDirectory          string
	LogMaxAge             int
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
