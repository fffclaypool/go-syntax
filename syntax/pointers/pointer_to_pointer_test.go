package main

import "testing"

func Test_pointertopointer(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 3000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointertopointer(); got != tt.want {
				t.Errorf("pointertopointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
