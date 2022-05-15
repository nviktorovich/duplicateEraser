package LogicProcessors

import (
	"io/fs"
	"os"
	"path/filepath"
)

type AbsGetter interface {
	abs(string) (string, error)
}

type AbsGetterCaller struct{}

func (a AbsGetterCaller) abs(path string) (string, error) {
	return filepath.Abs(path)
}

type ABSVerifier interface {
	stat(string) (fs.FileInfo, error)
}

type ABSVerifierCaller struct{}

func (a ABSVerifierCaller) stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}

var AbsGC AbsGetter
var AbsVC ABSVerifier

func init() {
	AbsGC = AbsGetterCaller{}
	AbsVC = ABSVerifierCaller{}
}

func GetAbsPath(path string) (string, error) {
	absPath, err := AbsGC.abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func VerifyAbsPath(s string) (bool, error) {
	_, err := AbsVC.stat(s)
	if err != nil {
		return false, err
	}
	return true, nil
}
