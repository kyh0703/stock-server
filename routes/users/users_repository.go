package users

import (
	"context"

	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/user"
)

type UsersRepository struct{}

func (repo *UsersRepository) Insert(ctx context.Context, username, email, password string) (*ent.User, error) {
	return database.Ent.User.
		Create().
		SetUsername(username).
		SetPassword(password).
		SetEmail(email).
		Save(ctx)
}

func (repo *UsersRepository) FetchOne(ctx context.Context, id int) (*ent.User, error) {
	return database.Ent.User.
		Query().
		Where(user.ID(id)).
		Only(ctx)
}

func (repo *UsersRepository) FetchByEmail(ctx context.Context, email string) (*ent.User, error) {
	return database.Ent.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)
}

func (repo *UsersRepository) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	return database.Ent.User.
		Query().
		Where(user.Email(email)).
		Exist(ctx)
}
