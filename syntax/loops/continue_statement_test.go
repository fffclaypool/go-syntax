package main

import "testing"

func Test_continuestatementBasicOperation(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := continuestatementBasicOperation(); got != tt.want {
				t.Errorf("continuestatementBasicOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
