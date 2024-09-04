package config

type Environment struct {
	LogLevel string `env:"LOG_LEVEL" required:"false" default:"debug"`

	HttpPort int    `env:"HTTP_PORT" required:"true" default:"8080"`
	HttpHost string `env:"HTTP_HOST" required:"true" default:"localhost"`
}
