package options_test

import (
	"reflect"
	"testing"

	"github.com/aqyuki/docat/internal/scanner/options"
)

func TestCreateDefaultOption(t *testing.T) {
	tests := []struct {
		name string
		want *options.ScannerOption
	}{
		{
			name: "Normal",
			want: &options.ScannerOption{
				IgnoreSubDirectory: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := options.CreateDefaultOption(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateDefaultOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithIgnoreSubDirectory(t *testing.T) {
	type args struct {
		opts *options.ScannerOption
	}
	tests := []struct {
		name string
		args args
		want *options.ScannerOption
	}{
		{
			name: "Normal",
			args: args{
				opts: &options.ScannerOption{
					IgnoreSubDirectory: false,
				},
			},
			want: &options.ScannerOption{
				IgnoreSubDirectory: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options.WithIgnoreSubDirectory(tt.args.opts)
			if !reflect.DeepEqual(tt.args.opts, tt.want) {
				t.Errorf("%s : got -> %+v want -> %+v", tt.name, tt.args.opts, tt.want)
			}
		})
	}
}
