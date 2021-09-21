package pkg

import "github.com/user0608/expertos/models/hea"

var factor = map[string]int{"S": 1, "M": 2, "P": 3, "A": 4, "N": 5}
var habitos = map[int]string{
	0: "ACTITUD FRENTE AL ESTUDIO",
	1: "ANSIEDAD",
	2: "AUTOEVALUACION",
	3: "AYUDAS PARA EL ESTUDIO",
	4: "CONCENTRACION",
	5: "ESTRATEGIAS PARA EL EXAMEN",
	6: "MOTIVACION",
	7: "PROCESAMIENTO DE LA INFORMACION",
	8: "SELECCION DE IDEAS PRINCIPALES",
	9: "USO DEL TIEMPO",
}
var table = [][]int{
	{5, 14, 18, 29, 38, 45, 51, 69},
	{1, 9, 25, 31, 35, 54, 57, 63},
	{4, 17, 21, 26, 30, 37, 65, 70},
	{7, 19, 24, 44, 50, 53, 62, 73},
	{6, 11, 39, 43, 46, 55, 61, 68},
	{20, 27, 34, 52, 59, 64, 71, 75},
	{10, 13, 16, 28, 33, 41, 49, 56},
	{12, 15, 23, 32, 40, 47, 67, 76},
	{2, 8, 60, 72, 77},
	{3, 22, 36, 42, 48, 59, 66, 74},
}

type HeaOperador struct {
	Resultado map[string]int
}

func (h *HeaOperador) SetAnswers(ws []hea.TestHea) {
	h.Resultado = make(map[string]int)
	for i, row := range table {
		sum := 0
		for i, v := range row {
			sum += factor[ws[v-1].Respuesta] * (i + 1)
		}
		h.Resultado[habitos[i]] = sum
	}
}
func (h *HeaOperador) Result() []string {
	var aux string
	var last string
	for k, v := range h.Resultado {
		if h.Resultado[aux] < v {
			last = aux
			aux = k
		}
	}
	return []string{aux, last}
}
