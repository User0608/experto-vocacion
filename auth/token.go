package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/user0608/expertos/models"
)

var (
	ErrInvalidToken error = errors.New("token no válido")
	ErrInvalidClams error = errors.New("No se pudo obtener la data")
)

//GenerageToken Genera un toquen para un usuario o cliente, el cual se empleara para futuras peticiones
// func GenerageToken(customer model.Customer) (string, error) {
func GenerageTokenAdmin(usuario models.Usuario) (string, error) {
	claim := Clain{
		UserName: usuario.Username,
		Role:     usuario.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(singKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

//ValidateToken verifica la firma de un token para comprobar su valides
func ValidateToken(t string) (Clain, error) {
	token, err := jwt.ParseWithClaims(t, &Clain{}, verifyFunction)
	if err != nil {
		return Clain{}, err
	}
	if !token.Valid {
		return Clain{}, ErrInvalidToken
	}
	claim, ok := token.Claims.(*Clain)
	if !ok {
		return Clain{}, ErrInvalidClams
	}
	return *claim, nil
}
func verifyFunction(token *jwt.Token) (interface{}, error) {
	return verifyKey, nil
}
