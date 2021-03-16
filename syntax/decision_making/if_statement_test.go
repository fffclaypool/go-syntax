package main

import "testing"

func Test_ifstatement(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ifstatement(); got != tt.want {
				t.Errorf("ifstatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
