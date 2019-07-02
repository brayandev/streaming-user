package user

import (
	"context"
	"database/sql"
	"time"

	"go.uber.org/zap"
)

// Database methods.
type Database interface {
	ExecInsertItem(ctx context.Context, query RepositoryQuery, args ...interface{}) (sql.Result, error)
}

// LoggableDatabase database dependencies.
type LoggableDatabase struct {
	target Database
	logger *zap.Logger
}

// ExecInsertItem insert item on database.
func (d *LoggableDatabase) ExecInsertItem(ctx context.Context, query RepositoryQuery, args ...interface{}) (sql.Result, error) {
	start := time.Now()
	result, err := d.target.ExecInsertItem(ctx, query, args)

	lastID, lErr := result.LastInsertId()
	if lErr != nil {
		return result, lErr
	}

	d.logger.Info(
		"db query",
		zap.String("query", query.Name),
		zap.Duration("duration", time.Since(start)),
		zap.Any("args", args),
		zap.Int64("id", lastID),
		zap.NamedError("error", err),
	)

	return result, err
}

// MySQLDatabase is a connection with a MySQL database.
type MySQLDatabase struct {
	db            *sql.DB
	healthTimeout time.Duration
}

// NewMySQL constructor of database.
func NewMySQL(driverName, dataSourceName string, healthTimeout time.Duration, logger *zap.Logger) (*LoggableDatabase, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return &LoggableDatabase{&MySQLDatabase{db, healthTimeout}, logger}, nil
}

// ExecInsertItem insert item on database.
func (d *MySQLDatabase) ExecInsertItem(ctx context.Context, query RepositoryQuery, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query.Query, args)
}
