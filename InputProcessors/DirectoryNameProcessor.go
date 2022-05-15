package InputProcessors

import (
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
	return nil
}
