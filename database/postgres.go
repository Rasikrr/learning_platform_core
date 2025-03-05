package database

import (
	"context"
	"github.com/Rasikrr/learning_platform_core/configs"
	"github.com/Rasikrr/learning_platform_core/configs/appenv"
	"github.com/avast/retry-go"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Postgres struct {
	pool *pgxpool.Pool
}

// nolint: gosec
func NewPostgres(ctx context.Context, cfg *configs.Config) (*Postgres, error) {
	conConfig, err := pgxpool.ParseConfig(cfg.Env.Get(appenv.PostgresDSN).GetString())
	if err != nil {
		return nil, err
	}
	conConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	conConfig.MaxConns = int32(cfg.Postgres.MaxConns)
	conConfig.MinConns = int32(cfg.Postgres.MinConns)
	conConfig.MaxConnIdleTime = cfg.Postgres.MaxIdleConnIdleTime

	pool, err := pgxpool.NewWithConfig(ctx, conConfig)
	if err != nil {
		return nil, err
	}

	err = retry.Do(func() error {
		return pool.Ping(ctx)
	},
		retry.Attempts(3),
		retry.Delay(1*time.Second),
		retry.OnRetry(func(_ uint, err error) {
			log.Printf("Failed to connect to database: %v\n", err)
		}),
	)

	if err != nil {
		return nil, err
	}
	return &Postgres{
		pool: pool,
	}, nil
}

func (p *Postgres) Pool() *pgxpool.Pool {
	return p.pool
}

func (p *Postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Query: %s; elapsed: %v; args: %v\n", sql, elapsed, args)
	}()
	return p.pool.Query(ctx, sql, args...)
}

func (p *Postgres) Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Exec: %s; elapsed: %v; args: %v\n", sql, elapsed, args)
	}()
	return p.pool.Exec(ctx, sql, args...)
}

func (p *Postgres) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("QueryRow: %s; elapsed: %v; args: %v\n", sql, elapsed, args)
	}()
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *Postgres) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("BeginTx: %v; elapsed: %v\n", txOptions, elapsed)
	}()
	return p.pool.BeginTx(ctx, txOptions)
}

func (p *Postgres) Begin(ctx context.Context) (pgx.Tx, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("Begin: %v; elapsed: %v\n", nil, elapsed)
	}()
	return p.pool.Begin(ctx)
}

func (p *Postgres) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("CopyFrom: %s; elapsed: %v; args: %v\n", tableName, elapsed, columnNames)
	}()
	return p.pool.CopyFrom(ctx, tableName, columnNames, rowSrc)
}

func (p *Postgres) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("SendBatch: %v; elapsed: %v\n", b, elapsed)
	}()
	return p.pool.SendBatch(ctx, b)
}

func (p *Postgres) Close(_ context.Context) error {
	p.pool.Close()
	log.Println("Postgres closed gracefully")
	return nil
}
