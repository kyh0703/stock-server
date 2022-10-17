package database

import (
	"context"

	"entgo.io/ent/dialect"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var Ent *ent.Client

func ConnectDb(ctx context.Context) (*ent.Client, error) {
	var err error
	// create ent client
	Ent, err = ent.Open(
		dialect.MySQL,
		"root:dnflth@tcp(localhost:3306)/stock?parseTime=true",
	)
	if err != nil {
		return nil, err
	}
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
