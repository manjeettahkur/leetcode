package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoSum1(tt.args.arr, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoSum1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkTwoSum1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tests{
			TwoSum1(tc.args.arr, tc.args.s)
		}
	}
}

