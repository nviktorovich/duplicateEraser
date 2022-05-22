package InputProcessors

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
)

var EmptyFileNameErr = errors.New("ошибка имени файла, название не може быть пустым")

// CommandLineOptions структура, предназначенная для хранения параметров командной строки и аргументов
type CommandLineOptions struct {
	M bool
	D bool
	A []string
}

// NewCommandLineOptions функция для анализа ввода командной строки. Возвращает
// объект с состоянием флагов, и список аргументов командной строки
func NewCommandLineOptions() CommandLineOptions {
	var (
		FlagInputSelector  bool
		FlagEraserSelector bool
	)

	flag.BoolVar(&FlagInputSelector, "m", false, "выбор режима ввода имени корневой директории: наличие флага - ручной; отстутвие - по умолчанию")
	flag.BoolVar(&FlagEraserSelector, "d", false, "выбор режима удаления дубликатов: наличие флага - удаление; отстутвие - режим просмотра")
	flag.Parse()

	return CommandLineOptions{FlagInputSelector, FlagEraserSelector, flag.Args()}
}

// SettingsOptions структура которая хранит состояние режима очистки true - удалять, false - нет и корневую директорию.
type SettingsOptions struct {
	Erase bool
	Root  string
}

// NewSettingsOptions принимает на вход структуру с данными ввода командной
// строки, возвращает структуру SettingsOptions, или ошибку
func NewSettingsOptions(c CommandLineOptions, fn func() (string, error)) (SettingsOptions, error) {
	s := new(SettingsOptions)
	var err error
	// 1 определить режим очистки
	s.Erase = c.D

	// 2 определить имя директории
	var rootName string
	switch c.M {
	// флаг m установлен, проверяем наличие аргументов. Если нет - ошибка. Проверяем
	// содержание первого аргумента, если пустая строка - ошибка, в остальных
	// сслучаях считаем, что имя корневой директории - это первый аргуммент
	case true:
		if len(c.A) < 1 {
			err = EmptyFileNameErr
		} else {
			if c.A[0] == "" {
				err = EmptyFileNameErr
			} else {
				s.Root = c.A[0]
			}
		}
		return *s, err

	// флаг m - не установлен, необходимо отработать с конфигруационным файлом
	default:
		rootName, err = fn()
		if err != nil {
			return *s, err
		}
		s.Root = rootName
		return *s, nil
	}

	// описать ошибку когда есть флаг ручного ввода, но нет имени директории

	// описать ошибку когда не удалось отработать с Config.yaml

}

// GetAbsPath получает на фход строку, и функцию. В рабочем варианте, функция
// должна быть filepath.Abs(path string)(string, error). Возвращает строку -
// абсолютный путь, или ошибку
func (s *SettingsOptions) GetAbsPath(fn func(path string) (string, error)) error {
	rootAbsPath, err := fn(s.Root)
	if err != nil {
		return err
	}
	s.Root = rootAbsPath
	return nil
}

// Print выводит в командную строку настройки и путь к корневой директории
func Print(c CommandLineOptions, s SettingsOptions) {
	fmt.Println(color.YellowString("Настройки программы duplicateEraser:"))
	fmt.Println()
	fmt.Println(color.BlueString("Корневая директория:"))

	if c.M {
		fmt.Println(color.GreenString("\t-ручной ввод, директория:"), s.Root)
	} else {
		fmt.Println(color.CyanString("\t-автоматический ввод, директория:"), s.Root)
	}
	fmt.Println(color.BlueString("Режим очистки:"))
	if s.Erase {
		fmt.Println(color.RedString("\t-режим очистки включен"))
	} else {
		fmt.Println(color.CyanString("\t-режим очистки выключен"))
	}
}
