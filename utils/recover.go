package utils

import "github.com/rs/zerolog"

func Recover(zlog *zerolog.Logger) {
	if rec := recover(); rec != nil {
		zlog.Warn().Str("recover", rec.(string))
	}
}
