package _797_All_Paths_From_Source_to_Target_

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := make([][]int, 0)
	curr := 0
	backtracking(graph, curr, []int{0}, &paths)
	return paths
}

func backtracking(graph [][]int, curr int, path []int, paths *[][]int) {
	if curr == len(graph)-1 {
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
		return
	}
	for _, e := range graph[curr] {
		path = append(path, e)
		backtracking(graph, e, path, paths)
		path = path[0 : len(path)-1]
	}
}
