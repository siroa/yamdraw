package app

import (
	"encoding/xml"
	"generate/domain/drawio"
	"generate/domain/layout"
	"generate/domain/yaml"
	"generate/service"
	"generate/utils/logger"
	"io/ioutil"
)

type GenerateHandlers struct {
	Yaml   service.YamlService
	Drawio service.DefaultDrawioService
	Layout service.DefaultLayoutService
}

func (gh *GenerateHandlers) ReadOverall() (*yaml.Overall, error) {
	b, err := gh.Yaml.Read()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (gh *GenerateHandlers) CreateDrawioFile(nls []*layout.LayoutNetwork, pls []layout.LayoutProcedure) drawio.Mxfile {
	service := gh.Drawio.NewDrawioService()
	file := gh.Drawio.Build(service, nls, pls)
	return file
}

func (gh *GenerateHandlers) WriteDrawioFile(file drawio.Mxfile, filename string) error {
	output, err := xml.MarshalIndent(&file, "", "  ")
	if err != nil {
		logger.Error("Can't Marshal to xml ..." + err.Error())
		return err
	}

	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		logger.Error("Can't write to file ..." + err.Error())
		return err
	}
	return nil
}

func (gh *GenerateHandlers) LocLayout(b *yaml.Overall) ([]*layout.LayoutNetwork, []layout.LayoutProcedure) {
	networks := b.Networks
	procedure := b.Procedure
	nls := gh.Layout.CreateNetworkLayout(networks)
	pls := []layout.LayoutProcedure{}
	if procedure != nil {
		pl := procedure.ToProcedureList()
		pls = gh.Layout.CreateProcedure(pl)
	}
	return nls, pls
}
