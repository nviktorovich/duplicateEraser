package main

import (
	"fmt"
	"github.com/nviktorovich/duplicateEraser/InputSettings"
	"log"
)

func main() {
	val, err := InputSettings.AnaliseOptions()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(val)
}
