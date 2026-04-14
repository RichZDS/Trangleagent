//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	account := os.Getenv("ACCOUNT")
	if account == "" {
		account = "DuBuque1671714"
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": account,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("TrangleAgent"))
	fmt.Println(tokenString)
}
