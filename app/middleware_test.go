package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIdControl(t *testing.T) {
	app, user := Construct()
	userJson, _ := json.Marshal(user)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(userJson), output: "id is missing in Path\n", status: 400, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/user/signin", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(IdControl(app.SignupHandler))

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		}
		body, _ := ioutil.ReadAll(rr.Body)
		if string(body) != string(test.output) {
			t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
		}

	}
	Destruct(app)
}
