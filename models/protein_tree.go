package models

type Protein string
type set map[Protein]bool

type ProteinNode struct {
	Protein  Protein
	Children []*ProteinNode
}

func CreateProteinTree(sourceProtein Protein, network ProteinNetwork) *ProteinNode {
	root := &ProteinNode{Protein: sourceProtein, Children: nil}
	queue := []*ProteinNode{root}
	visited := set{root.Protein: true}

	for len(queue) > 0 {
		currNode := queue[0]
		queue = queue[1:]
		visited[currNode.Protein] = true

		for _, child := range network[currNode.Protein] {
			if visited[child] {
				continue
			}
			childNode := &ProteinNode{Protein: child, Children: nil}
			currNode.Children = append(currNode.Children, childNode)
			queue = append(queue, childNode)
		}
	}

	return root
}
