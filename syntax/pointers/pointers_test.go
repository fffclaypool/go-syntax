/* Pointers

 */

package main

import "testing"

func Test_pointers(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pointers()
			if got != tt.want {
				t.Errorf("pointers() got = %v, want %v", got, tt.want)
			}
		})
	}
}
