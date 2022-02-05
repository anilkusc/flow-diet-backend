package recommendation

import (
	"sort"
)

type Pair struct {
	Key   uint
	Value uint
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func (r *Recommendation) ReverseSortRecipeIdsByPoint() []uint {
	var SortedIDs []uint
	p := make(PairList, len(r.Recipe_IDsPoints))

	i := 0
	for k, v := range r.Recipe_IDsPoints {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	for _, k := range p {
		SortedIDs = append(SortedIDs, k.Key)
	}
	return SortedIDs
}
