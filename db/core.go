package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Code-Hex/sqlx-transactionmanager"
)

func OpenDB(driver, dsn string) *sqlx.DB {
	db := sqlx.MustOpen(driver, dsn)

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(30)

	return db;
}

func (c *Core) Ping(timeout time.Duration) (err error) {
	ctx, cancel := context.WithTimeout(*c.Ctx, timeout)
	defer cancel()

	err = c.DB.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("ping db failed: %v", err)
	}

	return
}

func (c *Core) Do(opts *JobOptions) {
	ctx, cancel := context.WithTimeout(*c.Ctx, opts.Timeout)
	defer cancel()

	txm, err := c.DB.BeginTxmx(ctx, opts.TxOpts)
	if err != nil {
		opts.Fail(err)
		return
	}

	defer func() {
		if err := recover(); err != nil {
			txm.Rollback()
			opts.Fail(fmt.Errorf("%v", err))
		}
	}()

	opts.Job(&Core{DB: c.DB, Txm: txm, Ctx: &ctx})

	txm.Commit()
}

func (c *Core) DoSimple(opts *JobOptions) {
	ctx, cancel := context.WithTimeout(*c.Ctx, opts.Timeout)
	defer cancel()

	defer func() {
		if err := recover(); err != nil {
			opts.Fail(fmt.Errorf("%v", err))
		}
	}()

	opts.Job(&Core{DB: c.DB, Ctx: &ctx})
}

func (c *Core) MustExec(sql string, args... interface{}) sql.Result {
	if c.Txm != nil {
		return c.Txm.MustExecContext(*c.Ctx, sql, args...)
	} else {
		return c.DB.MustExecContext(*c.Ctx, sql, args...)
	}
}

func NewJobOpts(job Job, fail Fail) *JobOptions {
	return &JobOptions{ Timeout: 5 * time.Second, TxOpts: &sql.TxOptions{}, Job: job, Fail: fail }
}
