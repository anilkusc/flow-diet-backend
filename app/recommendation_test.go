package app

import (
	"testing"
)

func TestRecommendRecipes(t *testing.T) {
	app, _, usr, _, rcp, _, _, _ := Construct()
	usr.Create(app.DB)
	rcp.Create(app.DB)

	tests := []struct {
		input uint
		err   error
	}{
		{input: 1, err: nil},
	}
	for _, test := range tests {
		_, err := app.RecommendRecipes(test.input)
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

	}
	Destruct(app)
}
