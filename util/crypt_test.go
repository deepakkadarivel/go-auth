package util_test

import (
	"testing"
	"go-auth-sample/util"
	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	var password = "secret"
	hash, err := util.HashPassword(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)

}

func TestCheckPasswordForValidHashAndPasswordMatchSucceeds(t *testing.T) {
	var password = "secret"
	hash, err := util.HashPassword(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	isValid := util.CheckPassword(password, hash)
	assert.True(t, isValid)
}

func TestCheckPasswordForInValidHashAndPasswordMatchFails(t *testing.T) {
	var password = "secret"
	hash, err := util.HashPassword(password)
	assert.Nil(t, err)
	assert.NotNil(t, hash)
	password = "notsecret"
	isValid := util.CheckPassword(password, hash)
	assert.False(t, isValid)
}
