package monotonicstack

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "示例1: 包含多个温度变化",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "示例2: 温度持续升高",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "示例3: 温度阶梯升高",
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			name:         "边界测试: 空输入",
			temperatures: []int{},
			want:         []int{},
		},
		{
			name:         "边界测试: 单个元素",
			temperatures: []int{50},
			want:         []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temperatures)
			if len(got) != len(tt.want) {
				t.Errorf("dailyTemperatures() 长度不匹配: 得到 %d, 期望 %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("dailyTemperatures() 索引 %d 不匹配: 得到 %d, 期望 %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func Test_finalPrices(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   []int
	}{
		{
			name:   "示例1: 存在多个折扣情况",
			prices: []int{8, 4, 6, 2, 3},
			want:   []int{4, 2, 4, 2, 3},
		},
		{
			name:   "示例2: 无折扣情况",
			prices: []int{1, 2, 3, 4, 5},
			want:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "示例3: 包含相同价格折扣",
			prices: []int{10, 1, 1, 6},
			want:   []int{9, 0, 1, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finalPrices(tt.prices)
			// 比较两个切片是否完全相等
			if len(got) != len(tt.want) {
				t.Errorf("finalPrices() 长度不匹配: 得到 %d, 期望 %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("finalPrices() 索引 %d 不匹配: 得到 %d, 期望 %d", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestNextGreaterElement(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{
			name:  "example 1",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		{
			name:  "example 2",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
