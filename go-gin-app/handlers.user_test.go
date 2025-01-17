// handlers.user_test.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// Test that a GET request to the login page returns
// an HTTP error with code 401 for an authenticated user
func TestShowLoginPageAuthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	// Define the route similar to its definition in the routes file
	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/u/login", nil)
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a GET request to the login page returns the login page with
// the HTTP code 200 for an unauthenticated user
func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/u/login", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Login"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to the login route returns
// an HTTP error with code 401 for an authenticated user
func TestLoginAuthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	// Define the route similar to its definition in the routes file
	r.POST("/u/login", ensureNotLoggedIn(), performLogin)

	// Create a request to send to the above route
	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a POST request to login returns a success message for
// an unauthenticated user
func TestLoginUnauthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/u/login", ensureNotLoggedIn(), performLogin)

	// Create a request to send to the above route
	loginPayload := getLoginPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page title is "Successful Login"
	// You can carry out a lot more detailed tests using libraries that can
	// parse and process HTML pages
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful Login</title>") < 0 {
		t.Fail()
	}
}

// Test that a POST request to login returns an error when using
// incorrect credentials
func TestLoginUnauthenticatedIncorrectCredentials(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/u/login", ensureNotLoggedIn(), performLogin)

	// Create a request to send to the above route
	loginPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

// Test that a GET request to the registration page returns
// an HTTP error with code 401 for an authenticated user
func TestShowRegistrationPageAuthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	// Define the route similar to its definition in the routes file
	r.GET("/u/register", ensureNotLoggedIn(), showRegistrationPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/u/register", nil)
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a GET request to the registration page returns the registration
// page with the HTTP code 200 for an unauthenticated user
func TestShowRegistrationPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.GET("/u/register", ensureNotLoggedIn(), showRegistrationPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/u/register", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Login"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Register</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to the registration route returns
// an HTTP error with code 401 for an authenticated user
func TestRegisterAuthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "%7B%22username%22%3A%22user1%22%2C%22password%22%3A%22pass1%22%7D"})

	// Define the route similar to its definition in the routes file
	r.POST("/u/register", ensureNotLoggedIn(), register)

	// Create a request to send to the above route
	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header = http.Header{"Cookie": w.Result().Header["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Println("TestRegisterAuthenticated 235", w.Code)
	// Test that the http status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a POST request to register returns a success message for
// an unauthenticated user
func TestRegisterUnauthenticated(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/u/register", ensureNotLoggedIn(), register)

	// Create a request to send to the above route
	registrationPayload := getRegistrationPOSTPayload()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Println("TestRegisterUnauthenticated 263", w.Code)
	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page title is "Successful registration &amp; Login"
	// You can carry out a lot more detailed tests using libraries that can
	// parse and process HTML pages
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful registration &amp; Login</title>") < 0 {
		t.Fail()
	}

	num, err := deleteUser("u1")
	fmt.Println("TestRegisterUnauthenticated 278", num, err)
	if num == 0 || err != nil {
		t.Fail()
	}
}

// Test that a POST request to register returns an error when
// the username is already in use
func TestRegisterUnauthenticatedUnavailableUsername(t *testing.T) {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Get a new router
	r := getRouter(true)

	// Define the route similar to its definition in the routes file
	r.POST("/u/register", ensureNotLoggedIn(), register)

	// Create a request to send to the above route
	registrationPayload := getRegistrationPOSTPayloadUsedUsername()
	req, _ := http.NewRequest("POST", "/u/register", strings.NewReader(registrationPayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(registrationPayload)))

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	fmt.Println("TestRegisterUnauthenticatedUnavailableUsername 304", w.Code)
	// Test that the http status code is 400
	if w.Code != http.StatusBadRequest {
		t.Fail()
	}
}

func getLoginPOSTPayload() string {
	//params := url.Values{}
	//params.Add("username", "user1")
	//params.Add("password", "pass1")
	//
	//return params.Encode()
	testUser := `{
		"username": "user1",
		"password": "pass1"
	}`
	return testUser
}

func getRegistrationPOSTPayload() string {
	//params := url.Values{}
	//params.Add("username", "u1")
	//params.Add("password", "p1")
	//
	//return params.Encode()

	testUser := `{
		"username": "u1",
		"password": "p1",
		"gender": "unknown",
		"gatorlink": "user1@ufl.edu",
		"gatorPW": "1111"
	}`
	return testUser
}

func getRegistrationPOSTPayloadUsedUsername() string {
	testUser := `{
		"username": "user1",
		"password": "p1",
		"gender": "unknown",
		"gatorlink": "user1@ufl.edu",
		"gatorPW": "1111"
	}`
	return testUser
}
