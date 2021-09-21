package pkg

import (
	"errors"
	"fmt"

	"github.com/user0608/expertos/models/casm"
)

var HEAD = []string{"CCFM", "CCSS", "CCNA", "CCCO", "ARTE", "BURO", "CCLP", "IIAA", "FINA", "LING", "JURI"}
var CASM = map[string]string{
	"CCFM": "Ciencias físicas matemáticas",
	"CCSS": "Ciencias sociales",
	"CCNA": "Ciencias naturales",
	"CCCO": "Ciencias de la comunicación",
	"ARTE": "Artes",
	"BURO": "Burocracia",
	"CCLP": "Ciencias económicas políticas",
	"IIAA": "Fuerzas armadas de Perú",
	"FINA": "Finanzas",
	"LING": "Lingüística",
	"JURI": "Jurisprudencia",
}
var percentiles = [][]int{
	{99, 21, 20, 21, 20, 19, 17, 20, 21, 18, 20, 22},
	{95, 18, 17, 19, 17, 16, 14, 16, 19, 25, 16, 19},
	{90, 17, 15, 17, 15, 13, 12, 14, 17, 13, 14, 16},
	{85, 16, 14, 16, 14, 12, 11, 13, 16, 12, 13, 14},
	{80, 15, 14, 15, 12, 12, 10, 12, 15, 12, 12, 12},
	{75, 14, 13, 14, 12, 11, 9, 11, 14, 11, 10, 10},
	{70, 13, 12, 13, 11, 10, 8, 10, 13, 10, 10, 9},
	{65, 13, 11, 12, 10, 9, 7, 9, 12, 9, 9, 7},
	{60, 12, 11, 11, 9, 9, 6, 8, 11, 9, 8, 6},
	{55, 11, 10, 10, 8, 8, 6, 7, 10, 8, 7, 5},
	{50, 11, 9, 9, 8, 7, 5, 6, 10, 7, 7, 4},
	{45, 10, 9, 8, 7, 7, 4, 5, 9, 6, 6, 3},
	{40, 10, 8, 7, 6, 6, 3, 4, 8, 5, 5, 2},
	{35, 9, 7, 6, 5, 5, 3, 4, 7, 4, 4, 1},
	{30, 8, 7, 5, 4, 5, 2, 3, 6, 4, 3, 1},
	{25, 7, 6, 4, 4, 4, 2, 2, 4, 3, 3, -1},
	{20, 7, 5, 3, 3, 3, 1, 2, 3, 2, 2, -1},
	{15, 6, 4, 2, 2, 3, 1, 1, 2, 1, 1, -1},
	{10, 4, 1, 2, 1, 2, -1, -1, 1, 1, 1, -1},
	{5, 2, 2, 1, -1, 1, -1, -1, -1, -1, -1, -1},
	{1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
}
var values = []int{14, 13, 14, 11, 11, 12, 13, 13, 11, 10, 11}

type Result struct {
	Per int    `json:"per"`
	Res string `json:"res"`
}
type CasmOperator struct {
	answers        [][]casm.QuestionResponse
	PuntajeDirecto []int
	Percentil      []int
	Result         []Result
}

func (o *CasmOperator) Consistencia() bool {
	var num int
	for i := 0; i < 11; i++ {
		if i == 10 {
			if o.answers[0][0].Respuesta_a == o.answers[10][12].Respuesta_a &&
				o.answers[0][0].Respuesta_b == o.answers[10][12].Respuesta_b {
				num++
			}
		} else {
			if o.answers[10][i].Respuesta_a == o.answers[i][12].Respuesta_a &&
				o.answers[10][i].Respuesta_b == o.answers[i][12].Respuesta_b {
				num++
			}
		}
	}
	fmt.Println(num)
	return num <= 5
}
func (o *CasmOperator) CalculoVerasidad() bool {
	var num int
	for i := 0; i < 11; i++ {
		if o.answers[i][12].Respuesta_a {
			num++
		}
	}
	fmt.Println(num)
	return num < 5
}
func (o *CasmOperator) SetAnswer(ws []casm.QuestionResponse) error {
	if len(ws) != 143 {
		return errors.New("Cantidad de preguntas insuficientes")
	}
	for i := 0; i < 143; i += 13 {
		o.answers = append(o.answers, ws[i:i+13])
	}
	return nil
}
func (o *CasmOperator) CalculoPuntajeDirecto() {
	for i := 0; i < 11; i++ {
		var suma int
		for j := 0; j < 11; j++ {
			if o.answers[i][j].Respuesta_b {
				suma += 1
			}
			if o.answers[j][i].Respuesta_a {
				suma += 1
			}
		}
		o.PuntajeDirecto = append(o.PuntajeDirecto, suma)
	}
}
func (o *CasmOperator) CalculoPercentil() error {
	if len(o.PuntajeDirecto) == 0 {
		return errors.New("No hay puntajes directos")
	}
	var last int
	for i, p := range o.PuntajeDirecto {
		for j := 0; j < 21; j++ {
			if percentiles[j][i+1] < p {
				o.Percentil = append(o.Percentil, last)
				break
			}
			last = percentiles[j][0]
		}
	}
	return nil
}

// func (o *CasmOperator) orderna() {

// }
func (o *CasmOperator) Resultado() []Result {
	for i := 0; i < 11; i++ {
		if o.PuntajeDirecto[i] > values[i] {
			o.Result = append(o.Result, Result{Per: o.Percentil[i], Res: HEAD[i]})
		}
	}
	return o.Result
}
