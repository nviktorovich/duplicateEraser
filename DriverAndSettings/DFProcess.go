package DriverAndSettings

import "errors"

// этот пакет предназначен для получения данных из командной строки и их валидации

// DFBase папка/файл база - название файла, вспомогательная информация
type DFBase struct {
	Name string // Name имя файла/папки
	Path string // Path путь к файлу/папке
}

func UserRequest(d []string) (*DFBase, error) {
	return &DFBase{}, errors.New("df")
}
