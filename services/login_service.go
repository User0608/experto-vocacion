package services

import (
	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"github.com/user0608/expertos/repository"
)

type LoginService struct {
	repo *repository.LoginRepository
}

func NewLoginService(er *repository.LoginRepository) *LoginService {
	return &LoginService{
		repo: er,
	}
}

func (s *LoginService) LogginEstudiante(d *models.Session) (*models.Estudiante, error) {
	if d.Password == "" || d.Username == "" {
		return nil, errs.ErrInvalidData
	}
	return s.repo.Loging(d)
}

func (s *LoginService) LogginAdmin(d *models.Session) error {
	if d.Password == "" || d.Username == "" {
		return errs.ErrInvalidData
	}
	return s.repo.LogingAdmin(d)
}
