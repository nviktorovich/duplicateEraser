package InputProcessors

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
)

var EmptyFileNameErr = errors.New("ошибка имени файла, название не може быть пустым")

type Analyser interface {
	Analise() error
}

type Setter interface {
	Set()
}

type FlagOptions struct {
	Mode     bool
	Eraser   bool
	FileName string
}

func (f *FlagOptions) Analise() error {

	fmt.Printf(color.YellowString("\nSELECTED OPTIONS:\n"))
	// Option name -m
	fmt.Printf(color.BlueString("Режим ввода пути корневой директории (-m):\n"))
	if f.Mode {
		fmt.Println(color.GreenString("\t- ручной ввод корневой директории"))
		if f.FileName == "" {
			return EmptyFileNameErr
		}
	} else {
		fmt.Println(color.CyanString("\t- корневая директороия по умолчанию"))
	}

	// Option name -d
	fmt.Printf(color.BlueString("Режим очистки (-d):\n"))
	if f.Eraser {
		fmt.Println(color.RedString("\t- режим очистки влючен, дубликаты будут удалены"))
	} else {
		fmt.Println(color.CyanString("\t- режим очистки выключен"))
	}

	fmt.Println()

	return nil

}

var (
	FlagInputSelector  bool
	FlagEraserSelector bool
)

func (f *FlagOptions) Set() {

	flag.BoolVar(&FlagInputSelector, "m", false, "выбор режима ввода имени корневой директории: наличие флага - ручной; отстутвие - по умолчанию")
	flag.BoolVar(&FlagEraserSelector, "d", false, "выбор режима удаления дубликатов: наличие флага - удаление; отстутвие - режим просмотра")
	flag.Parse()
	f.Mode = FlagInputSelector
	f.Eraser = FlagEraserSelector
	f.FileName = flag.Arg(0)
}

func MakeSet(s Setter) {
	s.Set()
}
