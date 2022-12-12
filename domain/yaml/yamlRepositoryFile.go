package yaml

import (
	"generate/driver"
	"generate/utils/logger"

	"gopkg.in/yaml.v3"
)

type YamlRepositoryFile struct {
	handler driver.YamlHandler
}

func NewYamlRepositoryFile(yml *driver.YamlHandler) YamlRepositoryFile {
	return YamlRepositoryFile{*yml}
}

func (y YamlRepositoryFile) Read() (*Overall, error) {
	var b Overall
	if err := yaml.Unmarshal([]byte(y.handler.File), &b); err != nil {
		logger.Error("yaml cannot unmarshal ..." + err.Error())
		return nil, err
	}
	return &b, nil
}
