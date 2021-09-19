package models

import (
	"fmt"
	"strings"
)

type Usuario struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Estado   bool   `json:"estado" gorm:"->"`
}

func (u *Usuario) valString(s string, name string) error {
	if strings.Contains(s, " ") {
		return fmt.Errorf("El campo %s no puede contener espacios", name)
	}
	if len(s) < 6 {
		return fmt.Errorf("El campo %s no pude contener menos de 6 caracteres", name)
	}
	return nil
}
func (u *Usuario) Validate() error {
	if err := u.valString(u.Username, "username"); err != nil {
		return err
	}
	if err := u.valString(u.Password, "password"); err != nil {
		return err
	}
	return nil
}
