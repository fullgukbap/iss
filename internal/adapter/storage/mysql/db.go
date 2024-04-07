package mysql

import (
	"context"

	"fmt"
	"letsgo-mini-is/internal/adapter/config"
	"letsgo-mini-is/internal/adapter/storage/mysql/ent"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*ent.Client
}

func New(ctx context.Context, config *config.DB) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return &DB{Client: client}, nil
}

func (d *DB) Close() error {
	return d.Close()
}
