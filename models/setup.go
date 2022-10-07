package models

import (
	"sync"

	"github.com/kyh0703/stock-server/ent"
	"gorm.io/gorm"
)

var (
	Client *ent.Client
	once   sync.Once
)

func DataBase()
func ConnectDatabase(database, uri string) error {
	var err error
	once.Do(func() {
		Client, err = gorm.Open(database)
		if err != nil {
			return
		}
	})
	return err
}

func CloseDatabase(client *ent.Client) {
	client.Close()
}
