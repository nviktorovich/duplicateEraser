package LogicProcessors

import (
	"io/fs"
	"path/filepath"
	"strconv"
)

// DuplicateFileInfo объединяем имя файла и размер
type DuplicateFileInfo struct {
	UniteFileNameAndFileSize string
	FilePaths                []string
}

type Walker interface {
	walk(root string, fn func(path string, info fs.FileInfo, err error) error) error
}

type WorkAbs struct{}

func (w WorkAbs) walk(root string, fn func(path string, info fs.FileInfo, err error) error) error {
	return filepath.Walk(root, fn)
}

var WalkProc Walker

func init() {
	WalkProc = WorkAbs{}
}

type RootMap map[string][]string

// NewRootMap возвращает объект RootMap, который содержит полную карту в указанной корневой директории
func NewRootMap(path string) (RootMap, error) {
	r := make(RootMap)
	if err := WalkProc.walk(path, func(path string, info fs.FileInfo, err error) error {
		name := info.Name() + "_" + strconv.FormatInt(info.Size(), 10)
		r[name] = append(r[name], path)
		return nil
	}); err != nil {
		return r, err
	}
	return r, nil
}

// Filter метод структуры RootMap отфильтровывающий дубликаты, в карте остаются только дубли
func (m RootMap) Filter() {
	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
		} else {
			m[k] = v[1:]
		}
	}
}
