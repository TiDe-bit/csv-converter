package converter

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tushar2708/altcsv"
)

const inDir = "operation"

type FromToPair struct {
	From int
	To   int
}

func WithOptions(noKd bool) *ConvertOptions {
	opts := &ConvertOptions{NoKD: noKd}
	return opts
}

type ConvertOptions struct {
	FromToPairs map[string]FromToPair
	NoKD        bool
}

func (c *ConvertOptions) AddOption(pair FromToPair) *ConvertOptions {
	fmt.Printf("add %+v", pair)
	if c.FromToPairs == nil {
		c.FromToPairs = make(map[string]FromToPair)
	}
	c.FromToPairs[fmt.Sprintf("%d", pair.To)] = pair
	return c
}

func (c *ConvertOptions) RemoveOption(pair FromToPair) *ConvertOptions {
	fmt.Printf("remove %+v", pair)
	if c.FromToPairs == nil {
		return c
	}
	_, ok := c.FromToPairs[fmt.Sprintf("%d", pair.To)]
	if ok {
		delete(c.FromToPairs, fmt.Sprintf("%d", pair.To))
	}
	return c
}

func SpreadColumns() []string {
	wd, _ := os.Getwd()
	inOperationDir := path.Dir(wd) == inDir

	log.Debugf("spreadColums, %s", wd)

	dir, err := os.ReadDir(inDir)
	if err != nil {
		os.Mkdir(inDir, 0700)
		log.Error(log.WithError(err))
	}

	if !inOperationDir {
		err = os.Chdir(inDir)
		if err != nil {
			log.Error(log.WithError(err))
			return []string{"nothing here - not even a dir"}
		}
		defer os.Chdir("../")
	}

	if len(dir) <= 0 {
		return []string{"nothing here"}
	}
	data, err := readData(dir[0].Name())
	if err != nil {
		return nil
	}

	return data[0]
}

func Convert(options *ConvertOptions) {
	wd, _ := os.Getwd()
	inOperationDir := path.Dir(wd) == inDir

	if inOperationDir {
		os.Chdir("../")
	}
	dir, err := os.ReadDir(inDir)
	if err != nil {
		os.Mkdir(inDir, 0700)
		log.Warning(log.WithError(err))
	}

	if !inOperationDir {
		err = os.Chdir(inDir)
		if err != nil {
			log.Warning(log.WithError(err))
		}
		defer os.Chdir("../")
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
		write(&w, d.Name(), options)
	}
}

func write(w *altcsv.Writer, name string, opts *ConvertOptions) {
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

		tmpRow := newRow
		for _, option := range opts.FromToPairs {
			tmpRow[option.To] = newRow[option.From]
		}

		if opts.NoKD {
			noKD := strings.Split(sourceStrings[7], " KD"+
				" ")

			newField := fmt.Sprintf("%v - %v", noKD[0], sourceStrings[115])

			tmpRow[7] = newField
		}

		err = w.Write(tmpRow)
		if err != nil {
			log.Fatalln(log.WithError(err))
		}
	}
}

func readData(fileName string) ([][]string, error) {
	log.Debugf("filename %s", fileName)
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
