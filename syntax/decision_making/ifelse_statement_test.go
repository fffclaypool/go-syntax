package main

import "testing"

func Test_basicOperationX(t *testing.T) {
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
			if got := basicOperationX(); got != tt.want {
				t.Errorf("basicOperationX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_basicOperationY(t *testing.T) {
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
			if got := basicOperationY(); got != tt.want {
				t.Errorf("basicOperationY() = %v, want %v", got, tt.want)
			}
		})
	}
}
