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

func findMaxFish(grid [][]int) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			return 0
		}
		cnt := grid[x][y]
		grid[x][y] = 0
		cnt += dfs(x+1, y)
		cnt += dfs(x-1, y)
		cnt += dfs(x, y+1)
		cnt += dfs(x, y-1)
		return cnt
	}
	for x := range n {
		for y := range m {
			if grid[x][y] != 0 {
				ans = max(ans, dfs(x, y))
			}
		}
	}
	return ans
}

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}
	origin := grid[row][col]
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] != origin || visited[x][y] {
			return
		}
		visited[x][y] = true
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < m {
				if grid[nx][ny] != origin && !visited[nx][ny] {
					grid[x][y] = color
				} else if grid[nx][ny] == origin && !visited[nx][ny] {
					dfs(nx, ny)
				}
			} else {
				grid[x][y] = color
			}
		}
	}
	dfs(row, col)
	return grid
}

func numEnclaves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		dfs(x-1, y)
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
	}
	for x := range n {
		dfs(x, 0)
		dfs(x, m-1)
	}
	for y := range m {
		dfs(0, y)
		dfs(n-1, y)
	}
	ans := 0
	for i := range n {
		for j := range m {
			if grid[i][j] == 1 {
				ans++
			}
		}
	}
	return ans
}

func maxMoves(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	ans := 0
	var dfs func(int, int)
	dfs = func(x, y int) {
		ans = max(ans, y)
		if ans == m-1 {
			return
		}
		for i := max(x-1, 0); i < min(x+2, n); i++ {
			if grid[i][y+1] > grid[x][y] {
				dfs(i, y+1)
			}
		}
		grid[x][y] = 0
	}
	for i := range n {
		dfs(i, 0)
	}
	return ans
}

func closedIsland(grid [][]int) int {
	ans := 0
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m {
			return false
		}
		if grid[x][y] == 1 {
			return true
		}
		if visited[x][y] {
			return true
		}
		visited[x][y] = true
		d1 := dfs(x-1, y)
		d2 := dfs(x+1, y)
		d3 := dfs(x, y-1)
		d4 := dfs(x, y+1)
		return d1 && d2 && d3 && d4
	}
	for i := range n {
		for j := range m {
			if grid[i][j] == 0 && !visited[i][j] && dfs(i, j) {
				ans++
			}
		}
	}

	return ans
}

func solve(board [][]byte) {
	n, m := len(board), len(board[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}
	var isSurrounded func(x, y int) bool
	isSurrounded = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m {
			return false
		}
		if visited[x][y] || board[x][y] == 'X' {
			return true
		}
		visited[x][y] = true
		b1 := isSurrounded(x-1, y)
		b2 := isSurrounded(x+1, y)
		b3 := isSurrounded(x, y-1)
		b4 := isSurrounded(x, y+1)
		return b1 && b2 && b3 && b4
	}
	var change func(x, y int)
	change = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] == 'X' {
			return
		}
		board[x][y] = 'X'
		change(x-1, y)
		change(x+1, y)
		change(x, y-1)
		change(x, y+1)
	}
	for x := range n {
		for y := range m {
			if board[x][y] == 'O' && !visited[x][y] && isSurrounded(x, y) {
				change(x, y)
			}
		}
	}
}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	n, m := len(grid1), len(grid1[0])
	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, m)
	}
	var dfs func(x, y int) bool
	dfs = func(x, y int) bool {
		if x < 0 || x >= n || y < 0 || y >= m || visited[x][y] {
			return true
		}
		visited[x][y] = true
		if grid2[x][y] == 0 {
			return true
		}

		d1 := dfs(x-1, y)
		d2 := dfs(x+1, y)
		d3 := dfs(x, y-1)
		d4 := dfs(x, y+1)
		if grid1[x][y] != grid2[x][y] {
			return false
		}

		return d1 && d2 && d3 && d4
	}
	ans := 0
	for x := range n {
		for y := range m {
			if !visited[x][y] && grid2[x][y] == 1 && dfs(x, y) {
				ans++
			}
		}
	}
	return ans
}

func hasValidPath(grid [][]int) bool {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return false
	}

	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// 定义每个数字可以前往的方向: 上、右、下、左
	// 索引0-3分别代表上、右、下、左，true表示可以朝该方向移动
	canGo := [][]bool{
		{},                         // 0 占位
		{false, true, false, true}, // 1: 右、左
		{true, false, true, false}, // 2: 上、下
		{false, false, true, true}, // 3: 下、左
		{false, true, true, false}, // 4: 右、下
		{true, false, false, true}, // 5: 上、左
		{true, true, false, false}, // 6: 上、右
	}

	// 相反方向映射: 上<->下, 左<->右
	oppositeDir := []int{2, 3, 0, 1} // 上的反方向是下，右的反方向是左，下的反方向是上，左的反方向是右

	// 方向对应的坐标变化 (行变化, 列变化)
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} // 上、右、下、左

	var dfs func(row, col int) bool
	dfs = func(row, col int) bool {
		// 到达终点
		if row == rows-1 && col == cols-1 {
			return true
		}

		visited[row][col] = true
		cellType := grid[row][col]

		// 尝试所有可能的方向
		for dir := 0; dir < 4; dir++ {
			if canGo[cellType][dir] {
				newRow, newCol := row+dirs[dir][0], col+dirs[dir][1]
				// 检查新位置是否有效且未访问过
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol] {
					// 检查新单元格是否接受来自当前方向的移动
					newCellType := grid[newRow][newCol]
					if canGo[newCellType][oppositeDir[dir]] {
						if dfs(newRow, newCol) {
							return true
						}
					}
				}
			}
		}

		return false
	}

	return dfs(0, 0)
}
