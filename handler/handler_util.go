package handler

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func encrypt(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateAccessToken() string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		fmt.Println("yo")
	}

	return token
}
