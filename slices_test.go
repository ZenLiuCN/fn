package fn

import (
	"reflect"
	"testing"
)

func TestChunks(t *testing.T) {
	type args struct {
		a []int
		n int
	}
	tests := []struct {
		name  string
		args  args
		wantR [][]int
	}{
		{"normal1", args{[]int{1, 2, 3, 4, 5, 6}, 2}, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"normal2", args{[]int{1, 2, 3, 4, 5, 6, 7}, 2}, [][]int{{1, 2}, {3, 4}, {5, 6}, {7}}},
		{"zero", args{[]int{1, 2, 3, 4, 5, 6, 7}, 0}, nil},
		{"gt", args{[]int{1, 2, 3, 4, 5, 6, 7}, 8}, [][]int{{1, 2, 3, 4, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := SliceChunkBy(tt.args.a, tt.args.n); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Chunks() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
func acc(i int, s []int) int {
	for _, i3 := range s {
		i += i3
	}
	return i
}
func TestFoldN(t *testing.T) {
	type args struct {
		s    []int
		n    int
		init int
		act  func(int, []int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal", args{[]int{1, 2, 3, 4, 5, 6}, 2, 0, acc}, 21},
		{"zero", args{[]int{1, 2, 3, 4, 5, 6}, -1, 0, acc}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceFoldChunk(tt.args.s, tt.args.n, tt.args.init, tt.args.act); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceFoldChunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
