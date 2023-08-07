package document_test

import (
	"errors"
	"testing"

	"github.com/aqyuki/docat/internal/document"
)

func TestNewNotFindError(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		errorType *document.NotFindError
	}{
		{
			name:      "Create error",
			wantErr:   true,
			errorType: &document.NotFindError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			err := document.NewNotFindError()
			if !errors.As(err, &tt.errorType) {
				t.Errorf("NewNotFindError returned Unexpected error")
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNotFindError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotFindError_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *document.NotFindError
		want string
	}{
		{
			name: "Not Find Error Message Check",
			e:    &document.NotFindError{},
			want: "no file matching the pattern was found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("NotFindError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvalidArgument_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *document.InvalidArgument
		want string
	}{
		{
			name: "Invalid Argument Error Check",
			e:    &document.InvalidArgument{},
			want: "invalid argument",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("InvalidArgument.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewInvalidArgument(t *testing.T) {
	tests := []struct {
		name      string
		wantErr   bool
		errorType *document.InvalidArgument
	}{
		{
			name:      "Create Error",
			wantErr:   true,
			errorType: &document.InvalidArgument{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := document.NewInvalidArgument()
			if !errors.As(err, &tt.errorType) {
				t.Errorf("NewInvalidArgument returned Unexpected error")
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInvalidArgument() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
