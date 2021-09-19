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
func (s *QuestionService) RegisterCASMQuestionAnswer(testID int, answers []casm.TestCasm) (*models.CreateQuestionResponse, error) {
	responses := []casm.TestCasm{}
	if testID == 0 {
		return nil, errors.New("Test ID no puede ser 0")
	}
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
		w.TestID = testID
		responses = append(responses, w)
	}
	return s.casm.RegisterAnswer(responses)
}

func (s *QuestionService) GetBergerQuestions(testID, NumOfItems, Page int) ([]berger.BergerQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.berger.FindQuestionsByPage(testID, NumOfItems, Page)
}
func (s *QuestionService) RegisterBergerQuestionAnswer(testID int, answers []berger.TestBerger) (*models.CreateQuestionResponse, error) {
	responses := []berger.TestBerger{}
	if testID == 0 {
		return nil, errors.New("Test ID no puede ser 0")
	}
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
		w.TestID = testID
		responses = append(responses, w)
	}
	return s.berger.RegisterAnswer(responses)
}

func (s *QuestionService) GetHEAQuestions(testID, NumOfItems, Page int) ([]hea.HEAQuestion, error) {
	if testID == 0 || NumOfItems == 0 || Page == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.hea.FindQuestionsByPage(testID, NumOfItems, Page)
}

func (s *QuestionService) RegisterHEAQuestionAnswer(testID int, answers []hea.TestHea) (*models.CreateQuestionResponse, error) {
	responses := []hea.TestHea{}
	if testID == 0 {
		return nil, errors.New("Test ID no puede ser 0")
	}
	if len(answers) == 0 {
		return nil, errors.New("Datos no encontrados! No puede ser null")
	}
	for i, w := range answers {
		if err := w.Valid(); err != nil {
			return nil, fmt.Errorf("%v, registro %d", err, i+1)
		}
		w.TestID = testID
		responses = append(responses, w)
	}
	return s.hea.RegisterAnswer(responses)
}
