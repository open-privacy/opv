package functional_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestTokenizationAndDetokenization(t *testing.T) {
	// Simulate request data
	requestData := []map[string]string{
		{"name_first": "John", "name_last": "Dsilva"},
		{"name_first": "Tera"},
		{"name_first": "Maria"},
	}

	// Send request to server for tokenization
	tokenizedData, err := sendRequest("http://127.0.0.1:28001/tokenize", requestData)
	if err != nil {
		t.Fatalf("Error sending tokenize request: %v", err)
	}

	// Send request to server for detokenization
	detokenizedData, err := sendRequest("http://127.0.0.1:28001/detokenize", tokenizedData)
	if err != nil {
		t.Fatalf("Error sending detokenize request: %v", err)
	}

	// Compare original data with detokenized data
	for i := range requestData {
		if !isEqual(requestData[i], detokenizedData[i]) {
			t.Errorf("Detokenized data doesn't match original data at index %d", i)
		}
	}
}

// Send request to server and decode the response
func sendRequest(url string, requestData []map[string]string) ([]map[string]string, error) {
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var responseData []map[string]string
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}

// Function to compare two maps
func isEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
