package pkg

import "github.com/user0608/expertos/models/berger"

type BergerOperation struct {
	Emotivo int
	Activo  int
	Orden   int
}
type BergerResult struct {
	Emotivo string `json:"emotivo"`
	Activo  string `json:"activo"`
	Orden   string `json:"orden"`
}

func (b *BergerOperation) SetAnswer(ws []berger.TestBerger) {
	if len(ws) != 30 {
		return
	}
	for i := 0; i < 10; i++ {
		b.Emotivo += ws[i].Response
		b.Activo += ws[i+10].Response
		b.Orden += ws[i+20].Response
	}
}
func (b *BergerOperation) Resultado() (res BergerResult) {
	if b.Emotivo > 47 {
		res.Emotivo = "EMOTIVO"
	} else {
		res.Emotivo = "NO EMOTIVO"
	}
	if b.Activo > 54 {
		res.Activo = "ACTIVO"
	} else {
		res.Activo = "NO ACTIVO"
	}
	if b.Emotivo > 54 {
		res.Orden = "SECUNDARIO"
	} else {
		res.Orden = "PRIMARIO"
	}
	return
}
