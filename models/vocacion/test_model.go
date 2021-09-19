package vocacion

type RequestTest struct {
	ID           int
	EstudianteID int
}

type Test struct {
	ID              int    `json:"test_id"`
	ResultadoCasm   string `json:"resultado_casm"`
	ResultadoBerger string `json:"resultado_berger"`
	ResultadoLea    string `json:"resultado_lea"`
	Done            bool   `json:"done"`
	Fecha           string `json:"created_at"`
}

type TestTable struct {
	Test
	EstudianteID int `json:"estudiante_id"`
}

func (*TestTable) TableName() string {
	return "test"
}
