package users

import (
	"context"

	"github.com/kyh0703/stock-server/database"
	"github.com/kyh0703/stock-server/ent"
	"github.com/kyh0703/stock-server/ent/users"
)

type UsersRepository struct{}

func (repo *UsersRepository) Save(ctx context.Context, username, email, password string) (*ent.Users, error) {
	return database.Ent.Users.
		Create().
		SetUsername(username).
		SetPassword(password).
		SetEmail(email).
		Save(ctx)
}

func (repo *UsersRepository) FetchOne(ctx context.Context, id int) (*ent.Users, error) {
	return database.Ent.Users.
		Query().
		Where(users.ID(id)).
		Only(ctx)
}

func (repo *UsersRepository) FetchByEmail(ctx context.Context, email string) (*ent.Users, error) {
	return database.Ent.Users.
		Query().
		Where(users.Email(email)).
		Only(ctx)
}

func (repo *UsersRepository) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	return database.Ent.Users.
		Query().
		Where(users.Email(email)).
		Exist(ctx)
}
