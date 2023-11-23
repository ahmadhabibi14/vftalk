package middlewares

import (
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var LoggerConfig = logger.Config{
	Format:     "${time} | ${status} | ${latency} | ${method} | ${path}\n",
	TimeFormat: "2006/01/02 03:04 PM",
	TimeZone:   "Asia/Makassar",
}
