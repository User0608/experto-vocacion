package errs

import "errors"

var (
	ErrInternal           = errors.New("Hubo problemas al procesar la peticion!")
	ErrServiceResponse    = errors.New("Hubo un problema interno!")
	ErrApiConnection      = errors.New("No se pudo connectar con el servicio!")
	ErrDataTypeOrStruct   = errors.New("No se pudo procesar la data enviada!")
	ErrDataBaseError      = errors.New("No se pudo completar la operacion!")
	ErrUpdateDNI          = errors.New("Usted no puede actualizar su DNI!")
	ErrDNIExist           = errors.New("Ya existe un usuario registrado con el mismo numero de DNI!")
	ErrUsernameOrPassword = errors.New("Usuario o contraseña incorrectas!")
	ErrInvalidData        = errors.New("No se pudo completar la operacion!")
	ErrNothingFind        = errors.New("No se encontro el registro!")
)
