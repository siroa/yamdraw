package yaml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadServer(t *testing.T) {
	sid := "server1"
	sname := "server1"
	var smws *[]Mws = nil
	s := Server{ID: sid, Name: sname, Mws: nil}
	id, name, mws := s.ReadServer()
	require.Equal(t, sid, id)
	require.Equal(t, sname, name)
	require.Equal(t, smws, mws)
}
