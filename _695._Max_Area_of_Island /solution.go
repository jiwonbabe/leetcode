package _695__Max_Area_of_Island_

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	visited := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]int, len(grid[0]))
	}
	maxArea := 0
	stack := make([][2]int, 0)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 && visited[i][j] != 1 {
				rowCol := [2]int{i, j}
				stack = append(stack, rowCol)
				visited[i][j] = 1
				currArea := 0
				for len(stack) > 0 {
					top := stack[len(stack)-1]
					row, col := top[0], top[1]
					stack = stack[:len(stack)-1]
					currArea++

					up := row - 1
					down := row + 1
					left := col - 1
					right := col + 1

					if up >= 0 && grid[up][col] == 1 && visited[up][col] != 1 {
						stack = append(stack, [2]int{up, col})
						visited[up][col] = 1
					}
					if down <= len(grid)-1 && grid[down][col] == 1 && visited[down][col] != 1 {
						stack = append(stack, [2]int{down, col})
						visited[down][col] = 1
					}
					if left >= 0 && grid[row][left] == 1 && visited[row][left] != 1 {
						stack = append(stack, [2]int{row, left})
						visited[row][left] = 1
					}
					if right <= len(grid[0])-1 && grid[row][right] == 1 && visited[row][right] != 1 {
						stack = append(stack, [2]int{row, right})
						visited[row][right] = 1
					}
				}
				if maxArea < currArea {
					maxArea = currArea
				}
			}
		}
	}

	return maxArea
}
