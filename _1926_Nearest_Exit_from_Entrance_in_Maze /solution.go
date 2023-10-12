package _1926_Nearest_Exit_from_Entrance_in_Maze_

func nearestExit(maze [][]byte, entrance []int) int {
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{entrance[0], entrance[1]})
	maze[entrance[0]][entrance[1]] = '+'

	step := 0
	dir := [][]int{
		[]int{-1, 0},
		[]int{1, 0},
		[]int{0, -1},
		[]int{0, 1},
	}
	for len(queue) > 0 {
		size := len(queue) // number of nodes at the current depth level
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			for _, d := range dir {
				row := curr[0] + d[0]
				col := curr[1] + d[1]
				if row >= 0 && row < len(maze) && col >= 0 && col < len(maze[0]) && maze[row][col] == '.' {
					if row == 0 || row == len(maze)-1 || col == 0 || col == len(maze[0])-1 {
						return step + 1
					}
					queue = append(queue, [2]int{row, col})
					maze[row][col] = '+'
				}
			}
		}
		step++
	}
	return -1
}
