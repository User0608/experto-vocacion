package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"gorm.io/gorm"
)

type LoginRepository struct {
	gorm *gorm.DB
}

func NewLoginRepository(db *gorm.DB) *LoginRepository {
	return &LoginRepository{
		gorm: db,
	}
}

func (r *LoginRepository) Loging(s *models.Session) (*models.Estudiante, error) {
	estudiante := &models.Estudiante{}
	result := r.gorm.Raw("select * from ps_sign_in(?,?)", s.Username, s.Password).Scan(estudiante)
	if result.Error != nil {
		log.Println("Error : LoginRepository.Loging :", result.Error.Error())
		return nil, errs.ErrDataBaseError
	}
	if result.RowsAffected == 0 {
		return nil, errs.ErrUsernameOrPassword
	}
	estudiante.Password = ""
	return estudiante, nil
}
func (r *LoginRepository) LogingAdmin(s *models.Session) error {
	result := r.gorm.Raw("select * from ps_sign_admin(?,?)", s.Username, s.Password).Scan(&models.Usuario{})
	if result.Error != nil {
		log.Println("Error : LoginRepository.LogingAdmin :", result.Error.Error())
		return errs.ErrDataBaseError
	}
	if result.RowsAffected == 0 {
		return errs.ErrUsernameOrPassword
	}
	return nil
}
