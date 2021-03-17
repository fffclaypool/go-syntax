/* Method
  Go programming language supports special types of functions called methods. In
method declaration syntax, a "receiver" is present to represent the container of
the function. This receiver can be used to call a function using "." operator.
*/

package main

import "testing"

func TestCircle_area(t *testing.T) {
	type fields struct {
		x      float64
		y      float64
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			circle := Circle{
				x:      tt.fields.x,
				y:      tt.fields.y,
				radius: tt.fields.radius,
			}
			if got := circle.area(); got != tt.want {
				t.Errorf("Circle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_method(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method()
		})
	}
}
