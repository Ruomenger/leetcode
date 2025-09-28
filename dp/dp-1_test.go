package dp

import "testing"

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]byte
	}

	matrix1 := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	matrix2 := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	matrix3 := [][]byte{
		{'0'},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{matrix1}, 4},
		{"test2", args{matrix2}, 1},
		{"test3", args{matrix3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountSquares(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name: "示例1",
			matrix: [][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			want: 15,
		},
		{
			name: "示例2",
			matrix: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSquares(tt.matrix); got != tt.want {
				t.Errorf("countSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
