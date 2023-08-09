package printer_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/aqyuki/docat/internal/printer"
)

func PickStdOut(t *testing.T, fn func()) string {
	t.Helper()
	backup := os.Stdout
	defer func() {
		os.Stdout = backup
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	var buffer bytes.Buffer
	if _, err := buffer.ReadFrom(r); err != nil {
		t.Fatalf("fail read buf: %v", err)
	}
	s := buffer.String()
	return s[:len(s)-1]
}

func TestShowSimpleList(t *testing.T) {
	type args struct {
		list []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				list: []string{
					"sample",
					"sample",
				},
			},
			want: "01 : sample\n02 : sample",
		},
		{
			name: "no detected <nil>",
			args: args{list: nil},
			want: "item no detected",
		},
		{
			name: "no detected <0>",
			args: args{list: []string{}},
			want: "item no detected",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PickStdOut(t, func() { printer.ShowSimpleList(tt.args.list) })
			if got != tt.want {
				t.Errorf("Unexpected output returned want -> %+v got -> %+v\n", tt.want, got)
			}
		})
	}
}
