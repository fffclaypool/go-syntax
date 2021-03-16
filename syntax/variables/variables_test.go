package main

import "testing"

func Test_staticTypeDeclaration(t *testing.T) {
	tests := []struct {
		name string
		want float64
	}{
		{
			want: 20.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := staticTypeDeclaration(); got != tt.want {
				t.Errorf("staticTypeDeclaration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dynamicTypeDeclaration(t *testing.T) {
	tests := []struct {
		name  string
		want  float64
		want1 int
	}{
		{
			want:  20.0,
			want1: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := dynamicTypeDeclaration()
			if got != tt.want {
				t.Errorf("dynamicTypeDeclaration() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("dynamicTypeDeclaration() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_mixedVariableDeclaration(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
		want2 string
	}{
		{
			want:  3,
			want1: 4,
			want2: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := mixedVariableDeclaration()
			if got != tt.want {
				t.Errorf("mixedVariableDeclaration() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("mixedVariableDeclaration() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("mixedVariableDeclaration() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
