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
	paths := make(map[Protein][]Protein)

	for node := range graph {
		distances[node] = -1 // Initialize distances to -1 (representing infinity)
		paths[node] = []Protein{}
	}

	distances[source] = 0
	paths[source] = []Protein{source}

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
	for node := range graph {
		if node != source {
			result = append(result, paths[node])
		}
	}
	return result
}

