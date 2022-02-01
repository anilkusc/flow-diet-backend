package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignupHandler(t *testing.T) {
	app, user := Construct()
	userJson, _ := json.Marshal(user)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(userJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/user/signup", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.SignupHandler)

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
func TestSigninHandler(t *testing.T) {
	app, user := Construct()
	userJson, _ := json.Marshal(user)
	app.Signup(string(userJson))
	user.Password = ""
	userJsonOutput, _ := json.Marshal(user)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(userJson), output: string(userJsonOutput) + "\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/user/signin", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.SigninHandler)

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
func TestLogoutHandler(t *testing.T) {
	app, user := Construct()

	userJson, _ := json.Marshal(user)
	app.Signup(string(userJson))
	req, _ := http.NewRequest("POST", "/user/signin", strings.NewReader(string(userJson)))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.SigninHandler)

	handler.ServeHTTP(rr, req)
	sessionCookie := rr.Header()["Set-Cookie"][0]
	ck := strings.Split(sessionCookie, " ")
	ck = strings.Split(ck[0], "=")
	cookie := ck[1] + "="
	tests := []struct {
		cookie string
		output string
		status int
		err    error
	}{
		{cookie: cookie, output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/user/logout", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: test.cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.LogoutHandler))

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
