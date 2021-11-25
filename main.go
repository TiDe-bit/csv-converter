package main

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tushar2708/altcsv"
)

const inDir = "operation"

func main() {
	dir, err := os.ReadDir(inDir)
	if err != nil {
		os.Mkdir(inDir, 0700)
		log.Fatalln(log.WithError(err))
	}

	err = os.Chdir(inDir)
	if err != nil {
		log.Fatalln(log.WithError(err))
	}

	writers := make([]altcsv.Writer, len(dir))
	for i, f := range dir {
		//out, err := os.OpenFile(f.Name(), os.O_WRONLY, 0600)
		log.Info(f.Name(), "\tGenerated?\t", strings.Contains(f.Name(), "NEW"))
		if strings.Contains(f.Name(), "NEW") {
			log.Info("Generated File Found")
			continue
		}

		fileName := strings.Split(f.Name(), ".")[0] + " - NEW.csv"

		out, err := os.Create(fileName)
		if err != nil {
			log.Fatal(log.WithError(err))
		}
		writer := altcsv.NewWriter(out)
		writer.UseCRLF = true
		writer.AllQuotes = true

		writers[i] = *writer
		defer writer.Flush()
	}

	log.Infof("%v Files", len(dir))

	for i, d := range dir {
		w := writers[i]
		write(&w, d.Name())
	}
}

func write(w *altcsv.Writer, name string) {
	recordsSource, err := readData(name)
	if err != nil {
		log.Fatalln(log.WithError(err))
	}

	for _, sourceStrings := range recordsSource {
		emptyRow := sourceStrings
		newRow := make([]string, 0)

		for _, field := range emptyRow {
			newField := field
			//newField := fmt.Sprintf(`'%v'`, field)
			//newField = strings.Replace(newField, "'", "\042",2)
			newRow = append(newRow, newField)
			//emptyRow = append(emptyRow, fmt.Sprintf("%q", field))
		}

		noKD := strings.Split(sourceStrings[7], " KD"+
			" ")

		newField := fmt.Sprintf("%v - %v", noKD[0], sourceStrings[115])

		newRow[7] = newField

		err = w.Write(newRow)
		if err != nil {
			log.Fatalln(log.WithError(err))
		}
	}
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := altcsv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
