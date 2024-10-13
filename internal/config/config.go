package config

type Environment struct {
	LogLevel string `env:"LOG_LEVEL" required:"false" default:"debug"`

	HttpPort int    `env:"HTTP_PORT" required:"true" default:"8080"`
	HttpHost string `env:"HTTP_HOST" required:"true" default:"localhost"`

	PGDB
	RedisDB
}

type PGDB struct {
	PGDBHost string `env:"PG_DB_SERVER" required:"true" default:"192.168.1.242"`
	Port     int    `env:"PG_DB_PORT" required:"true"  default:"5432"`
	Name     string `env:"PG_DB_NAME" required:"true"  default:"editor"`
	User     string `env:"PG_DB_USER" required:"true"  default:"orc"`
	Password string `env:"PG_DB_PASS" required:"true"  default:"pAEiYSOu04"`
	SslMode  string `env:"PG_DB_SSLMODE" required:"true"  default:"disabled"`
}

type RedisDB struct {
	Host     string `env:"REDIS_SERVER" required:"true" default:"192.168.1.242"`
	Port     string `env:"REDIS_HOST" required:"true" default:"6380"`
	Password string `env:"REDIS_PASSWORD" required:"true" default:"qwerty"`
}
