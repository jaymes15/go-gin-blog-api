package starwars

import (
	"fmt"
	"io"
	"net/http"
)

type StarWarsInterface interface {
	GetAllCast()
}

type StarWarsRespository struct{}

func NewStarWarsRespository() *StarWarsRespository {
	return &StarWarsRespository{}

}

func (controllers *StarWarsRespository) GetAllCast() {
	url := "https://swapi.dev/api/people"

	// Data to be sent in the request body
	// requestData := []byte(`{"key": "value"}`)

	// Create a new POST request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set request headers if needed
	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := http.DefaultClient

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response status code and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(responseBody))
}
