package repository

import (
	"log"
	"time"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/vocacion"
	"gorm.io/gorm"
)

type TestRepository struct {
	gorm *gorm.DB
}

func NewTestRepository(db *gorm.DB) *TestRepository {
	return &TestRepository{
		gorm: db,
	}
}

func (r *TestRepository) CreateTest(dni string) (*vocacion.Test, error) {
	estudianteID, err := FindEstudianteIDWithDNI(r.gorm, dni)
	if err != nil {
		return nil, err
	}
	row := r.gorm.Raw("insert into test(estudiante_id) values(?) returning id", estudianteID).Row()
	var testID int
	if err := row.Err(); err != nil {
		log.Println("Error-0: TestRepository.CreateTest:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	if err := row.Scan(&testID); err != nil {
		log.Println("Error-1: TestRepository.CreateTest:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	return &vocacion.Test{ID: testID, Done: false, Fecha: time.Now().Format(time.RFC3339)}, nil
}
func (r *TestRepository) DeleteTest(testID int) error {
	var ok bool
	row := r.gorm.Raw("select * from check_test_isnot_used(?)", testID).Row()
	if err := row.Err(); err != nil {
		log.Println("Error-0: TestRepository.DeleteTest:", err.Error())
		return errs.ErrDataBaseError
	}
	if err := row.Scan(&ok); err != nil {
		log.Println("Error-1: TestRepository.DeleteTest:", err.Error())
		return errs.ErrDataBaseError
	}
	if res := r.gorm.Delete(&vocacion.Test{ID: testID}); res.Error != nil {
		log.Println("Error-1: TestRepository.DeleteTest:", res.Error.Error())
		return errs.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return errs.ErrNothingWasDeleted
		}
	}
	return nil
}
func (r *TestRepository) FindByID(ID int) (*vocacion.TestTable, error) {
	test := &vocacion.TestTable{}
	res := r.gorm.Find(test, ID)
	if err := res.Error; err != nil {
		log.Println("Error-0: TestRepository.FindByID:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	if res.RowsAffected == 0 {
		return nil, errs.ErrNothingFind
	}
	return test, nil
}
func (r *TestRepository) FindAll() ([]vocacion.TestTable, error) {
	tests := []vocacion.TestTable{}
	res := r.gorm.Find(&tests)
	if err := res.Error; err != nil {
		log.Println("Error-1: TestRepository.FindAll:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	return tests, nil
}
func (r *TestRepository) FindAllByEstudianteDNI(dni string) ([]vocacion.Test, error) {
	tests := []vocacion.Test{}
	estudianteID, err := FindEstudianteIDWithDNI(r.gorm, dni)
	if err != nil {
		return nil, err
	}
	res := r.gorm.Find(&tests, "estudiante_id = ?", estudianteID)
	if err := res.Error; err != nil {
		log.Println("Error-2: TestRepository.FindAll:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	return tests, nil

}
