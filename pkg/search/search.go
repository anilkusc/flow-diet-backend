package search

import (
	"strings"
)

type Search struct {
	Word string // word for searching
}

func (s *Search) FindRecipesByName(recipes map[uint]string) ([]uint, error) {

	var ids []uint
	for key, element := range recipes {
		if strings.Contains(element, s.Word) {
			ids = append(ids, key)
		}
	}

	return ids, nil
}
