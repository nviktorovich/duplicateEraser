package LogicProcessors

import (
	"io/fs"
	"path/filepath"
	"strconv"
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

type DirPoolStruct struct {
	M map[string][]string
}

func NewDirPoolStruct(m map[string][]string) *DirPoolStruct {
	return &DirPoolStruct{m}

}

func (m *DirPoolStruct) Filter() {
	for k, v := range m.M {
		if len(v) < 2 {
			delete(m.M, k)
		}
	}
}

// GetKeyName функция, которая принимает на вход строку и число в формате int64 и
// возврашает строку, где входные аргументы соединены через "_"
func GetKeyName(name string, size int64) string {
	strSize := strconv.FormatInt(size, 10)
	return name + "_" + strSize
}

// GetFileList обрабатывает рекурсивно указанную директорию. Формирует мапу.
// дабавить тестирование всего пакета необходимо!!!
func GetFileList(root string) (map[string][]string, error) {
	m := make(map[string][]string)

	if err := WalkVar.walk(root, func(path string, info fs.FileInfo, err error) error {
		key := GetKeyName(info.Name(), info.Size())
		val := path
		if err != nil {
			return err
		}
		m[key] = append(m[key], val)
		return nil
	}); err != nil {
		return m, err
	}
	return m, nil
}

// находим все файлы и формируем список
// фильтруем список, так, чтобы оставались только файлы, которые дублируют другие
// признак дубля - одинаковое имя, одинаковое содержание, имя и содержание
// удалям все дубликаты, кроме одного файла
