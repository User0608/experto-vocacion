package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/hea"
	"gorm.io/gorm"
)

type HEARepository struct {
	gorm *gorm.DB
}

func NewHEARepository(db *gorm.DB) *HEARepository {
	return &HEARepository{
		gorm: db,
	}
}

func (r *HEARepository) FindQuestionsByPage(testID, NumOfItems, Page int) ([]hea.HEAQuestion, error) {
	questions := []hea.HEAQuestion{}
	tr := r.gorm.Raw("select * from preguntas_hea_page(?,?,?)", testID, NumOfItems, Page).Scan(&questions)
	if tr.Error != nil {
		log.Println("Error-0: CASMRepository.FindQuestionsByPage:", tr.Error.Error())
		return nil, errs.ErrDataBaseError
	}
	if tr.RowsAffected == 0 {
		return nil, errs.ErrNothingFind
	}
	return questions, nil
}
