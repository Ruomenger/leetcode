package array

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{1, 2, 3, 4}, 6}, []int{1, 3}},
		{"test2", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"test3", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"test4", args{[]int{3, 3}, 6}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum1(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum1() = %v, want %v", got, tt.want)
			}
			if got := twoSum2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum2() = %v, want %v", got, tt.want)
			}
			if got := twoSum3(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 3}, []int{2}}, 2},
		{"test2", args{[]int{1, 2}, []int{3, 4}}, 2.5},
		{"test3", args{[]int{2, 2, 4, 4}, []int{2, 2, 2, 4, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
