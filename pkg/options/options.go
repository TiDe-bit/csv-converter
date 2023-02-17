package options

import (
	"changeme/pkg/converter"
	"encoding/json"
	"os"
	"path"

	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Info("loading options failed")
		os.Create(fileName)
		bytes, _ = os.ReadFile(fileName)
	}

	if len(bytes) == 0 {
		logrus.Info("no options to load")
		return nil
	}

	err = json.Unmarshal(bytes, loadedOptions)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Debugf("loaded options, %+v", loadedOptions)
	return loadedOptions
}

func Save(options *converter.ConvertOptions) {
	if options == nil {
		logrus.Info("no settings to save")
		return
	}

	logrus.Infof("saving options %+v", options)

	bytes, err := json.Marshal(options)
	if err != nil {
		logrus.Fatal(err)
		return
	}

	logrus.Infof("save options, %s", string(bytes))

	wd, _ := os.Getwd()
	isInCorrectDir := path.Dir(wd) != "operation"

	if !isInCorrectDir {
		os.Chdir("..")
		Save(options)
		return
	}

	logrus.Infof("options to save %s", bytes)

	_, err = os.ReadFile(fileName)
	if err != nil {
		file, err := os.Create(fileName)
		if err != nil {
			logrus.Fatal(err)
		}
		file.Close()
	}

	err = os.WriteFile(fileName, bytes, 0700)
	if err != nil {
		logrus.Fatal(err)
	}
}
