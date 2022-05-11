package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nviktorovich/duplicateEraser/DriverAndSettings"
	"github.com/nviktorovich/duplicateEraser/InputSettings"
	"log"
	"path/filepath"
)

func main() {
	val, err := InputSettings.AnaliseOptions()
	if err != nil {
		log.Fatal(err)
	}

	isValid, err := DriverAndSettings.FileNameValidate(val)
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(val)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("file name is: ", color.GreenString(val))
	if isValid {
		fmt.Println("file name is valid: ", color.GreenString("существует"))
		fmt.Println("full path is : ", color.GreenString(path))

	} else {
		fmt.Println("file name is valid: ", color.RedString("не существует"))
	}

}
