package configs

import (
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
)

var (
	once sync.Once
	l    zerolog.Logger
)

func InitLogger() *zerolog.Logger {
	once.Do(func() {
		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel)
		}

		var logOutput io.Writer
		if os.Getenv("WEB_ENV") == "prod" {
			file, _ := os.OpenFile("log/application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			logOutput = file
		} else {
			var output io.Writer = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: `03:04 PM`,
				PartsOrder: []string{
					zerolog.TimestampFieldName,
					zerolog.LevelFieldName,
					zerolog.CallerFieldName,
					zerolog.MessageFieldName,
				},
			}
			logOutput = output
		}

		l = zerolog.New(logOutput).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Caller().
			Logger()
	})

	return &l
}
