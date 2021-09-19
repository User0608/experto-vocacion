package models

import (
	"fmt"
	"time"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/kcheck"
)

type Estudiante struct {
	ID              int       `json:"estudiante_id"`
	Nombre          string    `json:"nombre"`
	ApellidoPaterno string    `json:"apellido_paterno"`
	ApellidoMaterno string    `json:"apellido_materno"`
	FechaNacimiento time.Time `json:"fecha_nacimiento"`
	Dni             string    `json:"dni"`
	Password        string    `json:"password,omitempty"`
}

func (e Estudiante) Validate() error {
	chk := kcheck.New()
	if err := chk.Target("min=2 basic", e.Nombre, e.ApellidoPaterno).Ok(); err != nil {
		return errs.ErrInvalidData
	}
	if err := chk.Target("len=8", e.Dni).Ok(); err != nil {
		return fmt.Errorf("DNI Invalida")
	}
	if err := chk.Target("min=6", e.Password).Ok(); err != nil {
		return fmt.Errorf("La contraseña debe de contener al menos 6 caracteres")
	}
	return nil
}

// func main() {
// t, err := time.Parse(time.RFC3339, "2021-09-12T16:34:21.763Z")
// if err != nil {
// 	fmt.Println("Error", err.Error())
// 	return
// }

// 	fmt.Println(t)
// 	fmt.Println(time.Now())
// }
