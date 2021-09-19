package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/casm"
	"gorm.io/gorm"
)

type CASMRepository struct {
	gorm *gorm.DB
}

func NewCASMRepository(db *gorm.DB) *CASMRepository {
	return &CASMRepository{
		gorm: db,
	}
}

func (r *CASMRepository) FindQuestionsByPage(testID, NumOfItems, Page int) ([]casm.CASMQuestion, error) {
	questions := []casm.CASMQuestion{}
	tr := r.gorm.Raw("select * from preguntas_casm_page(?,?,?)", testID, NumOfItems, Page).Scan(&questions)
	if tr.Error != nil {
		log.Println("Error-0: CASMRepository.FindQuestionsByPage:", tr.Error.Error())
		return nil, errs.ErrDataBaseError
	}
	if tr.RowsAffected == 0 {
		return nil, errs.ErrNothingFind
	}
	return questions, nil
}
