package main

import "testing"

func Test_max(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}

	arg1 := args{
		num1: 2,
		num2: 5,
	}
	arg2 := args{
		num1: 5,
		num2: 2,
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: arg1,
			want: 5,
		},
		{
			args: arg2,
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_functionsswap(t *testing.T) {
	type args struct {
		x string
		y string
	}

	arg1 := args{
		x: "Kumar",
		y: "Mahesh",
	}
	arg2 := args{
		x: "Mahesh",
		y: "Kumar",
	}

	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		{
			args:  arg1,
			want:  "Mahesh",
			want1: "Kumar",
		},
		{
			args:  arg2,
			want:  "Kumar",
			want1: "Mahesh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := functionsswap(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("functionsswap() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("functionsswap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
