package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define an interface for HotelbedsService
type HotelbedsServiceInterface interface {
	GetCheapestRates(checkin, checkout, currency, guestNationality, hotelIds, occupancies string) (string, error)
}

type HotelsController struct {
	hotelbedsService HotelbedsServiceInterface
}

func NewHotelsController(service HotelbedsServiceInterface) *HotelsController {
	return &HotelsController{
		hotelbedsService: service,
	}
}

func (h *HotelsController) GetCheapestRates(c *gin.Context) {
	// Capture query parameters from the request
	checkin := c.Query("checkin")
	checkout := c.Query("checkout")
	currency := c.Query("currency")
	guestNationality := c.Query("guestNationality")
	hotelIds := c.Query("hotelIds")
	occupancies := c.Query("occupancies")

	// Validate required fields
	if checkin == "" || checkout == "" || currency == "" || guestNationality == "" || hotelIds == "" || occupancies == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameters"})
		return
	}

	// Call the service function with the captured parameters
	rates, err := h.hotelbedsService.GetCheapestRates(checkin, checkout, currency, guestNationality, hotelIds, occupancies)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hotel rates"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"rates": rates,
	})
}
