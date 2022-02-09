package app

import (
	"testing"
)

func TestRecommendRecipes(t *testing.T) {
	app, _, usr, _, rcp, _, _, _, _ := Construct()
	usr.Create(app.DB)
	rcp.Create(app.DB)

	tests := []struct {
		userid       uint
		timeinterval string
		err          error
	}{
		{userid: 1, timeinterval: `{"start_date":1643914403,"end_date":1644173603}`, err: nil},
	}
	for _, test := range tests {
		_, err := app.RecommendRecipes(test.userid, test.timeinterval)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

	}
	Destruct(app)
}
