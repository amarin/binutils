package binutils_test

import (
	"encoding"
	"testing"

	. "github.com/amarin/binutils"
)

func TestLoadBinary(t *testing.T) {
	for _, tt := range []struct {
		name    string
		dict    BufferUnmarshaler
		wantErr bool
	}{} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt again
			filename := ""
			if err := LoadBinary(filename, tt.dict); (err != nil) != tt.wantErr {
				t.Errorf("LoadBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveBinary(t *testing.T) {
	for _, tt := range []struct {
		name     string
		filename string
		dict     encoding.BinaryMarshaler
		wantErr  bool
	}{} {
		tt := tt // pin tt
		t.Run(tt.name, func(t *testing.T) {
			tt := tt // pin tt again
			if err := SaveBinary(tt.filename, tt.dict); (err != nil) != tt.wantErr {
				t.Errorf("SaveBinary() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
