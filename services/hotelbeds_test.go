package services

import (
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the HTTPClient interface
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) R() *resty.Request {
	// Properly initialize the Request and its Header as http.Header
	req := &resty.Request{
		Header: http.Header{},
	}

	// Mock the Execute method to return a valid response
	req.Execute = func(method, url string) (*resty.Response, error) {
		// Simulate a successful response
		resp := &resty.Response{
			RawResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       http.NoBody, // Simulate a body if needed
			},
		}
		return resp, nil
	}
	return req
}

func TestGetCheapestRates(t *testing.T) {
	// Create a mock HTTP client
	mockClient := new(MockHTTPClient)

	// Initialize the service with the mock client
	hotelbedsService := NewHotelbedsService(mockClient)

	// Call the method and simulate response
	rates, err := hotelbedsService.GetCheapestRates("2024-03-15", "2024-03-16", "USD", "US", "129410", "[{\"rooms\":2, \"adults\": 2}]")

	// Assertions
	assert.NoError(t, err)
	assert.NotEmpty(t, rates)
}
