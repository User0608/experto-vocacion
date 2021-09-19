package berger

import (
	"errors"
)

type BergerQuestion struct {
	ID        int    `json:"id"`
	QuestionA string `json:"question_a"`
	QuestionB string `json:"question_b"`
	Answer    int    `json:"answer"`
	Done      bool   `json:"done"`
}
type TestBerger struct {
	TestID   int `json:"test_id"`
	BergerID int `json:"berger_id"`
	Response int `json:"answer"`
}

func (t *TestBerger) Valid() error {
	if !(t.Response == 1 || t.Response == 5 || t.Response == 9) {
		return errors.New("Field: `answer` Invalido")
	}
	if t.BergerID == 0 {
		return errors.New("Field: `berger_id` Invalido")
	}
	if t.TestID == 0 {
		return errors.New("Field: `test_id` Invalido")
	}
	return nil
}
