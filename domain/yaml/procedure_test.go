package yaml

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProcedure(t *testing.T) {
	p := Procedure{
		A: "1->",
		X: "hogehoge",
	}
	want := map[string]string{"A": "1->", "X": "hogehoge"}
	require.Equal(t, want, p.ToProcedureList())
}
