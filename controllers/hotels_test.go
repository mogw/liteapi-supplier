package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocking the HotelbedsServiceInterface
type MockHotelbedsService struct {
	mock.Mock
}

func (m *MockHotelbedsService) GetCheapestRates(checkin, checkout, currency, guestNationality, hotelIds, occupancies string) (string, error) {
	args := m.Called(checkin, checkout, currency, guestNationality, hotelIds, occupancies)
	return args.String(0), args.Error(1)
}

func TestGetCheapestRates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a mock service
	mockService := new(MockHotelbedsService)

	// Define what the mock should return when GetCheapestRates is called
	mockService.On("GetCheapestRates", "2024-03-15", "2024-03-16", "USD", "US", "129410", "[{\"rooms\":2,\"adults\":2}]").
		Return(`{"rates": "Sample Rate Data"}`, nil)

	// Set up a Gin router and the controller
	r := gin.Default()
	hotelsController := NewHotelsController(mockService) // This now accepts the mock service

	// Create a test request
	req, _ := http.NewRequest("GET", "/hotels/cheapest?checkin=2024-03-15&checkout=2024-03-16&currency=USD&guestNationality=US&hotelIds=129410&occupancies=[{\"rooms\":2,\"adults\":2}]", nil)
	w := httptest.NewRecorder()

	// Register the handler and serve the request
	r.GET("/hotels/cheapest", hotelsController.GetCheapestRates)
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)                  // Expecting 200 OK
	assert.Contains(t, w.Body.String(), "Sample Rate Data") // Expecting the mock response to be returned
}
