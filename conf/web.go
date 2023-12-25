package conf

import "os"

type WebConf struct {
	Port           string
	WebProtoDomain string
}

func EnvWebConf() WebConf {
	webProtoDomain := os.Getenv("WEB_PROTO_DOMAIN")
	if webProtoDomain == `` {
		webProtoDomain = `http://localhost:8080`
	}

	return WebConf{
		Port:           os.Getenv("WEB_PORT"),
		WebProtoDomain: webProtoDomain,
	}
}

func (w WebConf) ListenAddr() string {
	return ":" + w.Port
}
