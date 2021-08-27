package postgres

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type (
	PostgresDriver struct {
		Client *sql.DB
	}
)

func (driver *PostgresDriver) Exec(ctx context.Context, sql string) (sql.Result, error) {
	return driver.Client.ExecContext(ctx, sql)
}

func (driver *PostgresDriver) Query(ctx context.Context, sql string) (*sql.Rows, error) {
	return driver.Client.QueryContext(ctx, sql)
}
