package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Casm struct {
	PreguntaA string `json:"pregunta_a"`
	PreguntaB string `json:"pregunta_b"`
}
type HEA string

func generateSQLInsert(datos []Casm, w io.StringWriter) {
	fmt.Printf("Se generaria %d registros\n", len(datos))
	_, err := w.WriteString("insert into berger(pregunta_a,pregunta_b) values\n")
	if err != nil {
		log.Fatalln("No se pudo escribir en el archivo1")
	}
	for _, c := range datos {
		_, err := w.WriteString(fmt.Sprintf("('%s','%s'),\n", c.PreguntaA, c.PreguntaB))
		if err != nil {
			log.Fatalln("No se pudo escribir en el archivo1")
		}
	}
}
func generateSQLInsertHEA(datos []HEA, w io.StringWriter) {
	fmt.Printf("Se generaria %d registros\n", len(datos))
	_, err := w.WriteString("insert into hea(pregunta) values\n")
	if err != nil {
		log.Fatalln("No se pudo escribir en el archivo1")
	}
	for _, c := range datos {
		_, err := w.WriteString(fmt.Sprintf("('%s'),\n", c))
		if err != nil {
			log.Fatalln("No se pudo escribir en el archivo1")
		}
	}
}

func prepareCasm(f *os.File) []Casm {
	reader := bufio.NewReader(f)
	flag := true
	datos := []Casm{}
	casm := Casm{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			datos = append(datos, casm)
			break
		} else {
			if err != nil {
				log.Fatalln("No se pudo continuar:", err.Error())
			}
		}
		if strings.TrimSpace(line) == "" {
			datos = append(datos, casm)
			flag = true
		} else {
			if flag {
				casm.PreguntaA = strings.TrimSpace(line)
				flag = false
			} else {
				casm.PreguntaB = strings.TrimSpace(line)
			}
		}
	}

	return datos
}
func PrepareHEA(f *os.File) []HEA {
	reader := bufio.NewReader(f)
	datos := []HEA{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatalln("No se pudo continuar:", err.Error())
			}
		}
		datos = append(datos, HEA(strings.TrimSpace(line)))
	}

	return datos
}

func main() {
	file, err := os.OpenFile("hea.txt", os.O_RDONLY, 0777)
	if err != nil {
		log.Fatalln("No se pudo continuar:", err.Error())
	}
	defer file.Close()
	outFile, err := os.OpenFile("hea_result.sql", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln("No se pudo continuar:", err.Error())
	}
	defer outFile.Close()
	generateSQLInsertHEA(PrepareHEA(file), outFile)
	//generateSQLInsert(prepareCasm(file), outFile)
	fmt.Println("OK")
}
