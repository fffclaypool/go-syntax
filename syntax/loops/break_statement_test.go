package main

import "testing"

func Test_breakstatementBasicOperation(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := breakstatementBasicOperation(); got != tt.want {
				t.Errorf("breakstatementBasicOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
