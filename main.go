package main

import (
	"github.com/nviktorovich/duplicateEraser/InputProcessors"
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

}
