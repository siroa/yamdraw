package layout

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectKind(t *testing.T) {
	kind := COMMUNICATION
	k, _ := selectKind(kind)
	require.Equal(t, "comm", k)
	kind = REFERENCE
	k, _ = selectKind(kind)
	require.Equal(t, "ref", k)
	kind = UPDATE
	k, _ = selectKind(kind)
	require.Equal(t, "update", k)
	kind = "hogehoge"
	_, err := selectKind(kind)
	require.Error(t, err)
}
