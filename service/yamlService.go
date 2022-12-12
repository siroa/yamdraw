package service

import (
	"generate/domain/yaml"
)

type DefaultYamlService struct {
	repo yaml.YamlRepository
}

// primary port
type YamlService interface {
	Read() (*yaml.Overall, error)
}

func NewYamlService(r yaml.YamlRepository) DefaultYamlService {
	return DefaultYamlService{r}
}

func (s DefaultYamlService) Read() (*yaml.Overall, error) {
	return s.repo.Read()
}
