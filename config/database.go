package config

import "time"

type DB struct {
	Host            string        `envconfig:"DB_HOST" env-required:"true"`
	Port            string        `envconfig:"DB_PORT" env-required:"true"`
	DBName          string        `envconfig:"DB_NAME" env-required:"true"`
	Username        string        `envconfig:"DB_USER_NAME" env-required:"true"`
	Password        string        `envconfig:"DB_PASSWORD" env-required:"true"`
	SslMode         string        `envconfig:"DB_SSL_MODE" env-default:"disable"`
	DriverName      string        `envconfig:"DB_DRIVER_NAME" env-default:"postgres"`
	MigrationUrl    string        `envconfig:"MIGRATION_URL" env-required:"true"`
	MaxConns        int           `envconfig:"DB_MAX_CONNS" env-default:"10"`
	MaxIdleConns    int           `envconfig:"DB_MAX_IDEL_CONNS" env-default:"5"`
	ConnMaxLifetime time.Duration `envconfig:"DB_CONNS_MAX_LIFE_TIME" env-default:"5"`
}
