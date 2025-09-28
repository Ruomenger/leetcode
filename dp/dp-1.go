package dp

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	ans := 0
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				ans = max(ans, dp[i][j])
			}
		}
	}

	return ans * ans
}

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	ans := 0
	for i, row := range matrix {
		for j, x := range row {
			if x == 1 {
				dp[i+1][j+1] = min(dp[i][j], dp[i][j+1], dp[i+1][j]) + 1
				ans += dp[i+1][j+1]
			}
		}
	}
	return ans
}
