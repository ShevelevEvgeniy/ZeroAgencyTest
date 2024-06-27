package sl

import (
	"log/slog"
	"time"
)

type AdapterLogger struct {
	logger *slog.Logger
}

func NewAdapterLogger(logger *slog.Logger) *AdapterLogger {
	return &AdapterLogger{logger: logger}
}

func (al *AdapterLogger) Before(query string, args []interface{}) {
	al.logger.Info("Executing query", "query", query, "args", args)
}

func (al *AdapterLogger) After(query string, args []interface{}, d time.Duration, err error) {
	al.logger.Info("Executed query", "query", query, "args", args, "duration", d, "error", err)
}
