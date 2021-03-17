package main

import "testing"

func Test_forloopBasicOperation(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			want:  15,
			want1: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := forloopBasicOperation()
			if got != tt.want {
				t.Errorf("forloopBasicOperation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("forloopBasicOperation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
