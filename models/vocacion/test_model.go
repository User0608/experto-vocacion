package vocacion

import "github.com/user0608/expertos/pkg"

type RequestTest struct {
	ID           int
	EstudianteID int
}

type Test struct {
	ID              int    `json:"test_id"`
	ResultadoCasm   string `json:"resultado_casm"`
	ResultadoBerger string `json:"resultado_berger"`
	ResultadoHea    string `json:"resultado_hea"`
	Done            bool   `json:"done"`
	Resultado       string `json:"resultado"`
	Fecha           string `json:"created_at"`
}
type TestResponse struct {
	ID              int               `json:"test_id"`
	ResultadoCasm   []pkg.Result      `json:"resultado_casm"`
	ResultadoBerger *pkg.BergerResult `json:"resultado_berger"`
	ResBerger       string            `json:"resultado_berger_final,omitempty"`
	ResultadoHea    []string          `json:"resultado_hea"`
	Done            bool              `json:"done"`
	Resultado       interface{}       `json:"resultado"`
	Fecha           string            `json:"created_at"`
}

type TestTable struct {
	Test
	EstudianteID int `json:"estudiante_id"`
}

func (*TestTable) TableName() string {
	return "test"
}
