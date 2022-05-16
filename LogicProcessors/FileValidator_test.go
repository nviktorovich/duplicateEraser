package LogicProcessors

import (
	"io/fs"
	"testing"
)

var AbsGetMockFunc func(path string) (string, error)

type AbsGetMockStruct struct{}

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

var AbsVerifyMockFunc func(name string) (fs.FileInfo, error)

type AbsVerifyMockStruct struct{}

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
