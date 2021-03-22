package main

import "testing"

func Test_callbyvalueswap(t *testing.T) {
	type args struct {
		x int
		y int
	}

	arg1 := args{
		x: 100,
		y: 200,
	}
	arg2 := args{
		x: 200,
		y: 100,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: arg1,
			want: 100,
		},
		{
			args: arg2,
			want: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callbyvalueswap(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("callbyvalueswap() = %v, want %v", got, tt.want)
			}
		})
	}
}
