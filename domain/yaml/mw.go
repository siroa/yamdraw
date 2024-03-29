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
	Doc     *[]map[interface{}]interface{} `yaml:"document,omitempty"`
}

func (m *Mw) ReadMw() (string, string, map[string]*[]map[interface{}]interface{}, error) {
	if m == nil {
		logger.Warn("Can't read mw. Please check yaml syntax...")
		return "", "", nil, errors.New("Unexpected nil pointer")
	}
	accessories := map[string]*[]map[interface{}]interface{}{"db": m.Db, "process": m.Process, "doc": m.Doc}
	return m.ID, m.Name, accessories, nil
}
