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

func TestFindMaxFish(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{0, 2, 1, 0},
				{4, 0, 0, 3},
				{1, 0, 0, 4},
				{0, 3, 2, 0},
			},
			expected: 7,
		},
		{
			name: "示例2",
			grid: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 1},
			},
			expected: 1,
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxFish(tt.grid)
			if result != tt.expected {
				t.Errorf("测试用例 %s 失败: 预期 %d, 实际 %d", tt.name, tt.expected, result)
			}
		})
	}
}

// 辅助函数：判断两个二维数组是否相等
func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestColorBorder(t *testing.T) {
	// 定义测试用例表格
	tests := []struct {
		name     string
		grid     [][]int
		row      int
		col      int
		color    int
		expected [][]int
	}{
		{
			name: "示例1",
			grid: [][]int{
				{1, 1},
				{1, 2},
			},
			row:   0,
			col:   0,
			color: 3,
			expected: [][]int{
				{3, 3},
				{3, 2},
			},
		},
		{
			name: "示例2",
			grid: [][]int{
				{1, 2, 2},
				{2, 3, 2},
			},
			row:   0,
			col:   1,
			color: 3,
			expected: [][]int{
				{1, 3, 3},
				{2, 3, 3},
			},
		},
		{
			name: "示例3",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			row:   1,
			col:   1,
			color: 2,
			expected: [][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
		},
	}

	// 遍历测试用例表格执行测试
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := colorBorder(tt.grid, tt.row, tt.col, tt.color)
			if !equal(result, tt.expected) {
				t.Errorf("测试用例 %s 失败:\n预期: %v\n实际: %v", tt.name, tt.expected, result)
			}
		})
	}
}

func TestNumEnclaves(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name: "示例1",
			input: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 0, 0, 0},
			},
			want: 3,
		},
		{
			name: "示例2",
			input: [][]int{
				{0, 1, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			want: 0,
		},
		{
			name: "示例3",
			input: [][]int{
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
				{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
				{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
				{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
				{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
				{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
				{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
			},
			want: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := numEnclaves(tc.input)
			if got != tc.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMaxMoves(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name: "示例1",
			input: [][]int{
				{2, 4, 3, 5},
				{5, 4, 9, 3},
				{3, 4, 2, 11},
				{10, 9, 13, 15},
			},
			want: 3,
		},
		{
			name: "示例2",
			input: [][]int{
				{3, 2, 4},
				{2, 1, 9},
				{1, 1, 7},
			},
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxMoves(tc.input)
			if got != tc.want {
				t.Errorf("maxMoves() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestClosedIsland(t *testing.T) {
	testCases := []struct {
		name  string
		input [][]int
		want  int
	}{
		{
			name: "示例1",
			input: [][]int{
				{1, 1, 1, 1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 1, 0},
				{1, 0, 1, 0, 1, 1, 1, 0},
				{1, 0, 0, 0, 0, 1, 0, 1},
				{1, 1, 1, 1, 1, 1, 1, 0},
			},
			want: 2,
		},
		{
			name: "示例2",
			input: [][]int{
				{0, 0, 1, 0, 0},
				{0, 1, 0, 1, 0},
				{0, 1, 1, 1, 0},
			},
			want: 1,
		},
		{
			name: "示例3",
			input: [][]int{
				{1, 1, 1, 1, 1, 1, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 1, 0, 1, 0, 1},
				{1, 0, 1, 1, 1, 0, 1},
				{1, 0, 0, 0, 0, 0, 1},
				{1, 1, 1, 1, 1, 1, 1},
			},
			want: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := closedIsland(tc.input)
			if got != tc.want {
				t.Errorf("closedIsland() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		input [][]byte
		want  [][]byte
	}{
		{
			name: "示例1",
			input: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			want: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "示例2",
			input: [][]byte{
				{'X'},
			},
			want: [][]byte{
				{'X'},
			},
		},
		{
			name: "示例3",
			input: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
			want: [][]byte{
				{'O', 'O', 'O', 'O', 'X', 'X'},
				{'O', 'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'X', 'O'},
				{'O', 'X', 'O', 'X', 'O', 'O'},
				{'O', 'X', 'O', 'O', 'O', 'O'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建输入副本避免测试间相互影响
			inputCopy := make([][]byte, len(tt.input))
			for i := range tt.input {
				inputCopy[i] = make([]byte, len(tt.input[i]))
				copy(inputCopy[i], tt.input[i])
			}

			solve(inputCopy)

			// 验证结果是否与预期一致
			if len(inputCopy) != len(tt.want) {
				t.Errorf("结果行数不符: 实际 %d, 预期 %d", len(inputCopy), len(tt.want))
				return
			}

			for i := range inputCopy {
				if len(inputCopy[i]) != len(tt.want[i]) {
					t.Errorf("第 %d 行列数不符: 实际 %d, 预期 %d", i, len(inputCopy[i]), len(tt.want[i]))
					return
				}

				for j := range inputCopy[i] {
					if inputCopy[i][j] != tt.want[i][j] {
						t.Errorf("位置 (%d,%d) 不符: 实际 %c, 预期 %c", i, j, inputCopy[i][j], tt.want[i][j])
						return
					}
				}
			}
		})
	}
}
