package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStarWars struct{}

func (m *MockStarWars) GetAllCast() {
	// Mocked behavior for GetAllCast in your test scenario
	// Define behavior or assertions as needed for testing purposes
	fmt.Println("::::::::ALLLLLLL")
}

type userArticle struct {
	ID    int    `json:"ID"`
	Image string `json:"Image"`
	Name  string `json:"Name"`
	Email string `json:"Email"`
}

type Article struct {
	ID        int         `json:"ID"`
	Image     string      `json:"Image"`
	Title     string      `json:"Title"`
	Content   string      `json:"Content"`
	CreatedAt string      `json:"CreatedAt"`
	User      userArticle `json:"User"`
}

type ArticleData struct {
	Data []Article `json:"Data"`
}

type AppData struct {
	AppName          *string     `json:"app name"`
	FeaturedArticles ArticleData `json:"featuredArticles"`
	StoriesArticles  ArticleData `json:"storiesArticles"`
}

// Test unauthenticated user can GET
// all articles
func TestUnAuthenticatedUserCanGetAllArticles(t *testing.T) {
	mockStarWars := &MockStarWars{}

	r := getRouter(false)
	controller := NewArticleController(mockStarWars)

	r.GET("/articles", controller.Show)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/articles", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// json.Unmarshal([]byte(w.Body.Bytes()), &jsonMap)
		// r := jsonMap["featuredArticles"].(map[string]interface{})

		// fmt.Println(r["Data"])

		// Test that the page title is "Home Page"
		// You can carry out a lot more detailed tests using libraries that can
		// parse and process HTML pages
		// p, err := ioutil.ReadAll(w.Body)

		// Create a variable to hold the parsed JSON data
		var appData AppData

		// Unmarshal the JSON data into the Go structs
		err := json.Unmarshal([]byte(w.Body.Bytes()), &appData)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return false

		}

		pageOK := len(appData.FeaturedArticles.Data) == 4 && len(appData.StoriesArticles.Data) == 8

		return statusOK && pageOK
	})
}
