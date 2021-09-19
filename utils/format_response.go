package utils

const (
	CODE_OK               = "OK"
	CODE_BAD_REQUEST      = "ERROR"
	CODE_NO_ENCONTRADO    = "NO_FIND"
	CODE_INTERNAL_REQUEST = "ERR_II"
	COD_USUARIO_EXISTE    = "X000K"
	CODE_FORBIDDEN        = "FF"
)
const (
	InternalMessage = "No se pudo completar la operacion, inténtelo más tarde."
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type LogginResponse struct {
	Code    string      `json:"code"`
	Usuario interface{} `json:"usuario,omitempty"`
	Token   string      `json:"token"`
}

func NewForbiddenResponse(message string) *Response {
	if message == "" {
		message = "permiso denegado"
	}
	return &Response{
		Code:    CODE_FORBIDDEN,
		Message: message,
	}
}
func NewOkResponse(data interface{}) *Response {
	return &Response{
		Code: CODE_OK,
		Data: data,
	}
}
func NewOkMesage(message string) *Response {
	if message == "" {
		message = "Operacion exitosa"
	}
	return &Response{
		Code:    CODE_OK,
		Message: message,
	}
}
func NewNoFindResponse(message string) *Response {
	return &Response{
		Code:    CODE_NO_ENCONTRADO,
		Message: message,
	}
}
func NewBadResponse(message string) *Response {
	if message == "" {
		message = "Estructura de datos incorrecta"
	}
	return &Response{
		Code:    CODE_BAD_REQUEST,
		Message: message,
	}
}
func NewInternalErrorResponse(message string) *Response {
	if message == "" {
		message = "No se pudo completar la operacion"
	}
	return &Response{
		Code:    CODE_INTERNAL_REQUEST,
		Message: message,
	}
}
