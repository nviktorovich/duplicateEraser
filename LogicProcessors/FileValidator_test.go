package LogicProcessors

import (
	"io/fs"
	"testing"
)

// AbsGetMockFunc объявляем переменную, которая является функцией, аргументы и
// возврат которой, соответствуют методу abs
var AbsGetMockFunc func(path string) (string, error)

// AbsGetMockStruct пустая структура для имплементации интерфейса AbsGetter
type AbsGetMockStruct struct{}

// abs реализация метода для структуры AbsGetMockStruct
func (a AbsGetMockStruct) abs(path string) (string, error) {
	return AbsGetMockFunc(path)
}

func TestGetAbsPath(t *testing.T) {
	testPath := "/users"
	AbsGC = AbsGetMockStruct{}
	AbsGetMockFunc = func(path string) (string, error) {
		return "/users", nil
	}

	t.Run("simple test", func(t *testing.T) {
		p, err := GetAbsPath(testPath)
		if err != nil {
			t.Errorf("not exp err, but get: %s", err)
		}
		if p != testPath {
			t.Errorf("exp AbsPath: %s, get path: %s", testPath, p)
		}
	})

}

// testing Verify func

// AbsVerifyMockFunc объявляем переменную, которая является функцией, аргументы и
// возврат которой, соответствуют методу stat
var AbsVerifyMockFunc func(name string) (fs.FileInfo, error)

// AbsVerifyMockStruct пустая структура для имплементации интерфейса ABSVerifier
type AbsVerifyMockStruct struct{}

// stat реализация метода для заглушки
func (a AbsVerifyMockStruct) stat(name string) (fs.FileInfo, error) {
	return AbsVerifyMockFunc(name)
}

func TestVerifyAbsPath(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		path := "/users"
		const expOk = true

		AbsVC = AbsVerifyMockStruct{}
		AbsVerifyMockFunc = func(name string) (fs.FileInfo, error) {
			return nil, nil
		}

		ok, err := VerifyAbsPath(path)
		if err != nil {
			t.Errorf("expectef nil err, but get: %s", err)
		}
		if ok != expOk {
			t.Errorf("expected ok is %t, but get: %t", expOk, ok)
		}

	})
}
