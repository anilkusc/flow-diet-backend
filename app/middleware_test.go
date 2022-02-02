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
	app, user, _, _ := Construct()
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
		handler := http.HandlerFunc(app.IdControl(app.TestHandler))

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
func TestAuth(t *testing.T) {
	app, user, _, _ := Construct()

	userJson, _ := json.Marshal(user)
	app.Signup(string(userJson))
	req, _ := http.NewRequest("POST", "/user/signin", strings.NewReader(string(userJson)))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SigninHandler)

	handler.ServeHTTP(rr, req)
	sessionCookie := rr.Header()["Set-Cookie"][0]
	ck := strings.Split(sessionCookie, " ")
	ck = strings.Split(ck[0], "session=")
	cookie := ck[1]
	tests := []struct {
		cookie string
		output string
		status int
		err    error
	}{
		{cookie: cookie, output: "Hello\n", status: 200, err: nil},
		{cookie: "XTY0MzcyMjU3NHxEdi1CQkFFQ180SUFBUkFCRUFBQVNmLUNBQUlHYzNSeWFXNW5EQVlBQkhKdmJHVUdjM1J5YVc1bkRBWUFCSFZ6WlhJR2MzUnlhVzVuREE4QURXRjFkR2hsYm5ScFkyRjBaV1FHYzNSeWFXNW5EQVlBQkhSeWRXVT183Iu8uJkD1hKMYMOx6N2GZqnnl-sXrfm_B_1Pey8ZbgA=", output: "Forbidden\n", status: 403, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/user/test", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: test.cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.TestHandler))

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

func TestAuthz(t *testing.T) {
	app, user, _, _ := Construct()
	user.Role = "admin"
	userJson, _ := json.Marshal(user)
	app.Signup(string(userJson))
	req, _ := http.NewRequest("POST", "/user/signin", strings.NewReader(string(userJson)))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SigninHandler)

	handler.ServeHTTP(rr, req)
	sessionCookie := rr.Header()["Set-Cookie"][0]
	ck := strings.Split(sessionCookie, " ")
	ck = strings.Split(ck[0], "session=")
	cookie := ck[1]
	tests := []struct {
		cookie string
		output string
		status int
		err    error
	}{
		{cookie: cookie, output: "Hello\n", status: 200, err: nil},
		{cookie: "XTY0MzcyMjU3NHxEdi1CQkFFQ180SUFBUkFCRUFBQVNmLUNBQUlHYzNSeWFXNW5EQVlBQkhKdmJHVUdjM1J5YVc1bkRBWUFCSFZ6WlhJR2MzUnlhVzVuREE4QURXRjFkR2hsYm5ScFkyRjBaV1FHYzNSeWFXNW5EQVlBQkhSeWRXVT183Iu8uJkD1hKMYMOx6N2GZqnnl-sXrfm_B_1Pey8ZbgA=", output: "Forbidden\n", status: 403, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/user/test", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: test.cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.Authz(app.TestHandler)))

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
