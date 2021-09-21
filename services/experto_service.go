package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/user0608/expertos/errs"
	"github.com/user0608/expertos/utils"
)

type ExpertoService struct {
	ApiUrl    string
	ApiBerger string
}

func NewExpertoService(url string, bergerUR string) *ExpertoService {
	return &ExpertoService{ApiUrl: url, ApiBerger: bergerUR}
}
func (s *ExpertoService) Consultar(r *utils.RequestExperto) (*utils.ExpertResponse, error) {
	d, err := json.Marshal(r)
	if err != nil {
		log.Println("Error-0: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrDataTypeOrStruct
	}
	req, err := http.NewRequest(http.MethodPost, s.ApiUrl, bytes.NewReader(d))
	if err != nil {
		log.Println("Error-1: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrInternal
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error-2: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrApiConnection
	}
	rr := &utils.ExpertResponse{}
	if err := json.NewDecoder(res.Body).Decode(rr); err != nil {
		log.Println("Error-3: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrServiceResponse
	}
	return rr, nil
}
func (s *ExpertoService) ConsultarBerger(r *utils.RequestBerger) (*utils.ExpertResponse, error) {
	d, err := json.Marshal(r)
	if err != nil {
		log.Println("Error-0: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrDataTypeOrStruct
	}
	req, err := http.NewRequest(http.MethodPost, s.ApiBerger, bytes.NewReader(d))
	if err != nil {
		log.Println("Error-1: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrInternal
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error-2: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrApiConnection
	}
	rr := &utils.ExpertResponse{}
	if err := json.NewDecoder(res.Body).Decode(rr); err != nil {
		log.Println("Error-3: ExpertoService.Consultar:", err.Error())
		return nil, errs.ErrServiceResponse
	}
	return rr, nil
}
