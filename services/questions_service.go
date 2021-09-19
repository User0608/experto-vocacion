package services

import (
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/berger"
	"github.com/user0608/expertos/models/casm"
	"github.com/user0608/expertos/repository"
)

type QuestionService struct {
	casm   *repository.CASMRepository
	berger *repository.BergerRepository
}

func NewQuestionService(c *repository.CASMRepository, b *repository.BergerRepository) *QuestionService {
	return &QuestionService{casm: c, berger: b}
}

func (s *QuestionService) GetCASMQuestion(testID, NumOfItems, Page int) ([]casm.CASMQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.casm.FindQuestionsByPage(testID, NumOfItems, Page)
}
func (s *QuestionService) GetBergerQuestions(testID, NumOfItems, Page int) ([]berger.BergerQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.berger.FindQuestionsByPage(testID, NumOfItems, Page)
}
