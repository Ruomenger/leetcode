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
