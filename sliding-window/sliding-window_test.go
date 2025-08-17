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

func Test_maxSum(t *testing.T) {
	type args struct {
		nums []int
		m    int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"test1", args{[]int{2, 6, 7, 3, 1, 7}, 3, 4}, 18},
		{"test2", args{[]int{5, 9, 9, 2, 4, 5, 4}, 1, 3}, 23},
		{"test3", args{[]int{1, 2, 1, 2, 1, 2, 1}, 3, 3}, 0},
		{"test4", args{[]int{1, 1, 1, 2}, 2, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.nums, tt.args.m, tt.args.k); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4, 5, 6, 1}, 3}, 12},
		{"test2", args{[]int{9, 7, 7, 9, 7, 7, 9}, 7}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.cardPoints, tt.args.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSatisfied(t *testing.T) {
	type args struct {
		customers []int
		grumpy    []int
		minutes   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3}, 16},
		{"test2", args{[]int{1}, []int{0}, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSatisfied(tt.args.customers, tt.args.grumpy, tt.args.minutes); got != tt.want {
				t.Errorf("maxSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		k         int
		startTime []int
		endTime   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{5, 1, []int{1, 3}, []int{2, 5}}, 2},
		{"test2", args{10, 1, []int{0, 2, 9}, []int{1, 4, 10}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.k, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSwaps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{0, 1, 0, 1, 1, 0, 0}}, 1},
		{"test2", args{[]int{0, 1, 1, 1, 0, 0, 1, 1, 0}}, 2},
		{"test3", args{[]int{1, 1, 0, 0, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.nums); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decrypt(t *testing.T) {
	type args struct {
		code []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{[]int{5, 7, 1, 4}, 3}, []int{12, 10, 16, 13}},
		{"test2", args{[]int{2, 4, 9, 3}, -2}, []int{12, 5, 6, 13}},
		{"test3", args{[]int{1, 2, 3, 4}, 0}, []int{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decrypt(tt.args.code, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFreq(t *testing.T) {
	type args struct {
		s          string
		maxLetters int
		minSize    int
		maxSize    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"aababcaab", 2, 3, 4}, 2},
		{"test2", args{"aaaa", 1, 3, 3}, 2},
		{"test3", args{"aabcabcab", 2, 2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreq(tt.args.s, tt.args.maxLetters, tt.args.minSize, tt.args.maxSize); got != tt.want {
				t.Errorf("maxFreq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minFlips(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{"111000"}, 2},
		{"test2", args{"010"}, 0},
		{"test3", args{"00100110101"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minFlips(tt.args.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkInclusion(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"ab", "eidbaooo"}, true},
		{"test2", args{"ab", "eidboaoo"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subStrHash(t *testing.T) {
	type args struct {
		s         string
		power     int
		modulo    int
		k         int
		hashValue int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"leetcode", 7, 20, 2, 0}, "ee"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subStrHash(tt.args.s, tt.args.power, tt.args.modulo, tt.args.k, tt.args.hashValue); got != tt.want {
				t.Errorf("subStrHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryString(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"0110", 3}, true},
		{"test2", args{"0110", 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := queryString(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("queryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAnagrams(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test1", args{"cbaebabacd", "abc"}, []int{0, 6}},
		{"test2", args{"abab", "ab"}, []int{0, 1, 2}},
		{"test3", args{"abacbabc", "abc"}, []int{1, 2, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
