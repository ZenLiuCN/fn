package fn

import (
	"reflect"
	"testing"
)

func TestSliceInsertBefore(t *testing.T) {
	type args[A any] struct {
		s   []A
		idx int
		val []A
	}
	type testCase[A any] struct {
		name  string
		args  args[A]
		wantR []A
	}
	tests := []testCase[int]{
		{"zero", args[int]{[]int{1, 2, 3, 4}, 0, []int{-2, -1, 0}}, []int{-2, -1, 0, 1, 2, 3, 4}},
		{"neg", args[int]{[]int{1, 2, 3, 4}, -1, []int{-2, -1, 0}}, []int{1, 2, 3, -2, -1, 0, 4}},
		{"pos", args[int]{[]int{1, 2, 3, 4}, 2, []int{-2, -1, 0}}, []int{1, 2, -2, -1, 0, 3, 4}},
		{"edge", args[int]{[]int{1, 2, 3, 4}, 4, []int{-2, -1, 0}}, []int{1, 2, 3, 4, -2, -1, 0}},
		{"neg-edge", args[int]{[]int{1, 2, 3, 4}, -4, []int{-2, -1, 0}}, []int{1, 2, 3, 4, -2, -1, 0}},
		{"over", args[int]{[]int{1, 2, 3, 4}, 5, []int{-2, -1, 0}}, []int{1, 2, 3, 4}},
		{"neg-over", args[int]{[]int{1, 2, 3, 4}, -5, []int{-2, -1, 0}}, []int{-2, -1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := SliceInsertBefore(tt.args.s, tt.args.idx, tt.args.val...); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("SliceInsertBefore() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
