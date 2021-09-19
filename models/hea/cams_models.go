package hea

import (
	"errors"
	"strings"
)

type HEAQuestion struct {
	ID       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Done     bool   `json:"done"`
}

type TestHea struct {
	TestID    int    `json:"test_id"`
	HeaID     int    `json:"hea_id"`
	Respuesta string `json:"answer"`
}

func (t *TestHea) Valid() error {
	t.Respuesta = strings.TrimSpace(strings.ToUpper(t.Respuesta))
	answerLen := len(t.Respuesta)
	if answerLen == 0 || answerLen > 1 {
		return errors.New("Field: `answer` Invalido")
	}
	if !strings.Contains("SMPAN", t.Respuesta) {
		return errors.New("Field: `answer` Invalido")
	}
	if t.HeaID == 0 {
		return errors.New("Field: `hea_id` Invalido")
	}
	if t.TestID == 0 {
		return errors.New("Field: `test_id` Invalido")
	}
	return nil
}
