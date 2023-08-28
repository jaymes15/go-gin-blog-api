package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type userTest struct {
	ID    int    `json:"ID"`
	Image string `json:"Image"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type Data struct {
	User userTest `json:"user"`
}

func getRegistrationPOSTPayload() []byte {
	return []byte(`{"name":"u13","password":"123456789","email":"p1@mail.com"}`)
}

// Test new user can be created via endpouint
func TestCreateNewUser(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(false)

	r.POST("/register", NewAuthController().Register)

	signUpPayload := getRegistrationPOSTPayload()

	// Create a request to send to the above route
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(signUpPayload))

	req.Header.Add("Content-Type", "application/json")

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	var appData Data

	// Unmarshal the JSON data into the Go structs
	err := json.Unmarshal([]byte(w.Body.Bytes()), &appData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		t.Fail()

	}

	// Test that the http status code is 201
	if w.Code != http.StatusCreated && appData.User.Email != "p1@mail.com" {
		t.Fail()
	}

}
