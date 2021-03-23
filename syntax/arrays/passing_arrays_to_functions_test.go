package main

import "testing"

func Test_getAverage(t *testing.T) {
	type args struct {
		arr  []int
		size int
	}

	arg1 := args{
		arr:  []int{1000, 2, 3, 17, 50},
		size: 5,
	}

	arg2 := args{
		arr:  []int{237, 56, 14, 27, 50, 68, 31, 180, 2458},
		size: 9,
	}

	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			args: arg1,
			want: float32(214),
		},
		{
			args: arg2,
			want: float32(346),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverage(tt.args.arr, tt.args.size); got != tt.want {
				t.Errorf("getAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
