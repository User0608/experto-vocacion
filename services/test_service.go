package services

import (
	"strings"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/vocacion"
	"github.com/user0608/expertos/repository"
)

type TestService struct {
	repo *repository.TestRepository
}

func NewTestService(r *repository.TestRepository) *TestService {
	return &TestService{repo: r}
}
func (s *TestService) CreateTest(dni string) (*vocacion.Test, error) {
	if strings.TrimSpace(dni) == "" {
		return nil, errs.ErrInvalidData
	}
	return s.repo.CreateTest(dni)
}
func (s *TestService) FindAll() ([]vocacion.TestTable, error) {
	return s.repo.FindAll()
}
func (s *TestService) FindID(ID int) (*vocacion.TestTable, error) {
	if ID == 0 {
		return nil, errs.ErrInvalidData
	}
	return s.repo.FindByID(ID)
}
func (s *TestService) FindAllByEstudianteDNI(dni string) ([]vocacion.Test, error) {
	if strings.TrimSpace(dni) == "" {
		return nil, errs.ErrInvalidData
	}
	return s.repo.FindAllByEstudianteDNI(dni)
}

func (s *TestService) DeleteTest(ID int) error {
	if ID == 0 {
		return errs.ErrInvalidParam
	}
	return s.repo.DeleteTest(ID)
}
