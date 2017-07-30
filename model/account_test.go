package model_test

import (
	"testing"
	"go-auth-sample/model"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/lib/pq"
)

func TestRegister(t *testing.T) {
	var account = model.Account{}
	account.Username = "testuser6"
	account.Email = "test@test6.com"
	account.Password = "secret"
	db := initializeDB()
	user, err := account.Register(db, account)
	fmt.Println("User value : ", user)
	assert.NotNil(t, user)
	assert.Nil(t, err)

}


// TODO: Fix GetUserWithId for no values returned.

func TestAccount_GetUserWithId(t *testing.T) {
	var account = model.Account{}
	db := initializeDB()
	userId := 1100
	user, err := account.GetUserWithId(db, userId)
	fmt.Println("User value : ", user)
	assert.NotEmpty(t, user)
	assert.Nil(t, err)
}

func initializeDB() *sql.DB {
	const (
		host        = "localhost"
		port        = "5432"
		user    = "admin"
		password    = "C0mplexPwd!1234"
		dbname      = "goauth"
		search_path = "auth"
	)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s "+
		"dbname=%s search_path=%s sslmode=disable", host, port, user, password, dbname, search_path)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	return db
}
