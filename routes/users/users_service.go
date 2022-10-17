package users

import (
	"context"
	"fmt"

	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/user"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct{}

func (svc *UsersService) CreateUser(ctx context.Context, name, email, password string) (*ent.User, error) {
	ok, err := svc.CheckUserExist(ctx, email)
	if err != nil || ok {
		return nil, err
	}
	user, err := svc.SaveUser(ctx, name, email, password)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (svc *UsersService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to generate hash: %w", err)
	}
	return string(hash), nil
}

func CompareHashPassword(hash, password string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			return false, fmt.Errorf("mismatch password: %w", err)
		}
		return false, err
	}
	return true, nil
}

func (svc *UsersService) SaveUser(ctx context.Context, name, email, password string) (*ent.User, error) {
	return database.Ent.User.
		Create().
		SetUsername(name).
		SetPassword(password).
		SetEmail(email).
		Save(ctx)
}

func (svc *UsersService) CheckUserExist(ctx context.Context, email string) (bool, error) {
	return database.Ent.User.
		Query().
		Where(user.Email(email)).
		Exist(ctx)
}

func (svc *UsersService) FindOne(ctx context.Context, id int) (*ent.User, error) {
	return database.Ent.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
}

func (svc *UsersService) FindByEmail(ctx context.Context, email string) (*ent.User, error) {
	return database.Ent.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}
