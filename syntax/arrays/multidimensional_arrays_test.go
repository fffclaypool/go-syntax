package main

import (
	"reflect"
	"testing"
)

func Test_accessingTwodimensionalArrayElements(t *testing.T) {
	tests := []struct {
		name string
		want [5][2]int
	}{
		{
			want: [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := accessingTwodimensionalArrayElements(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("accessingTwodimensionalArrayElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
