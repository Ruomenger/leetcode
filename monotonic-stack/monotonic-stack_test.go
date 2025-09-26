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

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "example 1",
			nums: []int{1, 2, 1},
			want: []int{2, -1, 2},
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 4, 3},
			want: []int{2, 3, 4, -1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCarFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			name:     "示例1",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			name:     "示例2",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		{
			name:     "示例3",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.target, tt.position, tt.speed); got != tt.want {
				t.Errorf("carFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{
			name:    "示例1",
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			name:    "示例2",
			heights: []int{2, 4},
			want:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximumScore(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "示例1",
			nums: []int{1, 4, 3, 7, 4, 5},
			k:    3,
			want: 15,
		},
		{
			name: "示例2",
			nums: []int{5, 5, 4, 5, 4, 1, 1, 1},
			k:    0,
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumScore(tt.nums, tt.k); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{
			name: "示例1",
			matrix: [][]byte{
				[]byte("10100"),
				[]byte("10111"),
				[]byte("11111"),
				[]byte("10010"),
			},
			want: 6,
		},
		{
			name: "示例2",
			matrix: [][]byte{
				[]byte("0"),
			},
			want: 0,
		},
		{
			name: "示例3",
			matrix: [][]byte{
				[]byte("1"),
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangle(tt.matrix); got != tt.want {
				t.Errorf("maximalRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "示例1",
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want:   6,
		},
		{
			name:   "示例2",
			height: []int{4, 2, 0, 3, 2, 5},
			want:   9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trap(tt.height); got != tt.want {
				t.Errorf("trap() = %v, want %v", got, tt.want)
			}
		})
	}
}
