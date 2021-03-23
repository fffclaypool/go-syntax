package main

import (
	"reflect"
	"testing"
)

func Test_arraysofpointers(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arraysofpointers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arraysofpointers() = %v, want %v", got, tt.want)
			}
		})
	}
}
