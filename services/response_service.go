package services

import (
	"encoding/json"
	"log"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/models/vocacion"
	"github.com/user0608/expertos/pkg"
	"github.com/user0608/expertos/repository"
	"github.com/user0608/expertos/utils"
)

type ResponseService struct {
	repo    *repository.ResponseRepository
	experto *ExpertoService
}

func NewResponseService(r *repository.ResponseRepository, e *ExpertoService) *ResponseService {
	return &ResponseService{repo: r, experto: e}
}
func (s *ResponseService) prepare(v *vocacion.Test) (*vocacion.TestResponse, error) {
	casm := []pkg.Result{}
	berger := &pkg.BergerResult{}
	hea := []string{}
	if v.ResultadoCasm != "" {
		if err := json.Unmarshal([]byte(v.ResultadoCasm), &casm); err != nil {
			log.Println("Error-1:ResponseService.prepare:", err.Error())
			return nil, errs.ErrInnter
		}
	}
	if v.ResultadoBerger != "" {
		if err := json.Unmarshal([]byte(v.ResultadoBerger), berger); err != nil {
			log.Println("Error-2:ResponseService.prepare:", err.Error())
			return nil, errs.ErrInnter
		}
	}
	if v.ResultadoHea != "" {
		if err := json.Unmarshal([]byte(v.ResultadoHea), &hea); err != nil {
			log.Println("Error-3:ResponseService.prepare:", err.Error())
			return nil, errs.ErrInnter
		}
	}
	return &vocacion.TestResponse{
		ID:              v.ID,
		ResultadoCasm:   casm,
		ResultadoBerger: berger,
		ResultadoHea:    hea,
		Done:            v.Done,
		Fecha:           v.Fecha,
	}, nil
}
func (s *ResponseService) Procesar(t *vocacion.TestResponse) error {
	ber, err := s.experto.ConsultarBerger(&utils.RequestBerger{
		Emotivo: t.ResultadoBerger.Emotivo,
		Activo:  t.ResultadoBerger.Activo,
		Orden:   t.ResultadoBerger.Orden,
	})
	if err != nil {
		return err
	}
	t.ResBerger = ber.Respuesta
	var result []string
	for _, b := range t.ResultadoCasm {
		for _, h := range t.ResultadoHea {
			ex, err := s.experto.Consultar(&utils.RequestExperto{
				Area:     b.Res,
				Habito:   h,
				Caracter: t.ResBerger,
			})
			if err != nil {
				return err
			}
			if ex.Respuesta != "RESPUESTA NO ENCONTRADA" {
				flag := true
				for _, k := range result {
					if k == ex.Respuesta {
						flag = false
					}
				}
				if flag {
					result = append(result, ex.Respuesta)
				}
			}
		}
	}
	if len(result) == 0 {
		result = append(result, "NO SE A ENCONTRADO UNA RESPUESTA ADECUADA")
	}
	t.Resultado = result
	return nil
}
func (s *ResponseService) Response(TestID int) (*vocacion.TestResponse, error) {
	test, err := s.repo.Result(TestID)
	if err != nil {
		return nil, err
	}
	newTest, err := s.prepare(test)
	if newTest.Done {
		if err := s.Procesar(newTest); err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	return newTest, nil
}
