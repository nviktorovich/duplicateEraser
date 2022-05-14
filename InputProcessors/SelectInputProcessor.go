package InputProcessors

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
)

var FlagInputModeSelector bool

var EmptyFileNameErr = errors.New("ошибка имени файла, название не може быть пустым")

type Analyser interface {
	Analise() error
}

type Setter interface {
	Set()
}

type FlagOptions struct {
	Mode     bool
	FileName string
}

func (f *FlagOptions) Analise() error {
	if f.Mode {
		fmt.Println(color.GreenString("ручной режим"))
		if f.FileName == "" {
			return EmptyFileNameErr
		}
		return nil
	}
	fmt.Println(color.CyanString("режим по умолчанию"))
	return nil
}

func (f *FlagOptions) Set() {
	flag.BoolVar(&FlagInputModeSelector, "m", false, "выбор режима: наличие флага - ручной; отстутвие - по умолчанию")
	flag.Parse()
	f.Mode = FlagInputModeSelector
	f.FileName = flag.Arg(0)
}

func MakeSet(s Setter) {
	s.Set()
}
