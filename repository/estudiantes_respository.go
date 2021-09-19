package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
	"gorm.io/gorm"
)

type EstudianteRepository struct {
	gorm *gorm.DB
}

func NewEstudianteRepository(db *gorm.DB) *EstudianteRepository {
	return &EstudianteRepository{
		gorm: db,
	}
}

func (r *EstudianteRepository) Create(e *models.Estudiante) error {
	var consult string
	row := r.gorm.Raw("select * from valid_unique_dni(?)", e.Dni).Row()
	if err := row.Err(); err != nil {
		log.Println("Error-0: EstudianteRepository.Create:", err.Error())
		return errs.ErrDataBaseError
	}
	if err := row.Scan(&consult); err != nil {
		log.Println("Error-1: EstudianteRepository.Create:", err.Error())
		return errs.ErrDataBaseError
	}
	if consult != "OK" {
		return errs.ErrDNIExist
	}
	if result := r.gorm.Save(e); result.Error != nil {
		log.Println("Error-2: EstudianteRepository.Create:", result.Error)
		return errs.ErrDataBaseError
	}
	e.Password = ""
	return nil
}
func (r *EstudianteRepository) Update(e *models.Estudiante) error {
	if result := r.gorm.Omit("id").Where("dni = ?", e.Dni).Updates(e); result.Error != nil {
		log.Println("Error-2: EstudianteRepository.Update:", result.Error)
		return errs.ErrDataBaseError
	}
	e.Password = ""
	return nil
}

func (r *EstudianteRepository) Find() ([]models.Estudiante, error) {
	estudiantes := []models.Estudiante{}
	if result := r.gorm.Omit("password").Find(&estudiantes); result.Error != nil {
		log.Println("Error: EstudianteRepository.Find:", result.Error)
		return nil, errs.ErrDataBaseError
	}

	return estudiantes, nil
}
func (r *EstudianteRepository) FindByID(id int) (*models.Estudiante, error) {
	estudiante := &models.Estudiante{}
	if result := r.gorm.Limit(1).Omit("password").Find(estudiante, id); result.Error != nil {
		log.Println("Error: EstudianteRepository.FindByID:", result.Error)
		return nil, errs.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return nil, errs.ErrNothingFind
		}
	}
	estudiante.Password = ""
	return estudiante, nil
}
func (r *EstudianteRepository) FindByDNI(dni string) (*models.Estudiante, error) {
	estudiante := &models.Estudiante{}
	if result := r.gorm.Limit(1).Omit("password").Find(estudiante, "dni =?", dni); result.Error != nil {
		log.Println("Error: EstudianteRepository.FindByDNI:", result.Error)
		return nil, errs.ErrDataBaseError
	} else {
		if result.RowsAffected == 0 {
			return nil, errs.ErrNothingFind
		}
	}
	return estudiante, nil
}
