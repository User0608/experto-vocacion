package services

import (
	"fmt"
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"github.com/user0608/expertos/repository"
	"github.com/user0608/kcheck"
)

type EstudianteService struct {
	repo  *repository.EstudianteRepository
	check kcheck.KChecker
}

func NewEstudianteRepository(r *repository.EstudianteRepository) *EstudianteService {
	return &EstudianteService{
		repo:  r,
		check: kcheck.New(),
	}
}

func (s *EstudianteService) Create(e *models.Estudiante) error {
	if e == nil {
		return errs.ErrInvalidData
	}
	if err := e.Validate(); err != nil {
		return err
	}
	return s.repo.Create(e)
}
func (s *EstudianteService) Update(e *models.Estudiante) error {
	if e == nil {
		return errs.ErrInvalidData
	}
	if err := e.Validate(); err != nil {
		return err
	}
	return s.repo.Update(e)
}

func (s *EstudianteService) Find() ([]models.Estudiante, error) {
	return s.repo.Find()
}

func (s *EstudianteService) FindByID(id int) (*models.Estudiante, error) {
	if id == 0 {
		return nil, errs.ErrNothingFind
	}
	return s.repo.FindByID(id)
}
func (s *EstudianteService) FindByDNI(dni string) (*models.Estudiante, error) {
	if err := s.check.Target("len=8", dni).Ok(); err != nil {
		log.Println("->", err.Error())
		return nil, fmt.Errorf("DNI Invalida")
	}
	return s.repo.FindByDNI(dni)
}
