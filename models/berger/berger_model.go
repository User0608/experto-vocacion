package berger

type BergerQuestion struct {
	ID        int    `json:"id"`
	QuestionA string `json:"question_a"`
	QuestionB string `json:"question_b"`
	Answer    int    `json:"answer_a"`
	Done      bool   `json:"done"`
}
