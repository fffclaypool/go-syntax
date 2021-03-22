package main

import (
	"reflect"
	"testing"
)

func Test_creatingStrings(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{
			want: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := creatingStrings(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("creatingStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringLength(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringLength(); got != tt.want {
				t.Errorf("stringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concatenatingStrings(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			want: "Hello world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := concatenatingStrings(); got != tt.want {
				t.Errorf("concatenatingStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
