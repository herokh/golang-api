package configs

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	*zerolog.Logger
}

func ConfigureLogger(config *AppConfiguration) *Logger {
	var writers []io.Writer
	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}
	mw := io.MultiWriter(writers...)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := zerolog.New(mw).With().Timestamp().Logger()

	return &Logger{
		Logger: &logger,
	}
}

func newRollingFile(config *AppConfiguration) io.Writer {
	now := time.Now()
	return &lumberjack.Logger{
		Filename: path.Join(config.LogDirectory, fmt.Sprintf("%s.log", now.Format("20060102"))),
		MaxAge:   config.LogMaxAge,
	}
}
