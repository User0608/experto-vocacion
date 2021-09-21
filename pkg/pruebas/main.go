package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/user0608/expertos/models/casm"
)

func GenerarCasmQuestions() []casm.QuestionResponse {
	answers := []casm.QuestionResponse{}
	file, err := os.OpenFile("casm.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		for _, c := range line {
			q := casm.QuestionResponse{}
			if c == 'A' {
				q.Respuesta_a = true
			}
			if c == 'B' {
				q.Respuesta_b = true
			}

			answers = append(answers, q)
		}
		if err == io.EOF {
			break
		}
	}
	return answers
}
func OutFile(answers []casm.QuestionResponse) error {
	file, err := os.OpenFile("casm.sql", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	// return json.NewEncoder(file).Encode(answers)
	if _, err := file.WriteString("insert into test_casm(test_id,casm_id,respuesta_a,respuesta_b) values \n"); err != nil {
		return err
	}
	for i, w := range answers {
		if _, err := file.WriteString(fmt.Sprintf("(%d,%d,%t,%t),\n", 1, i+1, w.Respuesta_a, w.Respuesta_b)); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	// err := OutFile(GenerarCasmQuestions())
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println("OK")
	// }
	// c := pkg.CasmOperator{}
	// err := c.SetAnswer(GenerarCasmQuestions())
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(c.Consistencia())
	// fmt.Println(c.CalculoVerasidad())
	// c.CalculoPuntajeDirecto()
	// c.CalculoPercentil()
	// fmt.Println(c.PuntajeDirecto)
	// fmt.Println(c.Percentil)
	// fmt.Println(c.Resultado())
	GenerarHEA()
}

func GenerarHEA() {
	values := []string{"S", "M", "P", "A", "N"}
	_ = values
	file, err := os.OpenFile("hea.sql", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	_, _ = file.WriteString("insert into test_hea(test_id,hea_id,respuesta) values \n")
	for i := 0; i < 77; i++ {
		line := values[rand.Intn(5)]
		_, err = file.WriteString(fmt.Sprintf("(%d,%d,'%s'),\n", 1, i+1, line))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
