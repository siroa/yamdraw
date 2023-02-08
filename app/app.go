package app

import (
	"generate/domain/yaml"
	"generate/driver"
	"generate/service"
	"generate/utils/logger"
	"path/filepath"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	inputpath = kingpin.Arg("yaml", "Specify a yaml file.").Required().String()
)

func Run() {

	kingpin.Parse()
	if filepath.Ext(*inputpath) != ".yaml" {
		logger.Error(*inputpath + " is not a yaml file.")
		return
	}
	filename := filepath.Base(*inputpath)
	outputpath := strings.Split(filename, ".")
	outputpath[len(outputpath)-1] = "drawio"
	fn := strings.Join(outputpath, ".")

	file, err := driver.NewYamlHandler(*inputpath)
	if err != nil {
		return
	}
	yr := yaml.NewYamlRepositoryFile(&file)
	ys := service.NewYamlService(yr)
	gh := GenerateHandlers{Yaml: ys}
	b, err := gh.ReadOverall()
	if err != nil {
		return
	}
	nls, pl := gh.LocLayout(b)
	// Todo: 全体像の把握

	f := gh.CreateDrawioFile(nls, pl)
	err = gh.WriteDrawioFile(f, fn)
	if err != nil {
		return
	}
}
