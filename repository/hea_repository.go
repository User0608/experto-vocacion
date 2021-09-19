package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models"
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
func (r *HEARepository) RegisterAnswer(answer []hea.TestHea) (*models.CreateQuestionResponse, error) {
	created := make([]hea.TestHea, 0)
	tx := r.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		log.Println("Error-0:CASMRepository.RegisterAnswer:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	for _, w := range answer {
		var ok bool
		row := tx.Raw("select * from fn_consult_insertable_answer('hea',?,?)", w.TestID, w.HeaID).Row()
		if err := row.Err(); err != nil {
			tx.Rollback()
			log.Println("Error-1:CASMRepository.RegisterAnswer:", err.Error())
			return nil, errs.ErrDataBaseError
		}
		if err := row.Scan(&ok); err != nil {
			tx.Rollback()
			log.Println("Error-2:CASMRepository.RegisterAnswer:", err.Error())
			return nil, errs.ErrDataBaseError
		}
		if ok {
			if res := tx.Create(&w); res.Error != nil {
				tx.Rollback()
				log.Println("Error-3:CASMRepository.RegisterAnswer:", res.Error.Error())
				return nil, errs.ErrDataBaseError
			}
			created = append(created, w)
		}
	}
	if err := tx.Commit().Error; err != nil {
		log.Println("Error-4:CASMRepository.RegisterAnswer:", err.Error())
		return nil, errs.ErrDataBaseError
	}
	createdLen := len(created)
	return &models.CreateQuestionResponse{
		NumCreated: createdLen,
		NumOmitted: len(answer) - createdLen,
		Created:    created,
	}, nil
}
