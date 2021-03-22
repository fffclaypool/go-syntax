package main

import "testing"

func Test_basicFunctionasvalues(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := basicFunctionasvalues(); got != tt.want {
				t.Errorf("basicFunctionasvalues() = %v, want %v", got, tt.want)
			}
		})
	}
}
