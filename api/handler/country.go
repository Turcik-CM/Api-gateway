package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type CountriesHandler interface {
	CreateCountry(c *gin.Context)
	GetCountryByID(c *gin.Context)
	UpdateCountry(c *gin.Context)
	DeleteCountry(c *gin.Context)
	ListCountries(c *gin.Context)
}

type CountriesHandlers struct {
	countryService pb.NationalityServiceClient
	logger         *slog.Logger
}

func NewCountriesHandlers(CountryService service.Service, logger *slog.Logger) *CountriesHandlers {
	CountryClient := CountryService.Nationality()
	if CountryClient == nil {
		log.Fatalf("create country client failed")
		return nil
	}
	return &CountriesHandlers{
		countryService: CountryClient,
		logger:         logger,
	}
}

// CreateCountry godoc
// @Summary Create a new Country
// @Description Create a new country
// @Security BearerAuth
// @Tags Country
// @Accept json
// @Produce json
// @Param CreateCountryRequest body nationality.CreateCountryRequest true "Country Info"
// @Success 201 {object} nationality.CreateCountryResponse "Country successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /country/create [post]
func (h *CountriesHandlers) CreateCountry(c *gin.Context) {
	var req pb.CreateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.countryService.CreateCountry(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while creating country", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetCountryByID godoc
// @Summary Get Country by ID
// @Description Get a country by its ID
// @Security BearerAuth
// @Tags Country
// @Produce json
// @Param id path string true "Country ID"
// @Success 200 {object} nationality.GetCountryResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/get/{id} [get]
func (h *CountriesHandlers) GetCountryByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.countryService.GetCountry(context.Background(), &pb.GetCountryRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting country", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCountry godoc
// @Summary Update Country
// @Description Update a country's information
// @Security BearerAuth
// @Tags Country
// @Accept json
// @Produce json
// @Param UpdateCountryRequest body nationality.UpdateCountryRequest true "Update Country Info"
// @Success 200 {object} nationality.UpdateCountryResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/update [put]
func (h *CountriesHandlers) UpdateCountry(c *gin.Context) {
	var req pb.UpdateCountryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.countryService.UpdateCountry(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while updating country", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteCountry godoc
// @Summary Delete Country
// @Description Delete a country by its ID
// @Security BearerAuth
// @Tags Country
// @Produce json
// @Param id path string true "Country ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/delete/{id} [delete]
func (h *CountriesHandlers) DeleteCountry(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.countryService.DeleteCountry(context.Background(), &pb.DeleteCountryRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while deleting country", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListCountries godoc
// @Summary List Countries
// @Description Get a list of countries
// @Security BearerAuth
// @Tags Country
// @Produce json
// @Param filter query nationality.ListCountriesRequest false "Filter Countries"
// @Success 200 {object} nationality.ListCountriesResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/list [get]
func (h *CountriesHandlers) ListCountries(c *gin.Context) {
	var req pb.ListCountriesRequest

	limit := c.Query("limit")
	offset := c.Query("offset")

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 0
	}

	req.Limit = int64(limits)
	req.Offset = int64(offsets)

	resp, err := h.countryService.ListCountries(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while listing countries", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
