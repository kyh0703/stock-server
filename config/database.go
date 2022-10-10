package config

import (
	"context"
	"sync"

	"entgo.io/ent/entc/integration/idtype/ent"
)

var (
	instance *ent.Client
	once     sync.Once
)

func ConnectDatabase(ctx context.Context) error {
	var err error
	once.Do(func() {
		instance, err = ent.Open("sqllite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			return
		}
		// auto migration tool
		if err = instance.Schema.Create(ctx); err != nil {
			return
		}
	})
	return err
}

func CloseDatabase() error {
	if instance != nil {
		instance.Close()
	}
}

func DBClient() *ent.Client {
	return instance
}
