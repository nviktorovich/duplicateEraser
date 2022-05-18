package main

import (
	"github.com/nviktorovich/duplicateEraser/InputProcessors"
	"github.com/nviktorovich/duplicateEraser/LogicProcessors"
	"log"
)

func main() {
	a := new(InputProcessors.FlagOptions)
	InputProcessors.MakeSet(a)

	if err := a.Analise(); err != nil {
		log.Fatal(err)
	}
	if err := a.NameSet(); err != nil {
		log.Fatal(err)
	}
	absPath, err := LogicProcessors.GetAbsPath(a.FileName)
	if err != nil {
		log.Fatal(err)
	}

	ok, err := LogicProcessors.VerifyAbsPath(absPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("file path: %s is exist: %t", absPath, ok)

	err = LogicProcessors.GetFileList(absPath)
	if err != nil {
		log.Fatal(err)
	}
}
