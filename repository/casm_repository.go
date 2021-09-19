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
func (r *CASMRepository) RegisterAnswer(answer []casm.TestCasm) (*casm.CreateResponse, error) {
	created := make([]casm.TestCasm, 0)
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
		row := tx.Raw("select * from fn_consult_insertable_answer('casm',?,?)", w.TestID, w.CasmID).Row()
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
	return &casm.CreateResponse{
		NumCreated: createdLen,
		NumOmitted: len(answer) - createdLen,
		Created:    created,
	}, nil
}
