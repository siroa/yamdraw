package yaml

type Overall struct {
	Networks  *[]Networks `yaml:"networks"`
	Procedure *Procedure  `yaml:"procedure,omitempty"`
}

// secondary port
type YamlRepository interface {
	Read() (*Overall, error)
}
