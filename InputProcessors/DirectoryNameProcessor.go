package InputProcessors

import (
	"gopkg.in/yaml.v2"
	"os"
)

type DefaultSettings struct {
	PathName string `yaml:"Directory"`
}

func (f *FlagOptions) NameSet() error {
	if !f.Mode {
		n := new(DefaultSettings)
		file, err := os.ReadFile("InputProcessors/Config.yaml")
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(file, n)
		if err != nil {
			return err
		}
		f.FileName = n.PathName
		return nil
	}
	return nil
}

//fmt.Println("указано имя файла:", color.GreenString(f.FileName))
//fmt.Println("указано имя файла:", color.YellowString(f.FileName))
