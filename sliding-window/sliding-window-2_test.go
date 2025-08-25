package sliding_window

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcabcbb"}, 3},
		{"test2", args{"bbbbb"}, 1},
		{"test3", args{"pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumLengthSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"bcbbbcba"}, 4},
		{"test2", args{"aaaa"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumLengthSubstring(tt.args.s); got != tt.want {
				t.Errorf("maximumLengthSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 0, 1}}, 3},
		{"test2", args{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, 5},
		{"test3", args{[]int{1, 1, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minRemoval(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{2, 1, 5}, 2}, 1},
		{"test2", args{[]int{1, 6, 2, 9}, 3}, 2},
		{"test3", args{[]int{4, 6}, 2}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minRemoval(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("minRemoval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equalSubstring(t *testing.T) {
	type args struct {
		s       string
		t       string
		maxCost int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"abcd", "bcdf", 3}, 3},
		{"test2", args{"abcd", "cdef", 3}, 1},
		{"test3", args{"abcd", "acde", 0}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalFruit(t *testing.T) {
	type args struct {
		fruits []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 1}}, 3},
		{"test2", args{[]int{0, 1, 2, 2}}, 3},
		{"test3", args{[]int{1, 2, 3, 2, 2}}, 4},
		{"test4", args{[]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalFruit(tt.args.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumUniqueSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{4, 2, 4, 5, 6}}, 17},
		{"test2", args{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumUniqueSubarray(tt.args.nums); got != tt.want {
				t.Errorf("maximumUniqueSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSubarrayLength(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 1, 2, 3, 1, 2}, 2}, 6},
		{"test2", args{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 1}, 2},
		{"test3", args{[]int{5, 5, 5, 5, 5, 5, 5}, 4}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarrayLength(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("maxSubarrayLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"TTFF", 2}, 4},
		{"test2", args{"TFFT", 1}, 3},
		{"test3", args{"TTFTTFTT", 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestOnes(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2}, 6},
		{"test2", args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minOperations(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 1, 4, 2, 3}, 5}, 2},
		{"test2", args{[]int{5, 6, 7, 8, 9}, 4}, -1},
		{"test3", args{[]int{3, 2, 20, 1, 1, 3}, 10}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.nums, tt.args.x); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
