package configs

import (
	"io"
	"os"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
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
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
			zerolog.TimeFieldFormat = `2006/01/02 03:04 PM`
			var output io.Writer = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: `2006/01/02 03:04 PM`,
			}
			logOutput = output
		}

		buildInfo, _ := debug.ReadBuildInfo()

		l = zerolog.New(logOutput).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Caller().
			Str("go_version", buildInfo.GoVersion).
			Logger()
	})

	return &l
}
