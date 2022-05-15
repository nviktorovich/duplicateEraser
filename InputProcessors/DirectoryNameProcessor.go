package InputProcessors

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"os"
)

const ConfigPath = "InputProcessors/Config.yaml"

type DefaultSettings struct {
	PathName string `yaml:"Directory"`
}

func (f *FlagOptions) NameSet() error {
	if !f.Mode {
		n := new(DefaultSettings)
		file, err := os.ReadFile(ConfigPath)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(file, n)
		if err != nil {
			return err
		}
		f.FileName = n.PathName
	}
	fmt.Println(color.MagentaString("Название директории:"), f.FileName)
	return nil
}
