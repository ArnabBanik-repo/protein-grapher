package models

type ProteinNetwork map[Protein]ProteinList

func (p *ProteinNetwork) Insert(a, b Protein) {
	aList, ok := (*p)[a]
	if !ok {
		aList = ProteinList{}
	}

	aList.Append(b)
	(*p)[a] = aList
}

func (graph ProteinNetwork) ShortestPaths(source Protein) []ProteinList {
	distances := make(map[Protein]int)
	paths := make(map[Protein]ProteinList)

	for _, nodeList := range graph {
		for _, node := range nodeList {
			distances[node] = -1 // Initialize distances to -1 (representing infinity)
			paths[node] = ProteinList{}
		}
	}

	distances[source] = 0
	paths[source] = ProteinList{source}

	queue := []Protein{source}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if distances[neighbor] == -1 || distances[neighbor] > distances[current]+1 {
				distances[neighbor] = distances[current] + 1
				paths[neighbor] = append(paths[current], neighbor)
				queue = append(queue, neighbor)
			}
		}
	}

	result := []ProteinList{}
	for node := range paths {
		if node != source {
			result = append(result, paths[node])
		}
	}
	return result
}

func (p ProteinNetwork) HasCycle() bool {
	visited := set{}
	for protein := range p {
		if visited[protein] {
			continue
		}
		visited[protein] = true

		dfs_visited := set{}
		if dfs(p, protein, dfs_visited) {
			return true
		}
	}
	return false
}

func dfs(p ProteinNetwork, protein Protein, dfs_visited set) bool {
	if dfs_visited[protein] {
		return true
	}
	dfs_visited[protein] = true
	for _, neighbour := range p[protein] {
		if dfs(p, neighbour, dfs_visited) {
			return true
		}
	}
	return false
}
