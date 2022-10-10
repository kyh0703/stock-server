package config

import (
	"context"

	"github.com/kyh0703/stock-server/ent"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(ctx context.Context) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}
	// auto migration tool
	if err = client.Schema.Create(ctx); err != nil {
		return nil, err
	}
	return client, nil
}
