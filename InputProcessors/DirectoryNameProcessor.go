package InputProcessors

import (
	"gopkg.in/yaml.v2"
	"os"
)

const ConfigPath = "InputProcessors/Config.yaml"

type DefaultSettings struct {
	PathName string `yaml:"Directory"`
}

// GetDefaultRootName возвращает строку и ошибку. Предназначена для обработки конфигурационного файла ConfigPath
func GetDefaultRootName() (string, error) {
	d := new(DefaultSettings)

	file, err := os.ReadFile(ConfigPath)
	if err != nil {
		return "", err
	}
	err = yaml.Unmarshal(file, d)
	if err != nil {
		return "", err
	}
	return d.PathName, nil

}
