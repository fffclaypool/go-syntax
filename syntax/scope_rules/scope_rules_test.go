package main

import "testing"

func Test_localVariables(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := localVariables(); got != tt.want {
				t.Errorf("localVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_globalVariables(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			want:  30,
			want1: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := globalVariables()
			if got != tt.want {
				t.Errorf("globalVariables() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("globalVariables() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_formalParameters(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 30,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formalParameters(); got != tt.want {
				t.Errorf("formalParameters() = %v, want %v", got, tt.want)
			}
		})
	}
}
