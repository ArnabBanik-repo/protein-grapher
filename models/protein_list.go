package models

type ProteinList []Protein

func (p ProteinList) ContainsProtein(protein Protein) bool {
	// TODO: Make this faster

	for _, v := range p {
		if v == protein {
			return true
		}
	}

	return false
}

func (p *ProteinList) Append(protein Protein) {
	if p.ContainsProtein(protein) {
		return
	}
	*p = append(*p, protein)

}

func (p *ProteinList) RemoveDuplicates() {
	set := map[Protein]bool{}

  for _,v := range *p {
    set[v] = true
  }

  newList := ProteinList{}
  for k := range set {
    newList = append(newList, k)
  }

  *p = newList
}
