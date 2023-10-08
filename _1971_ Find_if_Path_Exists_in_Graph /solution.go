package _1971__Find_if_Path_Exists_in_Graph_

func validPath(n int, edges [][]int, source int, destination int) bool {
	adj := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		adj[edges[i][0]] = append(adj[edges[i][0]], edges[i][1])
		adj[edges[i][1]] = append(adj[edges[i][1]], edges[i][0])
	}

	seen := make(map[int]bool)
	seen[source] = true

	stack := []int{source}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top == destination {
			return true
		}
		stack = stack[:len(stack)-1]
		for _, v := range adj[top] {
			if _, exists := seen[v]; !exists {
				stack = append(stack, v)
				seen[v] = true
			}
		}
	}

	return false
}
