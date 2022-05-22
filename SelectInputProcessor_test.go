package main

import (
	"github.com/nviktorovich/duplicateEraser/InputProcessors"
	"os"
	"testing"
)

type TestingStructOptions struct {
	name      string
	inputData InputProcessors.CommandLineOptions
	fn        func() (string, error)
	expData   InputProcessors.SettingsOptions
	expErr    bool
}

// TestNewSettingsOptions в оригинале передаем функцию, которая открывает
// конфигурационный файл, в этом нет резона. Заменил заглушкой
func TestNewSettingsOptions(t *testing.T) {
	var TestSet = []TestingStructOptions{
		{
			"1. m, d, args: 0\t0\t0",
			InputProcessors.CommandLineOptions{false, false, []string{}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			false,
		},
		{
			"2. m, d, args:0\t0\t1",
			InputProcessors.CommandLineOptions{false, false, []string{"test"}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			false,
		},
		{
			"3. m, d, args: 0\t1\t0",
			InputProcessors.CommandLineOptions{false, true, []string{}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{true, "default"},
			false,
		}, {
			"4. m, d, args: 0\t1\t1",
			InputProcessors.CommandLineOptions{false, true, []string{"test"}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{true, "default"},
			false,
		}, {
			"5. m, d, args: 1\t0\t0",
			InputProcessors.CommandLineOptions{true, false, []string{}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			true,
		}, {
			"6. m, d, args: 1\t0\t1",
			InputProcessors.CommandLineOptions{false, false, []string{"default"}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			false,
		}, {
			"7. m, d, args: 1\t1\t0",
			InputProcessors.CommandLineOptions{true, true, []string{}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{true, "default"},
			true,
		}, {
			"8. m, d, args: 1\t1\t1",
			InputProcessors.CommandLineOptions{true, true, []string{"default"}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{true, "default"},
			false,
		}, {
			"9. m, d, args: 1\t0\t1",
			InputProcessors.CommandLineOptions{true, false, []string{""}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			true,
		}, {
			"10. m, d, args: 1\t1\t1",
			InputProcessors.CommandLineOptions{true, true, []string{""}},
			func() (string, error) { return "default", nil },
			InputProcessors.SettingsOptions{false, "default"},
			true,
		},
	}
	for _, tt := range TestSet {
		t.Run(tt.name, func(t *testing.T) {
			tData, tErr := InputProcessors.NewSettingsOptions(tt.inputData, tt.fn)
			switch tt.expErr {
			case true:
				if tErr == nil {
					t.Errorf("Ожидалась ошибка, но не была получена: %s, \t%v", tt.name, tt.inputData)
				}
			case false:
				if tErr != nil {
					t.Errorf("Возникла ошибка на тесте, где не должна была: %s, \t%v", tt.name, tt.inputData)
				}
				if tData != tt.expData {
					t.Errorf("Тест: %s, ожидался набор: %v, получен набор: %v", tt.name, tt.expData, tData)
				}
			}

		})
	}

}

type TestingStructAbs struct {
	name   string
	fn     func(path string) (string, error)
	in     InputProcessors.SettingsOptions
	exp    InputProcessors.SettingsOptions
	erxErr bool
}

// AbsImitator замена функции Abs для локального тестирования.
func AbsImitator(path string) (string, error) {
	fileSep := os.PathSeparator
	fistSymbol := rune(path[0])
	if fistSymbol == fileSep {
		return path, nil
	}
	return "/Test/" + path, nil
}

func TestGetAbsPath(t *testing.T) {
	var TestSet = []TestingStructAbs{
		{
			"1. with separator in start",
			AbsImitator,
			InputProcessors.SettingsOptions{false, "/Test/12"},
			InputProcessors.SettingsOptions{false, "/Test/12"},
			false,
		},
		{
			"2. without separator in start",
			AbsImitator,
			InputProcessors.SettingsOptions{false, "12"},
			InputProcessors.SettingsOptions{false, "/Test/12"},
			false,
		},
	}
	for _, tt := range TestSet {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.in.GetAbsPath(tt.fn)
			if err != nil || tt.in != tt.exp {
				t.Errorf("Тест: %s, ожидалась ошибка %t, ожидался набор: %v, получена ошибка: %v, получен набор: %v",
					tt.name, tt.erxErr, tt.exp, err, tt.in)
			}
		})
	}

}
