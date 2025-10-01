package graph

func numIslands(grid [][]byte) (ans int) {
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}

	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return
}

func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	ans := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] != 1 {
			return 0
		}
		grid[i][j] = 0
		cnt := 1
		cnt += dfs(i, j-1)
		cnt += dfs(i, j+1)
		cnt += dfs(i-1, j)
		cnt += dfs(i+1, j)
		return cnt
	}
	for i := range n {
		for j := range m {
			if grid[i][j] == 1 {
				ans = max(ans, dfs(i, j))
			}
		}
	}
	return ans
}

func countIslands(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	ans := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 {
			return 0
		}
		cnt := grid[i][j]
		grid[i][j] = 0
		cnt += dfs(i, j-1)
		cnt += dfs(i, j+1)
		cnt += dfs(i-1, j)
		cnt += dfs(i+1, j)
		return cnt
	}
	for i := range n {
		for j := range m {
			if grid[i][j] != 0 && dfs(i, j)%k == 0 {
				ans++
			}
		}
	}
	return ans
}
