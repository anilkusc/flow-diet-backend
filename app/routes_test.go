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
	app, user, _, _ := Construct()
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
	app, user, _, _ := Construct()
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

func TestGetCalendarRecipesHandler(t *testing.T) {
	app, user, clndr, _ := Construct()
	clndr.Create(app.DB)
	calendar, _ := clndr.List(app.DB)
	calendarJson, _ := json.Marshal(calendar)
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
		output string
		status int
		err    error
	}{
		{output: string(calendarJson) + "\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/calendar/recipes", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetCalendarRecipesHandler))

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
func TestCreateCalendarRecipeHandler(t *testing.T) {
	app, user, clndr, _ := Construct()
	calendarJson, _ := json.Marshal(clndr)
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
		input  string
		output string
		status int
		err    error
	}{
		{input: string(calendarJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/calendar/recipes/create", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.CreateCalendarRecipeHandler))

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

func TestUpdateCalendarRecipeHandler(t *testing.T) {
	app, user, clndr, _ := Construct()
	clndr.Create(app.DB)
	calendarJson, _ := json.Marshal(clndr)
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
		input  string
		output string
		status int
		err    error
	}{
		{input: string(calendarJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/calendar/recipes/update", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.UpdateCalendarRecipeHandler))

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

func TestDeleteCalendarRecipeHandler(t *testing.T) {
	app, user, clndr, _ := Construct()
	clndr.Create(app.DB)
	calendarJson, _ := json.Marshal(clndr)
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
		input  string
		output string
		status int
		err    error
	}{
		{input: string(calendarJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/calendar/recipes/delete", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.DeleteCalendarRecipeHandler))

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
func TestGetAllRecipesHandler(t *testing.T) {
	app, user, _, rcp := Construct()
	rcp.Create(app.DB)
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
		//output string
		status int
		err    error
	}{
		{
			//Since response cannot be compared for now , just compare status and error.
			//output: `[{"ID":1,"CreatedAt":"2022-02-02T18:09:26.285148+03:00","UpdatedAt":"2022-02-02T18:09:26.285148+03:00","DeletedAt":null,"Name":"Test Recipe","ingredients":[{"measurement":{"size":200,"quantity":"gram"},"material":{"type":"fruit","name":"banana","photo_urls":"['s3link1','s3link2']"},"isexist":false,"isoptional":false}],"preperation":"Cook the chickens!","preperation_time":15,"cooking_time_minute":15,"calori":255,"photo_urls":"['S3URL1','S3URL2']","video_urls":"['S3URL1','S3URL2']","for_how_many_people":2}]` + "\n",
			status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/recipes/all", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetAllRecipesHandler))

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		}
		/*
			body, _ := ioutil.ReadAll(rr.Body)
			if string(body) != string(test.output) {
				t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
			}
		*/

	}
	Destruct(app)
}

func TestCreateRecipeHandler(t *testing.T) {
	app, user, _, rcp := Construct()
	recipeJson, _ := json.Marshal(rcp)
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
		input  string
		output string
		status int
		err    error
	}{
		{input: string(recipeJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/recipes/create", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.CreateRecipeHandler))

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
func TestGetRecipeHandler(t *testing.T) {
	app, user, _, rcp := Construct()
	rcp.Create(app.DB)
	recipeJson, _ := json.Marshal(rcp)
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
		input  string
		output string
		status int
		err    error
	}{
		{input: string(recipeJson), output: string(recipeJson) + "\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/recipes/get", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetRecipeHandler))

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
