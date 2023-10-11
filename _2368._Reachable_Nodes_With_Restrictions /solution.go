package _2368__Reachable_Nodes_With_Restrictions_

func reachableNodes(n int, edges [][]int, restricted []int) int {
	adj := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}
	stack := []int{0}
	cnt := 0
	visited := make(map[int]bool)
	for _, r := range restricted {
		visited[r] = true
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited[top] = true
		cnt++
		for _, v := range adj[top] {
			if _, e := visited[v]; !e {
				stack = append(stack, v)
			}
		}
	}

	return cnt
}
