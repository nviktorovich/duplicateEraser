package main

import (
	"github.com/nviktorovich/duplicateEraser/InputProcessors"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

func TestFlagOptions_NameSet(t *testing.T) {
	type TestStruct struct {
		Name    string
		FO      *InputProcessors.FlagOptions
		expName string
		expErr  bool
	}

	file, err := os.ReadFile(InputProcessors.ConfigPath)
	if err != nil {
		t.Errorf("unavaible: %s", InputProcessors.ConfigPath)
	}

	n := new(InputProcessors.DefaultSettings)
	if err = yaml.Unmarshal(file, n); err != nil {
		t.Errorf("unmarshal problem: %s", InputProcessors.ConfigPath)
	}

	var TestSet = []TestStruct{
		{"flag and name", &InputProcessors.FlagOptions{Mode: true, FileName: "testname"}, "testname", false},
		{"flag and name 2", &InputProcessors.FlagOptions{Mode: true, FileName: "testname2"}, "testname2", false},
		{"no flag and name", &InputProcessors.FlagOptions{Mode: false, FileName: "testname3"}, n.PathName, false},
		{"no flag and no name", &InputProcessors.FlagOptions{Mode: false, FileName: ""}, n.PathName, false},
	}

	for _, tt := range TestSet {
		t.Run(tt.Name, func(t *testing.T) {
			err = tt.FO.NameSet()
			if tt.expName != tt.FO.FileName {
				t.Errorf("expected: %s, get: %s", tt.expName, tt.FO.FileName)
			}
			if err != nil && tt.expErr {
				t.Errorf("expected error: %t, get error: %s", tt.expErr, err)
			}
		})
	}
}
