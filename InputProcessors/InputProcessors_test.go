package InputProcessors

import (
	"flag"
	"testing"
)

// Analise() возвращает ошибку, только в одном случае, если выбрана опция ручного ввода -m и отстутсвует имя файла
// необходимо проверить комбинацию флагов в сочетании с именем файла

type InputOptionsAnaliseData struct {
	tName  string
	flags  FlagOptions
	expErr bool
}

func TestFlagOptions_Analise(t *testing.T) {
	var TestSet = []InputOptionsAnaliseData{
		{
			"1. Все опции выключены, имя отсутствует: ожидается отсутствие ошибки",
			FlagOptions{false, false, ""},
			false,
		},
		{
			"2. Все опции выключены, имя присутствует: ожидается отсутствие ошибки",
			FlagOptions{false, false, "/Users"},
			false,
		},
		{
			"3. -m вкл; -d выкл, имя отсутствует: ожидается ошибка",
			FlagOptions{true, false, ""},
			true,
		},
		{
			"4. -m вкл; -d выкл, имя присутствкет: ожидается отсутствие ошибки",
			FlagOptions{true, false, "/Users"},
			false,
		},

		{
			"5. -m выкл; -d вкл, имя отсутствует: ожидается отсутствие ошибки",
			FlagOptions{false, true, ""},
			false,
		},
		{
			"6. -m выкл; -d вкл, имя присутствкет: ожидается отсутствие ошибки",
			FlagOptions{false, true, "/Users"},
			false,
		},

		{
			"7. -m вкл; -d вкл, имя отсутствует: ожидается ошибка",
			FlagOptions{true, true, ""},
			true,
		},
		{
			"8. -m вкл; -d вкл, имя присутствкет: ожидается отсутствие ошибки",
			FlagOptions{true, true, "/Users"},
			false,
		},
	}

	for _, tt := range TestSet {
		t.Run(tt.tName, func(t *testing.T) {
			tErr := tt.flags.Analise()
			if tt.expErr && tErr == nil {
				t.Errorf("\tОжидается ошибка: %t, но она отстутсвует. Набор данных %v", tt.expErr, tt.flags)
			} else if !tt.expErr && tErr != nil {
				t.Errorf("\tПолучена ошибка: %s. Набор данных %v", tErr, tt.flags)
			}
		})
	}
}

// FlagsAndArgsComb замена флагов и аргументов командной строки.
type FlagsAndArgsComb struct {
	mode  *bool
	erase *bool
	a     []string
}

type SetOptionsData struct {
	tName string
	FO    FlagsAndArgsComb
	expFO FlagOptions
}

// TestFlagOptions_Set не получается, снова вернулся к тому, что нужно создавать
// иммитацию флагов и аргументов, но по сути, это не позволит протестировать
// метод Set()
func TestFlagOptions_Set(t *testing.T) {

	//var (
	//	mOn  = flag.Bool("m", true, "")
	//	mOff = flag.Bool("m", false, "")
	//	dOn  = flag.Bool("d", true, "")
	//	dOff = flag.Bool("d", false, "")
	//)

	var TestSet = []SetOptionsData{
		{
			"m, d, a: 0\t0\t0",
			FlagsAndArgsComb{
				flag.Bool("m", false, ""),
				flag.Bool("d", false, ""),
				[]string{""},
			},
			FlagOptions{false, false, ""},
		},
		{
			"m, d, a: 0\t0\t1",
			FlagsAndArgsComb{
				flag.Bool("m", false, ""),
				flag.Bool("d", false, ""),
				[]string{""},
			},
			FlagOptions{false, false, "/Users"},
		}, {
			"m, d, a: 0\t1\t0",
			FlagsAndArgsComb{
				flag.Bool("m", false, ""),
				flag.Bool("d", true, ""),
				[]string{""},
			},
			FlagOptions{false, true, ""},
		}, {
			"m, d, a: 0\t1\t1",
			FlagsAndArgsComb{
				flag.Bool("m", false, ""),
				flag.Bool("d", true, ""),
				[]string{""},
			},
			FlagOptions{false, true, "/Users"},
		}, {
			"m, d, a: 1\t0\t0",
			FlagsAndArgsComb{
				flag.Bool("m", true, ""),
				flag.Bool("d", false, ""),
				[]string{""},
			},
			FlagOptions{true, false, ""},
		}, {
			"m, d, a: 1\t0\t1",
			FlagsAndArgsComb{
				flag.Bool("m", true, ""),
				flag.Bool("d", false, ""),
				[]string{""},
			},
			FlagOptions{true, false, "/Users"},
		}, {
			"m, d, a: 1\t1\t0",
			FlagsAndArgsComb{
				flag.Bool("m", true, ""),
				flag.Bool("d", true, ""),
				[]string{""},
			},
			FlagOptions{true, true, ""},
		}, {
			"m, d, a: 1\t1\t1",
			FlagsAndArgsComb{
				flag.Bool("m", true, ""),
				flag.Bool("d", true, ""),
				[]string{""},
			},
			FlagOptions{true, true, "/Users"},
		},
	}
	for _, tt := range TestSet {
		t.Run(tt.tName, func(t *testing.T) {
			tF := new(FlagOptions)
			FlagInputSelector = *tt.FO.mode
			FlagEraserSelector = *tt.FO.erase
			tF.Set()
			if tF.Mode != tt.expFO.Mode || tF.Eraser != tt.expFO.Eraser || tF.FileName != tt.expFO.FileName {
				t.Errorf("ожидалось: %v; получено: %v", tt.expFO, tF)
			}
		})
	}

}
