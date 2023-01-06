package users

import (
	"strconv"
	"time"

	"github.com/kyh0703/stock-server/pkg/cache"
)

type AuthRepository struct{}

func (repo *AuthRepository) FindOne(userID int) (int64, error) {
	return cache.Redis.Get(strconv.Itoa(userID)).Int64()
}

func (repo *AuthRepository) Save(userID int, expire time.Time) error {
	return cache.Redis.Set(
		strconv.Itoa(userID),
		userID,
		expire.Sub(time.Now())).
		Err()
}

func (repo *AuthRepository) Remove(userID string) (int64, error) {
	return cache.Redis.Del(userID).Result()
}
