package monotonicstack

import "testing"

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
