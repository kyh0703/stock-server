package config

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open(dialect.MySQL, "root:1234@tcp(localhost:3306)/stock?parseTime=true")
	if err != nil {
		return nil, err
	}
	// auto migration tool
	err = client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
