package _909_Snakes_and_Ladders_

func snakesAndLadders(board [][]int) int {
	n := len(board)
	cells := make(map[int][2]int)
	num := 1
	for i := n - 1; i >= 0; i-- {
		if (n-1-i)%2 == 0 {
			for j := 0; j < n; j++ {
				cells[num] = [2]int{i, j}
				num++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				cells[num] = [2]int{i, j}
				num++
			}
		}
	}

	moves := 0
	queue := []int{1}
	visited := make(map[int]bool)
	visited[1] = true

	for len(queue) > 0 {
		currLevel := len(queue)
		for i := 0; i < currLevel; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == n*n {
				return moves
			}

			for j := 1; j <= 6 && curr+j <= n*n; j++ {
				next := curr + j
				row, col := cells[next][0], cells[next][1]
				if board[row][col] != -1 {
					next = board[row][col]
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		moves++
	}

	return -1
}
