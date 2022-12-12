package yaml

import (
	"fmt"
	"generate/driver"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestYamlRepositoryFile(t *testing.T) {
	// normal mode read
	fmt.Println(filepath.Abs("./"))
	testFilePath := filepath.Clean(`..\..\test\test01.yaml`)
	y, _ := driver.NewYamlHandler(testFilePath)
	nrf := NewYamlRepositoryFile(&y)
	b, _ := nrf.Read()

	var want Overall
	f, _ := ioutil.ReadFile(testFilePath)
	if err := yaml.Unmarshal([]byte(f), &want); err != nil {
		panic("cannot unmarshal test...")
	}
	require.Equal(t, &want, b)

	// Abnormal loading (structure not assumed by yaml)
	testFilePath = filepath.Clean(`..\..\test\test01_break_struct.yaml`)
	y, err := driver.NewYamlHandler(testFilePath)
	require.NoError(t, err)
	nrf = NewYamlRepositoryFile(&y)
	b, err = nrf.Read()
	require.Error(t, err)
}
