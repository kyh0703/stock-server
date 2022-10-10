package models

import (
	"context"

	"entgo.io/ent/entc/integration/idtype/ent/user"
	"github.com/kyh0703/stock-server/ent"
	"golang.org/x/crypto/bcrypt"
)

type User ent.User

func (u *User) SetPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hashedPassword, err
	}
	u.SetPassword(string(hashedPassword))
	return hashedPassword, nil
}

func (u *User) CheckPassword(hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		return false
	}
	return true
}

func (u *User) SaveUser(ctx context.Context) (*User, error) {
	user, err := client.User.Create().
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
	client.User.DeleteOne(u)
}

func (u *User) FindByUserName(ctx context.Context, username string) ([]*User, error) {
	users, err := client.User.Query().Where(user.NameContains(username)).All(ctx)
	if err != nil {
		return nil, err
	}
	return users, err
}
