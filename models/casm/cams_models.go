package casm

type CASMQuestion struct {
	ID        int    `json:"id"`
	QuestionA string `json:"question_a"`
	QuestionB string `json:"question_b"`
	AnswerA   bool   `json:"answer_a"`
	AnswerB   bool   `json:"answer_b"`
	Done      bool   `json:"done"`
}
