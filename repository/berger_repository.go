package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
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
func (r *BergerRepository) RegisterAnswer(answer []berger.TestBerger) (*models.CreateQuestionResponse, error) {
	created := make([]berger.TestBerger, 0)
	tx := r.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		log.Println("Error-0:BergerRepository.RegisterAnswer:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	for _, w := range answer {
		var ok bool
		row := tx.Raw("select * from fn_consult_insertable_answer('berger',?,?)", w.TestID, w.BergerID).Row()
		if err := row.Err(); err != nil {
			tx.Rollback()
			log.Println("Error-1:BergerRepository.RegisterAnswer:", err.Error())
			return nil, errs.ErrDataBaseError
		}
		if err := row.Scan(&ok); err != nil {
			tx.Rollback()
			log.Println("Error-2:BergerRepository.RegisterAnswer:", err.Error())
			return nil, errs.ErrDataBaseError
		}
		if ok {
			if res := tx.Create(&w); res.Error != nil {
				tx.Rollback()
				log.Println("Error-3:BergerRepository.RegisterAnswer:", res.Error.Error())
				return nil, errs.ErrDataBaseError
			}
			created = append(created, w)
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Println("Error-4:BergerRepository.RegisterAnswer:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	createdLen := len(created)
	return &models.CreateQuestionResponse{
		NumCreated: createdLen,
		NumOmitted: len(answer) - createdLen,
		Created:    created,
	}, nil
}
