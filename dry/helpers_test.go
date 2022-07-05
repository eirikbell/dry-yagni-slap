package dry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_union(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{[]int{}, []int{}}, []int{}},
		{"first only", args{[]int{1}, []int{}}, []int{1}},
		{"second only", args{[]int{}, []int{1}}, []int{1}},
		{"no intersect", args{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8}}, []int{1, 3, 5, 7, 9, 2, 4, 6, 8}},
		{"all intersect", args{[]int{1, 2, 3}, []int{1, 2, 3}}, []int{1, 2, 3}},
		{"some intersect", args{[]int{1, 3, 5, 6, 7}, []int{1, 2, 3}}, []int{1, 3, 5, 6, 7, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := union(tt.args.first, tt.args.second)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_intersect(t *testing.T) {
	type args struct {
		first  []int
		second []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"empty", args{[]int{}, []int{}}, []int{}},
		{"first only", args{[]int{1}, []int{}}, []int{}},
		{"second only", args{[]int{}, []int{1}}, []int{}},
		{"no intersect", args{[]int{1, 3, 5, 7, 9}, []int{2, 4, 6, 8}}, []int{}},
		{"all intersect", args{[]int{1, 2, 3}, []int{1, 2, 3}}, []int{1, 2, 3}},
		{"some intersect", args{[]int{1, 3, 5, 6, 7}, []int{1, 2, 3}}, []int{1, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.args.first, tt.args.second)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_mapKeys(t *testing.T) {
	tests := []struct {
		name string
		args map[int]struct{}
		want []int
	}{
		{"empty", map[int]struct{}{}, []int{}},
		{"single", map[int]struct{}{1: {}}, []int{1}},
		{"multi", map[int]struct{}{6: {}, 5: {}, 9: {}, 7: {}, 1: {}}, []int{6, 5, 9, 7, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mapKeys(tt.args)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func Test_sliceToMap(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want map[int]struct{}
	}{
		{"empty", []int{}, map[int]struct{}{}},
		{"single", []int{1}, map[int]struct{}{1: {}}},
		{"multi", []int{6, 5, 9, 7, 1}, map[int]struct{}{6: {}, 5: {}, 9: {}, 7: {}, 1: {}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sliceToMap(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_copy(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"multi", []int{6, 5, 9, 7, 1, 3, 4, 5, 1, 9, 8, 3, 5, 6, 2}, []int{6, 5, 9, 7, 1, 3, 4, 5, 1, 9, 8, 3, 5, 6, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copy(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
