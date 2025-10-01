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

func largestArea(grid []string) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	maxArea := 0
	// 方向数组：上、下、左、右
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 找到未访问的主题空间
			if grid[i][j] != '0' && !visited[i][j] {
				// 使用BFS探索整个主题空间
				queue := [][]int{{i, j}}
				visited[i][j] = true
				area := 0
				hasAdjacentCorridor := false
				spaceChar := grid[i][j]

				for len(queue) > 0 {
					cell := queue[0]
					queue = queue[1:]
					x, y := cell[0], cell[1]
					area++

					// 检查四个方向是否有走廊
					for _, dir := range dirs {
						nx, ny := x+dir[0], y+dir[1]
						if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
							// 相邻单元格是走廊
							if grid[nx][ny] == '0' {
								hasAdjacentCorridor = true
							} else if grid[nx][ny] == spaceChar && !visited[nx][ny] {
								// 相邻单元格是同一主题空间且未访问
								visited[nx][ny] = true
								queue = append(queue, []int{nx, ny})
							}
						} else {
							// 超出边界，边界外是走廊
							hasAdjacentCorridor = true
						}
					}
				}

				// 如果不与走廊相邻，更新最大面积
				if !hasAdjacentCorridor && area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

func islandPerimeter(grid [][]int) int {
	ans := 0
	n := len(grid)
	m := len(grid[0])
	for x := range n {
		for y := range m {
			if grid[x][y] != 1 {
				continue
			}
			if x == 0 {
				ans++
			}
			if x == n-1 {
				ans++
			}
			if y == 0 {
				ans++
			}
			if y == m-1 {
				ans++
			}
			if x-1 >= 0 && grid[x-1][y] == 0 {
				ans++
			}
			if x+1 < n && grid[x+1][y] == 0 {
				ans++
			}
			if y-1 >= 0 && grid[x][y-1] == 0 {
				ans++
			}
			if y+1 < m && grid[x][y+1] == 0 {
				ans++
			}
		}
	}
	return ans
}
