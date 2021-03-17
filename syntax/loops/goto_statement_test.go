package main

import "testing"

func Test_gotostatementBasicOperation(t *testing.T) {
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
			if got := gotostatementBasicOperation(); got != tt.want {
				t.Errorf("gotostatementBasicOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}
