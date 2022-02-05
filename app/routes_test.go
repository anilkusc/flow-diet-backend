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
	app, _, user, _, _, _, _, _ := Construct()
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
	app, _, user, _, _, _, _, _ := Construct()
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
	app, cookie, _, _, _, _, _, _ := Construct()

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
	app, cookie, _, clndr, _, _, _, _ := Construct()
	clndr.Create(app.DB)
	calendar, _ := clndr.List(app.DB)
	calendarJson, _ := json.Marshal(calendar)

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
	app, cookie, _, clndr, _, _, _, _ := Construct()
	calendarJson, _ := json.Marshal(clndr)

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
	app, cookie, _, clndr, _, _, _, _ := Construct()
	clndr.Create(app.DB)
	calendarJson, _ := json.Marshal(clndr)

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
	app, cookie, _, clndr, _, _, _, _ := Construct()
	clndr.Create(app.DB)
	calendarJson, _ := json.Marshal(clndr)

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
	app, cookie, _, _, rcp, _, _, _ := Construct()
	rcp.Create(app.DB)

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
	app, cookie, _, _, rcp, _, _, _ := Construct()
	recipeJson, _ := json.Marshal(rcp)

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
	app, cookie, _, _, rcp, _, _, _ := Construct()
	rcp.Create(app.DB)
	recipeJson, _ := json.Marshal(rcp)

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

func TestUpdateRecipeHandler(t *testing.T) {
	app, cookie, user, _, rcp, _, _, _ := Construct()
	rcp.Create(app.DB)
	rcp.ID = 1
	rcp.Calori = 10
	user.Role = "admin"
	recipeJson, _ := json.Marshal(rcp)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(recipeJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/recipes/update", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.UpdateRecipeHandler))

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
func TestDeleteRecipeHandler(t *testing.T) {
	app, cookie, _, _, rcp, _, _, _ := Construct()
	rcp.Create(app.DB)
	rcp.ID = 1
	recipeJson, _ := json.Marshal(rcp)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(recipeJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/recipes/delete", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.DeleteRecipeHandler))

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

func TestGetAllShoppingsHandler(t *testing.T) {
	app, cookie, _, _, _, shp, _, _ := Construct()
	shp.Start_Date = 1643937031
	shp.End_Date = 1644016231
	shp.Create(app.DB)
	tests := []struct {
		status       int
		shoppingJson string
		err          error
	}{
		{shoppingJson: `{"start_date":1643850631 ,"end_date":1644109831}`, status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/shopping/all", strings.NewReader(test.shoppingJson))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetAllShoppingsHandler))

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		}

	}
	Destruct(app)
}

func TestCreateShoppingHandler(t *testing.T) {
	app, cookie, _, _, _, shp, _, _ := Construct()
	shoppingJson, _ := json.Marshal(shp)

	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(shoppingJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/shopping/create", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.CreateShoppingHandler))

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
func TestGetShoppingHandler(t *testing.T) {
	app, cookie, _, _, _, shp, _, _ := Construct()
	shp.Create(app.DB)
	shoppingJson, _ := json.Marshal(shp)

	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(shoppingJson), output: string(shoppingJson) + "\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/shopping/get", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetShoppingHandler))

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

func TestUpdateShoppingHandler(t *testing.T) {
	app, cookie, user, _, _, shp, _, _ := Construct()
	shp.Create(app.DB)
	shp.ID = 1
	shp.Start_Date = 10
	user.Role = "admin"
	shoppingJson, _ := json.Marshal(shp)

	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(shoppingJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/shopping/update", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.UpdateShoppingHandler))

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
func TestDeleteShoppingHandler(t *testing.T) {
	app, cookie, _, _, _, shp, _, _ := Construct()
	shp.Create(app.DB)
	shp.ID = 1
	shoppingJson, _ := json.Marshal(shp)
	tests := []struct {
		input  string
		output string
		status int
		err    error
	}{
		{input: string(shoppingJson), output: "OK\n", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/shopping/delete", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.DeleteShoppingHandler))

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

func TestSearchRecipesHandler(t *testing.T) {
	app, cookie, _, _, rcp, _, _, _ := Construct()
	rcp.Create(app.DB)
	tests := []struct {
		input string
		//output string
		status int
		err    error
	}{
		{input: "Test", status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("POST", "/search/recipes", strings.NewReader(test.input))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.SearchRecipesHandler))

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		} /*
			body, _ := ioutil.ReadAll(rr.Body)
			if string(body) != string(test.output) {
				t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
			}*/

	}
	Destruct(app)
}
func TestGetRecommendationsHandler(t *testing.T) {
	app, cookie, user, _, rcp, _, _, _ := Construct()
	user.Signup(app.DB)
	rcp.Create(app.DB)
	//recommendationJson, _ := json.Marshal(recom)
	tests := []struct {
		//output string
		status int
		err    error
	}{
		{ /*output: string(recommendationJson),*/ status: 200, err: nil},
	}
	for _, test := range tests {
		req, err := http.NewRequest("GET", "/recommendation/getrecipes", strings.NewReader(""))
		if err != nil {
			t.Errorf("Error is: %v . Expected: %v", err, test.err)
		}
		req.AddCookie(&http.Cookie{Name: "session", Value: cookie})
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(app.Auth(app.GetRecommendationsHandler))

		handler.ServeHTTP(rr, req)

		if rr.Result().StatusCode != test.status {
			t.Errorf("Response status is: %v . Expected: %v", rr.Result().StatusCode, test.status)
		}
		/*
			body, _ := ioutil.ReadAll(rr.Body)
			if string(body) != string(test.output) {
				t.Errorf("Response is: %v . Expected: %v", string(body), test.output)
			}*/

	}
	Destruct(app)
}
