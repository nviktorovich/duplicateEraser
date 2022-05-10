package InputSettings

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var (
	Method = flag.Bool("m",
		false,
		"выбор опции автоматически/вручную, если флаг установлен, "+
			"необходимо указать название директории для поиска дубликатов")
)

type Config struct {
	Directory string `yaml:"Directory"`
}

// AnaliseOptions принимает параметры, переданные/не переданные в командной строке
func AnaliseOptions() (string, error) {
	flag.Parse()
	selector := *Method
	dir, err := GetData(&selector, flag.Args())
	if err != nil {
		return "", err
	}
	return dir, nil
}

// GetData функция, которая возвращает название директории для поиска дубликатов.
// Если установлен флаг -m, то осуществляется ручной ввод названия директории,
// где будет осуществляться поиск дубликатов файлов. Если флаг не установле, то
// информация забирается из файла Default.yml.
// Функция возвращает строку - название директории и ошибку, в случае невозможности возврата названия.
func GetData(b *bool, s []string) (string, error) {

	if *b {
		if len(s) != 1 {
			err := errors.New("ошибка. установлен флаг -m, но не передано название директории")
			return "", err
		}
		return s[0], nil
	} else {
		data, err := ReadDefault()
		if err != nil {
			return "", err
		}
		return data, nil
	}
}

// ReadDefault функцмя возврашает строку с названием директории по умолчанию. Для
// этого используется *.yml файл - Default.yml
func ReadDefault() (string, error) {
	var Data Config

	filename, err := filepath.Abs("./InputSettings/Default.yml")
	if err != nil {
		return "", err
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if err = yaml.Unmarshal(yamlFile, &Data); err != nil {
		return "", err
	}
	return Data.Directory, nil
}
