package conf

import (
	"github.com/joho/godotenv"
)

func LoadEnv() {
	zlog := InitLogger()
	err := godotenv.Load(".env")
	if err != nil {
		zlog.Panic().
			Str("ERROR", err.Error()).
			Msg("cannot load .env files")
	}
}
