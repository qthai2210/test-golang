package repo_impl

import (
	"backend-github/db"
	"backend-github/errorlist"
	"backend-github/model"
	"backend-github/model/req"
	"backend-github/repository"
	"context"
	"database/sql"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql}
}

func (u *UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `INSERT INTO users (user_id,email, password, role,full_name, created_at, updated_at)
	VALUES (:user_id,:email, :password, :role,:full_name, :created_at, :updated_at)`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errorlist.UserConflict
			}
		}
		return user, errorlist.SignUpfail
	}

	return user, nil
}

func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email = $1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errorlist.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
func (u *UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error) {

	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE user_id = $1", userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errorlist.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
func (u *UserRepoImpl) UpdateUser(context context.Context, user model.User) (model.User, error) {
	statement := `UPDATE users 
	SET 
	    full_name = (CASE WHEN LENGTH(:full_name) > 0 THEN :full_name ELSE full_name END),
	    email = (CASE WHEN LENGTH(:email) > 0 THEN :email ELSE email END),
	    updated_at = COALESCE(:updated_at, updated_at)
	WHERE user_id = :user_id`
	user.UpdatedAt = time.Now()
	result, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		return user, errorlist.UserUpdateFail
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Error(err.Error())
		return user, errorlist.UserUpdateFail
	}
	if count == 0 {
		return user, errorlist.UserNotUpdated
	}
	return user, nil
}
