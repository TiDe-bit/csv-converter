package options

import (
	"changeme/pkg/converter"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const fileName = "optionen.json"

func Load() *converter.ConvertOptions {
	loadedOptions := &converter.ConvertOptions{}
	wd, _ := os.Getwd()
	isInCorrectDir := path.Dir(wd) != "operation"

	if !isInCorrectDir {
		err := os.Chdir("../")
		if err != nil {
			logrus.Error(err)
		}
		return Load()
	}

	bytes, err := os.ReadFile(fileName)
	if err != nil {
		os.Create(fileName)
		bytes, _ = os.ReadFile(fileName)
	}

	err = json.Unmarshal(bytes, loadedOptions)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debugf("loaded options, %+v", loadedOptions)
	return loadedOptions
}

func Save(options *converter.ConvertOptions) {
	bytes, err := json.Marshal(options)
	if err != nil {
		return
	}

	fmt.Printf("save options, %s", string(bytes))

	wd, _ := os.Getwd()
	isInCorrectDir := path.Dir(wd) != "operation"

	if !isInCorrectDir {
		os.Chdir("..")
		Save(options)
		return
	}

	fmt.Printf("options to save %s", bytes)

	file, err := os.Open(fileName)
	if err != nil {
		file, err = os.Create(fileName)
		logrus.Warning(err)
	}
	defer file.Close()

	err = os.WriteFile(fileName, bytes, 0700)
	if err != nil {
		logrus.Error(err)
	}
}
