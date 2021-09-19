package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/berger"
	"gorm.io/gorm"
)

type BergerRepository struct {
	gorm *gorm.DB
}

func NewBergerRepository(db *gorm.DB) *BergerRepository {
	return &BergerRepository{
		gorm: db,
	}
}

func (r *BergerRepository) FindQuestionsByPage(testID, NumOfItems, Page int) ([]berger.BergerQuestion, error) {
	questions := []berger.BergerQuestion{}
	tr := r.gorm.Raw("select * from preguntas_berger_page(?,?,?)", testID, NumOfItems, Page).Scan(&questions)
	if tr.Error != nil {
		log.Println("Error-0: BergerRepository.FindQuestionsByPage:", tr.Error.Error())
		return nil, errs.ErrDataBaseError
	}
	if tr.RowsAffected == 0 {
		return nil, errs.ErrNothingFind
	}
	return questions, nil
}
