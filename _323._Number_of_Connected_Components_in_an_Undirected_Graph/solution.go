package _323__Number_of_Connected_Components_in_an_Undirected_Graph

func countComponents(n int, edges [][]int) int {
	adj := make(map[int][]int)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	cnt := 0
	seen := make(map[int]bool)
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		if _, e := seen[i]; e {
			continue
		}
		cnt++
		stack = append(stack, i)
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			seen[top] = true
			stack = stack[:len(stack)-1]
			for _, v := range adj[top] {
				if _, e := seen[v]; !e {
					seen[v] = true
					stack = append(stack, v)
				}
			}
		}
	}

	return cnt
}
