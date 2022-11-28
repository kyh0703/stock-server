package auth

import (
	"time"

	"github.com/kyh0703/stock-server/database"
)

type AuthRepository struct{}

func (repo *AuthRepository) FetchUserIdByUUID(uuid string) (int, error) {
	return database.Redis.Get(uuid).Int()
}

func (repo *AuthRepository) InsertToken(userID int, uuid string, expire, now time.Time) error {
	return database.Redis.Set(
		uuid,
		userID,
		expire.Sub(now)).
		Err()
}

func (repo *AuthRepository) Delete(id string) (int64, error) {
	return database.Redis.Del(id).Result()
}
