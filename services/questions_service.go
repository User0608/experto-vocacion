package services

import (
	"errors"
	"fmt"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"github.com/user0608/expertos/models/berger"
	"github.com/user0608/expertos/models/casm"
	"github.com/user0608/expertos/models/hea"
	"github.com/user0608/expertos/repository"
	repo "github.com/user0608/expertos/repository"
)

type QuestionService struct {
	casm   *repository.CASMRepository
	berger *repository.BergerRepository
	hea    *repository.HEARepository
}

func NewQuestionService(c *repo.CASMRepository, b *repo.BergerRepository, h *repo.HEARepository) *QuestionService {
	return &QuestionService{casm: c, berger: b, hea: h}
}

func (s *QuestionService) GetCASMQuestion(testID, NumOfItems, Page int) ([]casm.CASMQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.casm.FindQuestionsByPage(testID, NumOfItems, Page)
}
func (s *QuestionService) RegisterCASMQuestionAnswer(answers []casm.TestCasm) (*models.CreateQuestionResponse, error) {
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
	}
	return s.casm.RegisterAnswer(answers)
}

func (s *QuestionService) GetBergerQuestions(testID, NumOfItems, Page int) ([]berger.BergerQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.berger.FindQuestionsByPage(testID, NumOfItems, Page)
}
func (s *QuestionService) RegisterBergerQuestionAnswer(answers []berger.TestBerger) (*models.CreateQuestionResponse, error) {
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
	}
	return s.berger.RegisterAnswer(answers)
}

func (s *QuestionService) GetHEAQuestions(testID, NumOfItems, Page int) ([]hea.HEAQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.hea.FindQuestionsByPage(testID, NumOfItems, Page)
}

func (s *QuestionService) RegisterHEAQuestionAnswer(answers []hea.TestHea) (*models.CreateQuestionResponse, error) {
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
	}
	return s.hea.RegisterAnswer(answers)
}
