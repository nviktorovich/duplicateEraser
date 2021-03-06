package main

import (
	"github.com/nviktorovich/duplicateEraser/EraserProcessor"
	"github.com/nviktorovich/duplicateEraser/InputProcessors"
	"github.com/nviktorovich/duplicateEraser/LogicProcessors"
	"log"
	"path/filepath"
)

func main() {

	CommandLineParameters := InputProcessors.NewCommandLineOptions()

	SettingsParameters, err := InputProcessors.NewSettingsOptions(CommandLineParameters, InputProcessors.GetDefaultRootName)
	if err != nil {
		log.Fatal(err)
	}

	if err = SettingsParameters.GetAbsPath(filepath.Abs); err != nil {
		log.Fatal(err)
	}

	InputProcessors.Print(CommandLineParameters, SettingsParameters)

	if err = LogicProcessors.Validate(SettingsParameters.Root); err != nil {
		log.Fatal(err)
	}

	RootMap, err := LogicProcessors.NewRootMap(SettingsParameters.Root)
	if err != nil {
		log.Fatal(err)
	}
	RootMap.Filter()

	for _, v := range RootMap {
		if err = EraserProcessor.EraserOperator(v); err != nil {
			log.Fatal(err)
		}
	}

}
