package db_connection

import (
	"ZeroAgencyTest/config"
	"ZeroAgencyTest/lib/logger/sl"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log/slog"
)

func Connect(cfg config.DB, log *slog.Logger) (*reform.DB, error) {
	urlExample := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SslMode)

	sqlDb, err := sql.Open(cfg.DriverName, urlExample)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDb.SetMaxOpenConns(cfg.MaxConns)
	sqlDb.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	db := reform.NewDB(
		sqlDb,
		postgresql.Dialect,
		sl.NewAdapterLogger(log),
	)

	if err = sqlDb.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}
