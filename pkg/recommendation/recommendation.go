package recommendation

type Recommendation struct {
	Prohibits []string
}

////purpose: gain weight , lost weight
//func (r *Recommendation) WeeklyRecommend(diet string, meals []string, likes []string, dislikes []string, prohibits []string, purpose string, age uint, weight uint, height uint, gender string) {
//
//}
func (r *Recommendation) RemoveProhibits(nameID map[uint]string) ([]uint, error) {
	var recipeIDs []uint

	for key, value := range nameID {
		isProhibited := false
		for _, prohibit := range r.Prohibits {
			if value == prohibit {
				isProhibited = true
				break
			}
		}
		if !isProhibited {
			recipeIDs = append(recipeIDs, key)
		}
	}

	return recipeIDs, nil
}
