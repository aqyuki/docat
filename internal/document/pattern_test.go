package document_test

import (
	"reflect"
	"testing"

	"github.com/aqyuki/docat/internal/document"
)

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
