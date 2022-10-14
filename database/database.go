package database

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	ec *ent.Client
	eo sync.Once
)

func Ent() *ent.Client {
	return ec
}

func ConnectDb(ctx context.Context) (*ent.Client, error) {
	var err error
	eo.Do(func() {
		// creaete ent client
		ec, err = ent.Open(
			dialect.MySQL,
			"root:1234@tcp(localhost:3306)/stock?parseTime=true",
		)
		if err != nil {
			return
		}
		// auto migration tool
		err = ec.Schema.Create(
			ctx,
			migrate.WithDropColumn(true),
			migrate.WithDropIndex(true),
		)
		if err != nil {
			return
		}
	})
	return ec, err
}
