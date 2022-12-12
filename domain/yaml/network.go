package yaml

type Network struct {
	ID      string     `yaml:"id,omitempty"`
	Name    string     `yaml:"name,omitempty"`
	Kind    string     `yaml:"kind,omitempty"`
	Servers *[]Servers `yaml:"servers,omitempty"`
}

func (n *Network) ReadNetwork() (string, string, string, *[]Servers) {
	return n.ID, n.Name, n.Kind, n.Servers
}
