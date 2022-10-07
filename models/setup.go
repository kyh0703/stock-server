package models

import (
	"sync"

	"github.com/kyh0703/stock-server/ent"
)

var once sync.Once

func ConnectDatabase(database, uri string) error {
	var err error
	once.Do(func() {
	})
	return err
}

func CloseDatabase(client *ent.Client) {
	client.Close()
}
