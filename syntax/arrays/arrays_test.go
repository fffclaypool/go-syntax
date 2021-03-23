/* Arrays

 */

package main

import (
	"reflect"
	"testing"
)

func Test_arrays(t *testing.T) {
	tests := []struct {
		name string
		want [10]int
	}{
		{
			want: [10]int{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrays(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("arrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
