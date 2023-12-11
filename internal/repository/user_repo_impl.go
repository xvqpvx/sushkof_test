package repository

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"sushkof_test/internal/helper"
	"sushkof_test/internal/model"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) Save(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "INSERT INTO users(name, age, email, password) VALUES (?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, query, user.Name, user.Age, user.Email, user.Password)
	helper.PanicIfError(err)

	log.Info().Msgf("New user saved")
}

func (u *UserRepoImpl) GetIdByName(ctx context.Context, name string) int {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var id int

	query := "SELECT id_user FROM users WHERE name=?"
	err = tx.QueryRowContext(ctx, query, name).Scan(&id)
	helper.PanicIfError(err)

	return id
}

func (u *UserRepoImpl) FindAll(ctx context.Context) []model.User {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT id_user, name, age, email FROM users ORDER BY name ASC"
	result, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	defer result.Close()
	var users []model.User

	for result.Next() {
		user := model.User{}
		err = result.Scan(&user.IdUser, &user.Name, &user.Age, &user.Email)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}
