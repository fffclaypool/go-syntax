package main

import "testing"

func Test_constants(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constants(); got != tt.want {
				t.Errorf("constants() = %v, want %v", got, tt.want)
			}
		})
	}
}
