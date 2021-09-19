package injectors

import "github.com/user0608/expertos/handlers"

func GetLogginHandler() *handlers.LogginHandler {
	return loginHandler
}
func GetEstudianteHandler() *handlers.EstudianteHandler {
	return estudianteHandler
}
func GetQuestionHandler() *handlers.QuestionHandler {
	return questionHandler
}