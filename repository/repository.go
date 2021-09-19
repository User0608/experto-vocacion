package repository

import (
	"log"

	"github.com/user0608/expertos/errs"
	"gorm.io/gorm"
)

func FindEstudianteIDWithDNI(g *gorm.DB, dni string) (int, error) {
	var estudianteID int
	row := g.Raw("select id from estudiante where dni = ?", dni).Row()
	if err := row.Err(); err != nil {
		log.Println("Error-0: TestRepository.FindAll:", err.Error())
		return 0, errs.ErrDataBaseError
	}
	if err := row.Scan(&estudianteID); err != nil {
		log.Println("Error-1: TestRepository.FindAll:", err.Error())
		if err == gorm.ErrRecordNotFound {
			return 0, errs.ErrNothingFind
		}
	}
	return estudianteID, nil
}
