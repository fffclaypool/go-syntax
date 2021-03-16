package main

import (
	"reflect"
	"testing"
)

func Test_expressionSwitch(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expressionSwitch(); got != tt.want {
				t.Errorf("expressionSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_typeSwitch(t *testing.T) {
	tests := []struct {
		name string
		want interface{}
	}{
		{
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeSwitch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("typeSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
