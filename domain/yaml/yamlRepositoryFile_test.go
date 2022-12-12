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
	// 正常読み込み
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

	// 異常読み込み(yamlの想定していない構造)
	testFilePath = filepath.Clean(`..\..\test\test01_break_struct.yaml`)
	y, err := driver.NewYamlHandler(testFilePath)
	require.NoError(t, err)
	nrf = NewYamlRepositoryFile(&y)
	b, err = nrf.Read()
	require.Error(t, err)
}
