package database

import (
	"context"

	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var Ent *ent.Client

func ConnectDatabase(ctx context.Context, databaseType, databaseHost string) (*ent.Client, error) {
	var err error

	// create ent client
	Ent, err = ent.Open(databaseType, databaseHost)
	if err != nil {
		return nil, err
	}

	// set debug mode
	Ent = Ent.Debug()

	// auto migration tool
	err = Ent.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
	)
	if err != nil {
		return nil, err
	}

	return Ent, nil
}
