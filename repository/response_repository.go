package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/berger"
	"github.com/user0608/expertos/models/casm"
	"github.com/user0608/expertos/models/hea"
	"github.com/user0608/expertos/models/vocacion"
	"github.com/user0608/expertos/pkg"
	"gorm.io/gorm"
)

var NotOK = errors.New("No se encontro!")

type ResponseRepository struct {
	gorm *gorm.DB
}

func NewResponseRepository(g *gorm.DB) *ResponseRepository {
	return &ResponseRepository{gorm: g}
}

func (r *ResponseRepository) IsProcessableTest(entity string, testID int) (bool, error) {
	test := vocacion.Test{}
	if res := r.gorm.Find(&test); res.Error != nil {
		log.Println("Error-0: ResponseRepository.CheckCompleteAnswers: ", res.Error.Error())
		return false, errs.ErrDataBaseError
	} else {
		if res.RowsAffected == 0 {
			return false, errs.ErrNothingFind
		}
		if entity == "casm" && test.ResultadoCasm != "" {
			return false, nil
		}
		if entity == "berger" && test.ResultadoBerger != "" {
			return false, nil
		}
		if entity == "hea" && test.ResultadoHea != "" {
			return false, nil
		}
	}
	row := r.gorm.Raw(fmt.Sprintf("select count(tc) from test_%s tc where tc.test_id =?", entity), testID).Row()
	var num int
	if err := row.Err(); err != nil {
		log.Println("Error-1: ResponseRepository.CheckCompleteAnswers: ", err.Error())
		return false, errs.ErrDataBaseError
	}
	if err := row.Scan(&num); err != nil {
		log.Println("Error-2: ResponseRepository.CheckCompleteAnswers: ", err.Error())
		return false, errs.ErrDataBaseError
	}
	if entity == "casm" {
		return num == 143, nil
	}
	if entity == "berger" {
		return num == 30, nil
	}
	if entity == "hea" {
		return num == 77, nil
	}
	return false, nil
}

func (r *ResponseRepository) Result(testID int) (*vocacion.Test, error) {
	test := &vocacion.Test{}
	if res := r.gorm.Find(&test, testID); res.Error != nil {
		log.Println("Erro-1:ResponseRepository.Result", res.Error.Error())
		return nil, errs.ErrDataBaseError
	}
	if !test.Done {
		if test.ResultadoCasm == "" {
			if err := r.GenerarCasm(testID); err != nil {
				return nil, err
			}
		}
		if test.ResultadoBerger == "" {
			if err := r.GenerarBerger(testID); err != nil {
				return nil, err
			}
		}
		if test.ResultadoHea == "" {
			if err := r.GenerarHea(testID); err != nil {
				return nil, err
			}
		}
	} else {
		return test, nil
	}
	if res := r.gorm.Find(&test, testID); res.Error != nil {
		log.Println("Erro-1:ResponseRepository.Result", res.Error.Error())
		return nil, errs.ErrDataBaseError
	} else {
		if test.ResultadoBerger != "" && test.ResultadoCasm != "" && test.ResultadoHea != "" {
			res := r.gorm.Model(test).Where("id = ?", testID).Update("done", true)
			if res.Error != nil {
				return nil, errs.ErrDataBaseError
			}
		}
		return test, nil
	}
}
func (r *ResponseRepository) GenerarHea(testID int) error {
	answers := []hea.TestHea{}
	if ok, err := r.IsProcessableTest("hea", testID); err != nil {
		return err
	} else if ok {
		res := r.gorm.Find(&answers, "test_id = ?", testID)
		if res.Error != nil {
			log.Println("Erro-1:ResponseRepository.GenerarHea", res.Error.Error())
			return errs.ErrDataBaseError
		}
		if res.RowsAffected == 77 {
			c := pkg.HeaOperador{}
			c.SetAnswers(answers)
			rr := c.Result()
			js, err := json.Marshal(rr)
			if err != nil {
				log.Println("Erro-3:ResponseRepository.GenerarBerger", err.Error())
				return errs.ErrDataBaseError
			}
			res = r.gorm.Model(&vocacion.Test{}).Where("id = ?", testID).Update("resultado_hea", string(js))
			if res.Error != nil {
				log.Println("Erro-4:ResponseRepository.GenerarBerger", err.Error())
				return errs.ErrDataBaseError
			}
		}
	} else {
		return NotOK
	}
	return nil
}
func (r *ResponseRepository) GenerarBerger(testID int) error {
	answers := []berger.TestBerger{}
	if ok, err := r.IsProcessableTest("casm", testID); err != nil {
		return err
	} else if ok {
		res := r.gorm.Find(&answers, "test_id = ?", testID)
		if res.Error != nil {
			log.Println("Erro-1:ResponseRepository.GenerarBerger", res.Error.Error())
			return errs.ErrDataBaseError
		}
		if res.RowsAffected == 30 {
			c := pkg.BergerOperation{}
			c.SetAnswer(answers)
			rr := c.Resultado()
			js, err := json.Marshal(rr)
			if err != nil {
				log.Println("Erro-3:ResponseRepository.GenerarBerger", err.Error())
				return errs.ErrDataBaseError
			}
			res = r.gorm.Model(&vocacion.Test{}).Where("id = ?", testID).Update("resultado_berger", string(js))
			if res.Error != nil {
				log.Println("Erro-4:ResponseRepository.GenerarBerger", err.Error())
				return errs.ErrDataBaseError
			}
		}
	} else {
		return NotOK
	}
	return nil
}
func (r *ResponseRepository) GenerarCasm(testID int) error {
	answers := []casm.QuestionResponse{}
	if ok, err := r.IsProcessableTest("casm", testID); err != nil {
		return err
	} else if ok {
		res := r.gorm.Find(&answers, "test_id = ?", testID)
		if res.Error != nil {
			log.Println("Erro-1:ResponseRepository.GenerarCasm", res.Error.Error())
			return errs.ErrDataBaseError
		}
		if res.RowsAffected == 143 {
			c := pkg.CasmOperator{}
			_ = c.SetAnswer(answers)
			c.CalculoPuntajeDirecto()
			_ = c.CalculoPercentil()
			rs := c.Resultado()
			js, err := json.Marshal(rs)
			if err != nil {
				log.Println("Erro-3:ResponseRepository.GenerarCasm", err.Error())
				return errs.ErrDataBaseError
			}
			res := r.gorm.Model(&vocacion.Test{}).Where("id = ?", testID).Update("resultado_casm", string(js))
			if res.Error != nil {
				log.Println("Erro-4:ResponseRepository.GenerarCasm", err.Error())
				return errs.ErrDataBaseError
			}
		}
	} else {
		return NotOK
	}
	return nil
}
