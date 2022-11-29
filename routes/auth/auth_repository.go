package auth

import (
	"strconv"
	"time"

	"github.com/kyh0703/stock-server/database"
)

type AuthRepository struct{}

func (repo *AuthRepository) Fetch(userID int) (int64, error) {
	return database.Redis.Get(strconv.Itoa(userID)).Int64()
}

func (repo *AuthRepository) InsertToken(userID int, expire time.Time) error {
	return database.Redis.Set(
		strconv.Itoa(userID),
		userID,
		expire.Sub(time.Now())).
		Err()
}

func (repo *AuthRepository) Delete(userID string) (int64, error) {
	return database.Redis.Del(userID).Result()
}
