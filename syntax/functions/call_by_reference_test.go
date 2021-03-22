package main

import "testing"

func Test_callbyreferenceswap(t *testing.T) {
	type args struct {
		x *int
		y *int
	}

	var a int = 100
	var b int = 200

	arg := args{
		x: &a,
		y: &b,
	}

	tests := []struct {
		name string
		args args
	}{
		{
			args: arg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callbyreferenceswap(tt.args.x, tt.args.y)
		})
	}
}
