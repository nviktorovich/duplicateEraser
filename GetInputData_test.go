package main

import (
	"github.com/nviktorovich/duplicateEraser/InputSettings"
	"testing"
)

func TestGetData(t *testing.T) {
	var yes, no = new(bool), new(bool)
	*yes = true
	*no = false
	type args struct {
		b *bool
		s []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"correct test with command line options 1", args{yes, []string{"new"}}, "new", false},
		{"correct test with command line options 2", args{yes, []string{"./"}}, "./", false},
		{"incorrect test with command line options 1", args{yes, []string{}}, "", true},
		{"correct test without command line options 1", args{no, []string{}}, ".", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InputSettings.GetData(tt.args.b, tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetData() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadDefault(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"correct test 1", ".", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InputSettings.ReadDefault()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadDefault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadDefault() got = %v, want %v", got, tt.want)
			}
		})
	}
}
