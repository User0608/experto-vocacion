package auth

import (
	"log"

	"github.com/labstack/echo/v4"
)

const (
	USERNAME_KEY = "username"
	ROLE_KEY     = "role"
)

func itemIsContainIn(item string, values []string) bool {
	for _, val := range values {
		if item == val {
			return true
		}
	}
	return false
}
func jwtValidToken(token string, c echo.Context) error {
	if token == "" {
		return echo.ErrForbidden
	}
	clain, err := ValidateToken(token)
	if err != nil {
		log.Println("El token no es valido")
		return echo.ErrForbidden
	}
	c.Set(USERNAME_KEY, clain.UserName)
	c.Set(ROLE_KEY, clain.Role)
	return nil
}

func JWTMiddleware(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if err := jwtValidToken(token, c); err != nil {
			return echo.ErrForbidden
		}
		return f(c)
	}
}
func RolesMiddleware(f echo.HandlerFunc, roles ...string) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get(ROLE_KEY).(string)
		if !ok {
			return echo.ErrForbidden
		}
		for _, r := range roles {
			if role == r {
				return f(c)
			}
		}
		return echo.ErrForbidden
	}
}
