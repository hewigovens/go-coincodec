package coincodec

import (
	"encoding/hex"
	"reflect"
	"strings"
	"testing"
)

func TestStellarDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "Normal",
			input:  "GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER",
			output: "11b32750d877d03b29df85c0349b3ca5c6abf64786d773323c417f34df0b2fea",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StellarDecodeToBytes(tt.input)
			if tt.err != nil {
				if !strings.HasPrefix(err.Error(), tt.err.Error()) {
					t.Errorf("StellarDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("StellarDecodeToBytes() = %v, want %v", hex.EncodeToString(got), tt.output)
				}
			}
		})
	}
}

func TestStellarEncodeToString(t *testing.T) {
	pubkey, _ := hex.DecodeString("11b32750d877d03b29df85c0349b3ca5c6abf64786d773323c417f34df0b2fea")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "Good",
			input:  pubkey,
			output: "GAI3GJ2Q3B35AOZJ36C4ANE3HSS4NK7WI6DNO4ZSHRAX6NG7BMX6VJER",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StellarEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("StellarEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("StellarEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
