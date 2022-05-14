package InputProcessors

import (
	"testing"
)

type TestFlagOptionsImitator struct {
	testName string
	expErr   bool
	err      error
	In       *FlagOptions
}

var TestFlagOptionsImitatorSet = []TestFlagOptionsImitator{
	{"1 flag off, empty name", false, nil, &FlagOptions{false, ""}},
	{"2 flag off, not empty name", false, nil, &FlagOptions{false, "testName"}},
	{"3 flag on, empty name", true, EmptyFileNameErr, &FlagOptions{true, ""}},
	{"4 flag on, not empty name", false, nil, &FlagOptions{true, "testName"}},
}

func TestFlagOptionsAnalise(t *testing.T) {
	for _, tt := range TestFlagOptionsImitatorSet {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.In.Analise()
			if err != tt.err {
				t.Errorf("name: %s, expected err: %t, %v, getErr %v\n\n", tt.testName, tt.expErr, tt.err, err)
			}
		})
	}
}

func (c *CLFlagImitator) Set() {
	c.FO.FileName = c.a[0]
	if *c.f {
		c.FO.Mode = true
	} else {
		*c.f = false
	}

}

type CLFlagImitator struct {
	name    string
	f       *bool
	a       []string
	expFlag bool
	expName string
	FO      *FlagOptions
}

func TestMakeSet(t *testing.T) {
	var yes = true
	var no = false

	var TestCLFlagImitatorSet = []CLFlagImitator{
		{"1 flag off, name off", &no, []string{""}, false, "", &FlagOptions{}},
		{"2 flag off, name on", &no, []string{"testName"}, false, "testName", &FlagOptions{}},
		{"3 flag on, name off", &yes, []string{""}, true, "", &FlagOptions{}},
		{"4 flag on, name on", &yes, []string{"testName"}, true, "testName", &FlagOptions{}},
	}

	for _, tt := range TestCLFlagImitatorSet {
		t.Run(tt.name, func(t *testing.T) {
			MakeSet(&tt)
			if tt.FO.Mode != tt.expFlag || tt.FO.FileName != tt.expName {
				t.Errorf("exp:\n\tMode:\t%t\n\tName:\t%s\nget:\n\tMode:\t%t\n\tName:\t%s\n", tt.expFlag, tt.expName, tt.FO.Mode, tt.FO.FileName)
			}
		})
	}
}
