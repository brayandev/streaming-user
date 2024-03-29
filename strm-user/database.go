package user

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql" //Driver for MySQL.

	"go.uber.org/zap"
)

// Database methods.
type Database interface {
	ExecInsertItem(ctx context.Context, query RepositoryQuery, user User) (sql.Result, error)
	QuerySingleResult(ctx context.Context, scanner DatabaseMapperFn, query RepositoryQuery, args ...interface{}) (interface{}, error)
}

// LoggableDatabase database dependencies.
type LoggableDatabase struct {
	target Database
	logger *zap.Logger
}

// ExecInsertItem insert item on database.
func (d *LoggableDatabase) ExecInsertItem(ctx context.Context, query RepositoryQuery, user User) (sql.Result, error) {
	start := time.Now()
	result, err := d.target.ExecInsertItem(ctx, query, user)

	lastID, lErr := result.LastInsertId()
	user.ID = lastID

	d.logger.Info(
		"db query",
		zap.String("query", query.Name),
		zap.Duration("duration", time.Since(start)),
		zap.Any("user", user),
		zap.NamedError("error", err),
		zap.NamedError("error-last-id", lErr),
	)

	return result, err
}

// QuerySingleResult insert item on database.
func (d *LoggableDatabase) QuerySingleResult(ctx context.Context, scanner DatabaseMapperFn, query RepositoryQuery, args ...interface{}) (interface{}, error) {
	start := time.Now()
	result, err := d.target.QuerySingleResult(ctx, scanner, query, args)

	d.logger.Info(
		"db query",
		zap.String("query", query.Name),
		zap.Duration("duration", time.Since(start)),
		zap.Any("args", args),
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
func (d *MySQLDatabase) ExecInsertItem(ctx context.Context, query RepositoryQuery, user User) (sql.Result, error) {
	stm, err := d.db.PrepareContext(ctx, query.Query)
	if err != nil {
		return nil, err
	}

	return stm.ExecContext(ctx, user.Name, user.Email)
}

// RowScanner used to accept a sql.Row.
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// DatabaseMapperFn scanner.
type DatabaseMapperFn func(row RowScanner) (interface{}, error)

// QuerySingleResult get one result on database.
func (d *MySQLDatabase) QuerySingleResult(ctx context.Context, scanner DatabaseMapperFn, query RepositoryQuery, args ...interface{}) (interface{}, error) {
	return d.db.QueryContext(ctx, query.Query, args)
}
