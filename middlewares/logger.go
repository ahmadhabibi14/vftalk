package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiberLogger() logger.Config {
	var conf logger.Config

	if os.Getenv("WEB_ENV") == `prod` {
		file, _ := os.OpenFile("log/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		conf.Format = "{\"time\": \"${time}\", \"status\": \"${status}\", \"ip\": \"${ips}\", \"latency\": \"${latency}\", \"method\": \"${method}\", \"path\": \"${path}\", \"body\": '${body}'}\n"
		conf.TimeFormat = "2006-01-02T03:00:55+08:00"
		conf.TimeZone = "Asia/Makassar"
		conf.Output = file
	} else {
		conf.Format = "${time} | ${status} | ${latency} | ${method} | ${path}\n"
		conf.TimeFormat = "2006/01/02 03:04 PM"
		conf.TimeZone = "Asia/Makassar"
		conf.Output = os.Stdout
	}

	return conf
}
