package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const pgxDriverName = "pgx"

type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
	TimeZone string
}

func Sqlx(l *slog.Logger, cfg Config) (*sqlx.DB, error) {
	l.Info("Подклюечние БД...")

	connectionString := postgresConnectionString(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name, cfg.TimeZone)

	db, err := sqlx.Open(pgxDriverName, connectionString)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	l.Info("БД подключена!")

	return db, nil
}

func postgresConnectionString(host, port, user, password, name, timeZone string) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		host, port, user, password, name, timeZone)
}
