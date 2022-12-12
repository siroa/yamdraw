package driver

import (
	"generate/utils/logger"
	"io/ioutil"
)

type YamlHandler struct {
	File []byte
}

func NewYamlHandler(path string) (YamlHandler, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Error("Can't read yaml file..." + err.Error())
		return YamlHandler{}, err
	}
	yml := YamlHandler{File: buf}
	return yml, nil
}
