package models

import (
	"context"
	"sync"

	"entgo.io/ent/entc/integration/idtype/ent"
)

var (
	client *ent.Client
	once   sync.Once
)

func Connect(ctx context.Context) error {
	var err error
	once.Do(func() {
		client, err = ent.Open("sqllite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			return
		}
		// auto migration tool
		if err = client.Schema.Create(ctx); err != nil {
			return
		}
	})
	return err
}

func Close() error {
	if client != nil {
		client.Close()
	}
}
