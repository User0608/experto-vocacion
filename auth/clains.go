package auth

import "github.com/dgrijalva/jwt-go"

type Clain struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
