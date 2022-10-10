package models

import (
	"context"

	"github.com/kyh0703/stock-server/config"
	"github.com/kyh0703/stock-server/ent"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ent.User
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) SaveUser(ctx context.Context) (*User, error) {
	user, err := config.DBClient().User.Create().
		SetNickname(u.Nickname).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *User) DeleteUser(ctx context.Context) {
	config.DBClient().User.Delete(u)
})