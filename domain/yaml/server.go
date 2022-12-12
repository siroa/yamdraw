package yaml

type Server struct {
	ID   string `yaml:"id,omitempty"`
	Name string `yaml:"name,omitempty"`
	Mws  *[]Mws `yaml:"mws,omitempty"`
}

func (s *Server) ReadServer() (string, string, *[]Mws) {
	if s.Mws != nil {
		return s.ID, s.Name, s.Mws
	}
	return s.ID, s.Name, nil
}
