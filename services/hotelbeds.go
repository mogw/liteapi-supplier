package services

import (
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

// Define an interface for Resty client methods we need to mock
type HTTPClient interface {
	R() *resty.Request
}

// Update the service to use the interface
type HotelbedsService struct {
	Client HTTPClient
}

func NewHotelbedsService(client HTTPClient) *HotelbedsService {
	return &HotelbedsService{Client: client}
}

func (s *HotelbedsService) GetCheapestRates(checkin, checkout, currency, guestNationality, hotelIds, occupancies string) (string, error) {
	baseURL := os.Getenv("HOTELBEDS_BASE_URL")
	apiKey := os.Getenv("HOTELBEDS_API_KEY")
	secret := os.Getenv("HOTELBEDS_SECRET")
	liteAPIConfig := os.Getenv("LITEAPI_SUPPLIER_CONFIG")

	// Construct URL with the passed parameters
	url := fmt.Sprintf("%s/hotels/?checkin=%s&checkout=%s&currency=%s&guestNationality=%s&hotelIds=%s&occupancies=%s",
		baseURL, checkin, checkout, currency, guestNationality, hotelIds, occupancies)

	resp, err := s.Client.R().
		SetHeader("x-liteapi-supplier-config", liteAPIConfig).
		SetHeader("Api-key", apiKey).
		SetHeader("X-Signature", secret).
		Get(url)

	if err != nil {
		log.Println("Error calling Hotelbeds API: ", err)
		return "", err
	}

	return resp.String(), nil
}
