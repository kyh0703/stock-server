package config

import (
	"context"

	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDatabase(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("mysql", "root:1234@tcp(localhost:3306)/mydb")
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
