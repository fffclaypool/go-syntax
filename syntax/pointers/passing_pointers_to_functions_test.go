package main

import "testing"

func Test_passingPointersToFunctions(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			want:  200,
			want1: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := passingPointersToFunctions()
			if got != tt.want {
				t.Errorf("passingPointersToFunctions() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("passingPointersToFunctions() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
