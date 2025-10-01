package graph

import "testing"

func TestNumIslands(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "示例1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "示例2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numIslands(tt.grid)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestMaxAreaOfIsland(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			expected: 6,
		},
		{
			name: "示例2",
			grid: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: 0,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxAreaOfIsland(tt.grid)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestCountIslands(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]int
		k        int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{0, 2, 1, 0, 0},
				{0, 5, 0, 0, 5},
				{0, 0, 1, 0, 0},
				{0, 1, 4, 7, 0},
				{0, 2, 0, 0, 8},
			},
			k:        5,
			expected: 2,
		},
		{
			name: "示例2",
			grid: [][]int{
				{3, 0, 3, 0},
				{0, 3, 0, 3},
				{3, 0, 3, 0},
			},
			k:        3,
			expected: 6,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countIslands(tt.grid, tt.k)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestLargestArea(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     []string
		expected int
	}{
		{
			name: "示例1",
			grid: []string{
				"110",
				"231",
				"221",
			},
			expected: 1,
		},
		{
			name: "示例2",
			grid: []string{
				"11111100000",
				"21243101111",
				"21224101221",
				"11111101111",
			},
			expected: 3,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := largestArea(tt.grid)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}

func TestIslandPerimeter(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			},
			expected: 16,
		},
		{
			name: "示例2",
			grid: [][]int{
				{1},
			},
			expected: 4,
		},
		{
			name: "示例3",
			grid: [][]int{
				{1, 0},
			},
			expected: 4,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := islandPerimeter(tt.grid)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}
