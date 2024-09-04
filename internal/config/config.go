package config

type Environment struct {
	LogLevel string `env:"LOG_LEVEL" required:"false" default:"debug"`

	HttpPort int    `env:"HTTP_PORT" required:"true" default:"8080"`
	HttpHost string `env:"HTTP_HOST" required:"true" default:"localhost"`

	PGDB
}

type PGDB struct {
	PGDBHost     string `env:"PG_DB_SERVER" required:"true" default:"192.168.1.242"`
	PGDBPort     int    `env:"PG_DB_PORT" required:"true"  default:"5432"`
	PGDBName     string `env:"PG_DB_NAME" required:"true"  default:"editor"`
	PGDBUser     string `env:"PG_DB_USER" required:"true"  default:"orc"`
	PGDBPassword string `env:"PG_DB_PASS" required:"true"  default:"pAEiYSOu04"`
	PGDBSslMode  string `env:"PG_DB_SSLMODE" required:"true"  default:"disabled"`
}
