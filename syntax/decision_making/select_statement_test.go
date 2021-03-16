package main

import "testing"

func Test_basicOperation(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		{
			want:  0,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := basicOperation()
			if got != tt.want {
				t.Errorf("basicOperation() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("basicOperation() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
