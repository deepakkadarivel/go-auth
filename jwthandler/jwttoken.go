package jwthandler

import (
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
	"encoding/json"
	"github.com/auth0/go-jwt-middleware"
)

var MySigningKey = []byte("secret")

type accessToken struct {
	Token string      `json:"token"`
}

var GetToken = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["admin"] = true
	claims["name"] = "testing"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, _ := token.SignedString(MySigningKey)

	tokenObject := &accessToken{
		Token: tokenString,
	}
	payload, _ := json.Marshal(tokenObject)

	res.Header().Set("Content-Type", "application/json")
	res.Write([]byte(payload))
})

var JwtValidator = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return MySigningKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
