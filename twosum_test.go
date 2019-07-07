package leetcode

import (
	"reflect"
	"testing"
)

type args struct {
	arr []int
	s   int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{
		"One",
		args{[]int{3, 5, 2, -4, 8, 11},
			7,},
		[][]int{{5, 2}, {-4, 11}},
	},
}

func TestTwoSum(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum(tt.args.arr, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests{
			TwoSum(tc.args.arr, tc.args.s)
		}
	}
}
