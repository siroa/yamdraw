package yaml

import (
	"errors"
	"generate/utils/logger"
)

type Mw struct {
	ID      string                         `yaml:"id,omitempty"`
	Name    string                         `yaml:"name,omitempty"`
	Process *[]map[interface{}]interface{} `yaml:"process,omitempty"`
	Db      *[]map[interface{}]interface{} `yaml:"db,omitempty"`
}

func (m *Mw) ReadMw() (string, string, *[]map[interface{}]interface{}, *[]map[interface{}]interface{}, error) {
	if m == nil {
		logger.Warn("Can't read mw. Please check yaml syntax...")
		return "", "", nil, nil, errors.New("Unexpected nil pointer")
	}
	return m.ID, m.Name, m.Process, m.Db, nil
}

func (m *Mw) ReadProcess() (*[]map[interface{}]interface{}, error) {
	if m == nil {
		logger.Warn("Can't read mw. Please check yaml syntax...")
		return nil, errors.New("Unexpected nil pointer")
	}
	return m.Process, nil
}

func (m *Mw) ReadDb() (*[]map[interface{}]interface{}, error) {
	if m == nil {
		logger.Warn("Can't read mw. Please check yaml syntax...")
		return nil, errors.New("Unexpected nil pointer")
	}
	return m.Db, nil
}
