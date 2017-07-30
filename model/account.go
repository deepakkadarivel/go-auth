package model

import (
	"database/sql"
	"errors"
	"go-auth-sample/util"
	"fmt"
	"go-auth-sample/jwthandler"
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

func (acc *Account) Register(db *sql.DB, account Account) (User, error) {
	password_hash, err := util.HashPassword(account.Password)

	if err != nil {
		return User{}, err
	}

	var lastInsertId int
	err = db.QueryRow(
		"INSERT INTO account(username, email, password_hash) VALUES($1, $2, $3) RETURNING user_id",
		account.Username, account.Email, password_hash).Scan(&lastInsertId)

	if err != nil {
		return User{}, err
	}

	fmt.Println("last inserted id =", lastInsertId)

	user, err := acc.GetUserWithId(db, lastInsertId)
	if err != nil {
		return User{}, err
	}

	fmt.Println("User : ", user)
	return user, nil
}

func (acc *Account) resetPassword(db *sql.DB) (User, error) {
	return User{}, errors.New("Not Implemented.")
}

func (acc *Account) GetUserWithId(db *sql.DB, id int) (User, error)  {
	rows, err := db.Query("SELECT username, email FROM account WHERE user_id=$1", id)
	if err != nil {
		return User{}, err
	}

	var user = User{}
	for rows.Next() {
		err = rows.Scan(&user.Username, &user.Email)
		if err != nil {
			return User{}, err
		}
		user.Token = jwthandler.GetAuthToken(user.Username, user.Email)
		fmt.Println("Username | Email | Token")
		fmt.Printf("%v | %v | %v", user.Username, user.Email, user.Token)
	}

	return user, nil
}
