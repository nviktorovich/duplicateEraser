package LogicProcessors

import (
	"io/fs"
	"os"
	"path/filepath"
)

// AbsGetter интерфейс, включающий в себя метод abs, который прнимает на вход
// строку, возвращает строку и ошибку
type AbsGetter interface {
	abs(string) (string, error)
}

// AbsGetterCaller пустая структура, в которую будем имплиментирвоать интерфейс AbsGetter
type AbsGetterCaller struct{}

// abs метод для объекта AbsGetterCaller возвращает вызов метода сторонней
// библиотеки (filepath.Abs(string))
func (a AbsGetterCaller) abs(path string) (string, error) {
	return filepath.Abs(path)
}

// ABSVerifier интерфейс, включающий в себя метод stat, прнимает на вход строку,
// возвращает fs.FileInfo и ошибку
type ABSVerifier interface {
	stat(string) (fs.FileInfo, error)
}

// ABSVerifierCaller пустая структура для дальнейшей имплементации интерфейса
// ABSVerifier
type ABSVerifierCaller struct{}

// stat метод описанный для структуры ABSVerifierCaller, имплементирующий ее в
// интерфейс ABSVerifier
func (a ABSVerifierCaller) stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

// определение переменных, которые являются интерфейсами
var (
	AbsGC AbsGetter
	AbsVC ABSVerifier
)

// init инициализация где структуре присваивается интерфейс ??? почитать про init!!!
func init() {
	AbsGC = AbsGetterCaller{}
	AbsVC = ABSVerifierCaller{}
}

// GetAbsPath принимает строку на вход, вместо прямой реализации вызова сторонней
// библиотеки, используется вызов метода для стукруты в которую имплементирован
// интерфейс abs
func GetAbsPath(path string) (string, error) {
	absPath, err := AbsGC.abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

// VerifyAbsPath функция, принемает на вход строку, прямой вызов метода сторонней
// библиотеки заменен вызовом метода для структуры, в которую имлементирован
// интерфейс ABSVerifier
func VerifyAbsPath(s string) (bool, error) {
	_, err := AbsVC.stat(s)
	if err != nil {
		return false, err
	}
	return true, nil
}
