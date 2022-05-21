package LogicProcessors

import (
	"io/fs"
	"reflect"
	"testing"
)

var walkMockFunc func(root string, fn func(path string, info fs.FileInfo, err error) error) error

// WalkMockStruct пустая структура для последующей имплементации в интерфейс Walker
type WalkMockStruct struct {
}

func (w WalkMockStruct) walk(root string, fn func(path string, info fs.FileInfo, err error) error) error {
	return walkMockFunc(root, fn)
}

type MockData struct {
	name   string
	expErr bool
	path   string
	m      map[string][]string
}

func TestGetFileList(t *testing.T) {
	var TestSet = []MockData{
		{"1 test set", false, "/rootTest", map[string][]string{"1.txt_1": {"/rootTest/1/1.txt", "/rootTest/2/1.txt"}}},
		{"2 test set", false, "/newTest", map[string][]string{"1.txt_1": {"/newTest/1/1.txt", "/newTest/2/1.txt"}}},
	}
	WalkVar = WalkMockStruct{}
	walkMockFunc = func(root string, fn func(path string, info fs.FileInfo, err error) error) error {
		for _, v := range []string{"1", "2"} {
			m["1.txt_1"] = append(m["1.txt_1"], root+"/"+v+"/1.txt")
		}
		return nil
	}

	for _, tt := range TestSet {
		t.Run(tt.name, func(t *testing.T) {

			testMap, testErr := GetFileList(tt.path)
			if testErr != nil && !tt.expErr {
				t.Errorf("test: %s, expErr: %t, getErr: %s", t.Name(), tt.expErr, testErr)
			}
			for k, v := range testMap {
				if !reflect.DeepEqual(tt.m[k], v) {
					t.Errorf("test:  %s\n\nexpMap: %v\n\ngetMap: %v", tt.name, tt.m, testMap)
				}
			}
			m = map[string][]string{}

		})
	}
}
