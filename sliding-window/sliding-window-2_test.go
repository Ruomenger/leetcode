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
