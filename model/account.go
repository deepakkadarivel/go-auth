package model

import (
	"database/sql"
	"errors"
	"go-auth-sample/util"
	"fmt"
	"go-auth-sample/jwthandler"
	"github.com/lib/pq"
)

type Account struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func (acc *Account) login(db *sql.DB) (User, error) {
	return User{}, errors.New("Not Implemented.")
}

func (acc *Account) Register(db *sql.DB, account Account) (interface{}, error) {
	password_hash, err := util.HashPassword(account.Password)
	if err != nil {
		return nil, err
	}

	queryStmt, err := db.Prepare("INSERT INTO account(username, email, password_hash) VALUES($1, $2, $3) RETURNING user_id")

	if err != nil {
		return nil, err
	}

	var lastInsertId int
	err = queryStmt.QueryRow(account.Username, account.Email, password_hash).Scan(&lastInsertId)

	if err != nil {
		pqErr := err.(*pq.Error)
		if pqErr.Code.Name() == "unique_violation" {
			return nil, fmt.Errorf("User with username/email already exixts.")
		} else {
			return nil, err
		}
	}

	fmt.Println("last inserted id =", lastInsertId)

	user, err := acc.GetUserWithId(db, lastInsertId)
	if err != nil {
		return nil, err
	}

	fmt.Println("User : ", user)
	return user, nil
}

func (acc *Account) resetPassword(db *sql.DB) (User, error) {
	return User{}, errors.New("Not Implemented.")
}

func (acc *Account) GetUserWithId(db *sql.DB, id int) (interface{}, error) {
	var user = User{}
	queryStmt, err := db.Prepare("SELECT username, email FROM account WHERE user_id=$1")
	if err != nil {
		return nil, err
	}

	err = queryStmt.QueryRow(id).Scan(&user.Username, &user.Email)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("No Results Found")
	}
	if err != nil {
		return nil, err
	}

	user.Token = jwthandler.GetAuthToken(user.Username, user.Email)
	fmt.Println("User : ", user)
	return user, nil
}