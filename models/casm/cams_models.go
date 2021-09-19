package casm

import "errors"

type TestCasm struct {
	TestID      int  `json:"test_id"`
	CasmID      int  `json:"casm_id"`
	Respuesta_a bool `json:"answer_a"`
	Respuesta_b bool `json:"answer_b"`
}

func (t *TestCasm) Valid() error {
	if t.CasmID == 0 {
		return errors.New("Field: `casm_id` Invalido")
	}
	if t.TestID == 0 {
		return errors.New("Field: `test_id` Invalido")
	}
	return nil
}

type CASMQuestion struct {
	ID        int    `json:"id"`
	QuestionA string `json:"question_a"`
	QuestionB string `json:"question_b"`
	AnswerA   bool   `json:"answer_a"`
	AnswerB   bool   `json:"answer_b"`
	Done      bool   `json:"done"`
}
