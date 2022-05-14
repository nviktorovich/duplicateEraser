package InputProcessors

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
)

var FlagInputModeSelector bool

var EmpyFileNameErr = errors.New("ошибка имени файла, название не може быть пустым")

type Analiser interface {
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
		fmt.Println("указана опция -m, выбран", color.GreenString("ручной режим"))
		if f.FileName != "" {
			fmt.Println("указано имя файла:", color.GreenString(f.FileName))
			return nil
		}
		return EmpyFileNameErr
	}
	fmt.Println("отсутствует опция -m, выбран", color.YellowString("режим по умолчанию"))
	fmt.Println("указано имя файла:", color.YellowString(f.FileName))
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
