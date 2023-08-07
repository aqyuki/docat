package document_test

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aqyuki/docat/internal/document"
)

func GenerateFilePath(t *testing.T, filename string) string {
	t.Helper()
	return filepath.Join("project", filename)
}

func TestCreatePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want *document.DocumentPat
	}{
		{
			name: "Create new object",
			args: args{
				pattern: "sample",
			},
			want: &document.DocumentPat{
				Pattern: "sample",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := document.CreatePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNormalArguments(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error case 1",
			args: args{
				path: "",
			},
			wantErr: true,
		},
		{
			name: "Normal",
			args: args{
				path: filepath.Join("sample/sample.txt"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := document.IsNormalArguments(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("isNormalArguments() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_compareFileName(t *testing.T) {
	baseName := "README"
	type args struct {
		base string
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same name",
			args: args{
				base: baseName,
				name: "README",
			},
			want: true,
		},
		{
			name: "Same lowercase ",
			args: args{
				base: baseName,
				name: "readme",
			},
			want: true,
		},
		{
			name: "mixed case ",
			args: args{
				base: baseName,
				name: "ReadMe",
			},
			want: true,
		}, {
			name: "Same name with extension",
			args: args{
				base: baseName,
				name: "README.md",
			},
			want: true,
		},
		{
			name: "Same lowercase with extension",
			args: args{
				base: baseName,
				name: "readme.md",
			},
			want: true,
		},
		{
			name: "mixed case with extension",
			args: args{
				base: baseName,
				name: "ReadMe.md",
			},
			want: true,
		},
		{
			name: "vary name",
			args: args{
				base: baseName,
				name: "CHANGELOG",
			},
			want: false,
		},
		{
			name: "vary lowercase ",
			args: args{
				base: baseName,
				name: "changelog",
			},
			want: false,
		},
		{
			name: "mixed vary case ",
			args: args{
				base: baseName,
				name: "ChangeLog",
			},
			want: false,
		}, {
			name: "vary name with extension",
			args: args{
				base: baseName,
				name: "CHANGELOG.md",
			},
			want: false,
		},
		{
			name: "vary lowercase with extension",
			args: args{
				base: baseName,
				name: "changelog.md",
			},
			want: false,
		},
		{
			name: "mixed vary case with extension",
			args: args{
				base: baseName,
				name: "ChangeLog.md",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := document.CompareFileName(tt.args.base, tt.args.name); got != tt.want {
				t.Errorf("Case %v got = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestIsREADME(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "README.md"),
			},
			want: true,
		},
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "README"),
			},
			want: true,
		},
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "CHANGELOG.md"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := document.IsREADME(tt.args.path); got != tt.want {
				t.Errorf("IsREADME() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLICENSE(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "LICENSE.md"),
			},
			want: true,
		},
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "LICENSE"),
			},
			want: true,
		},
		{
			name: "Same case",
			args: args{
				path: GenerateFilePath(t, "README.md"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := document.IsLICENSE(tt.args.path); got != tt.want {
				t.Errorf("IsLICENSE() = %v, want %v", got, tt.want)
			}
		})
	}
}
