package sliding_window

import (
	"reflect"
	"testing"
)

func Test_maxVowels(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abciiidef", 3}, 3},
		{"test2", args{"aeiou", 2}, 2},
		{"test3", args{"leetcode", 3}, 2},
		{"test4", args{"rhythms", 4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxVowels(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("maxVowels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75},
		{"test2", args{[]int{5}, 1}, 5},
		{"test3", args{[]int{-1}, 1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr       []int
		k         int
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 2, 2, 2, 5, 5, 5, 8}, 3, 4}, 3},
		{"test2", args{[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}, 3, 5}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numOfSubarrays(tt.args.arr, tt.args.k, tt.args.threshold); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAverages(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{7, 4, 3, 9, 1, 8, 5, 2, 6}, 3}, []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}},
		{"test2", args{[]int{100000}, 0}, []int{100000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAverages(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAverages() = %v, want %v", got, tt.want)
			}
		})
	}
}
