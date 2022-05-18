package LogicProcessors

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

// Walker интерфейс, реализующий метод walk (подготовка для создания заглушки, для метода filepath.Walk()
type Walker interface {
	walk(root string, fn func(path string, info fs.FileInfo, err error) error) error
}

// WalkStruct пустая структура для последующей имплементации в интерфейс Walker
type WalkStruct struct{}

// описание метода walk для структуры WalkStruct
func (w WalkStruct) walk(root string, fn func(path string, info fs.FileInfo, err error) error) error {
	return filepath.Walk(root, fn)
}

// WalkVar объявление переменной, для последующей передачи внутрь структуры
// WalkStruct, т.о переменная будет содержать структуру, которой присущь интерфес
// Walker
var WalkVar Walker

// инициализация, передача структуры в переменную
func init() {
	WalkVar = WalkStruct{}
}

// GetFileList обрабатывает рекурсивно указанную директорию.
func GetFileList(path string) error {
	err := WalkVar.walk(path,
		func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() {
				fmt.Printf("filename is: %s, file size is: %d byte(s)\n", info.Name(), info.Size())
				fmt.Println(filepath.Abs(info.Name()))
				// вставить код, отвечающий за формирование словаря
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

// находим все файлы и формируем список
// фильтруем список, так, чтобы оставались только файлы, которые дублируют другие
// признак дубля - одинаковое имя, одинаковое содержание, имя и содержание
// удалям все дубликаты, кроме одного файла
