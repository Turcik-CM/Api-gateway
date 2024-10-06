package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/pkg/minio"
	"api-gateway/pkg/models"
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

	CreateCity(c *gin.Context)
	GetCityByID(c *gin.Context)
	UpdateCity(c *gin.Context)
	DeleteCity(c *gin.Context)
	ListCity(c *gin.Context)
	GetCityByCount(c *gin.Context)
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
// @Security BearerAuth
// @Tags Country
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload image file"
// @Param country_name formData string true "Country name"
// @Success 201 {object} models.AttractionResponse "Attraction successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error or invalid file"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /country/create [post]
func (h *CountriesHandlers) CreateCountry(c *gin.Context) {
	var r models.PostCountry

	if err := c.ShouldBind(&r); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Error occurred while getting file from form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := minio.UploadFlag(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading flag", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.CreateCountryRequest{
		Name:     r.Country,
		ImageUrl: url,
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
// @Success 200 {object} models.GetCountryResponse
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
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload image file"
// @Param id formData string true "Country id"
// @Param name formData string true "Country name"
// @Success 200 {object} models.UpdateCountryResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/update [put]
func (h *CountriesHandlers) UpdateCountry(c *gin.Context) {
	var form models.UpdateCountry

	if err := c.ShouldBind(&form); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := pb.UpdateCountryRequest{
		Id:   form.Id,
		Name: form.Name,
	}

	file, err := c.FormFile("file")
	if err != nil {
		resp, err := h.countryService.UpdateCountry(context.Background(), &req)
		if err != nil {
			h.logger.Error("Error occurred while updating country", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
		return
	}

	url, err := minio.UploadFlag(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading flag", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ImageUrl = url

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
// @Param filter query models.ListCountriesRequest false "Filter Countries"
// @Success 200 {object} models.ListCountriesResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /country/list [get]
func (h *CountriesHandlers) ListCountries(c *gin.Context) {
	var req pb.ListCountriesRequest
	var offset int

	limit := c.Query("limit")
	p := c.Query("page")
	name := c.Query("name")

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}

	if page == 0 || page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limits
	}

	req.Limit = int64(limits)
	req.Offset = int64(offset)
	req.Name = name

	resp, err := h.countryService.ListCountries(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while listing countries", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateCity godoc
// @Summary Create City
// @Description Create City by country id
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param city body models.CreateCityRequest false "Filter Countries"
// @Success 200 {object} models.CreateCityResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/create [post]
func (h *CountriesHandlers) CreateCity(c *gin.Context) {

	var req pb.CreateCityRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.countryService.CreateCity(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while creating city", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetCityByID godoc
// @Summary Get City
// @Description Get city by id
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param id path string true "Filter Countries"
// @Success 200 {object} models.CreateCityResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/get/{id} [get]
func (h *CountriesHandlers) GetCityByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.countryService.GetCity(context.Background(), &pb.GetCityRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting city", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateCity godoc
// @Summary Update City
// @Description Update city
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param city body models.CreateCityResponse true "Filter Countries"
// @Success 200 {object} models.CreateCityResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/update/ [put]
func (h *CountriesHandlers) UpdateCity(c *gin.Context) {
	var req pb.CreateCityResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(req.Id)

	res, err := h.countryService.UpdateCity(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while updating city", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteCity godoc
// @Summary Delete City
// @Description Delete city
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param id path string true "city id"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/delete/{id} [delete]
func (h *CountriesHandlers) DeleteCity(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.countryService.DeleteCountry(context.Background(), &pb.DeleteCountryRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while deleting country", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListCity godoc
// @Summary Get List City
// @Description Get list of city with filter
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param filter query models.FilterCountry false "Filter Countries"
// @Success 200 {object} models.CreateCityResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/get-all/ [get]
func (h *CountriesHandlers) ListCity(c *gin.Context) {
	var req pb.ListCityRequest

	limit := c.Query("limit")
	offset := c.Query("offset")
	name := c.Query("name")

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
	req.Name = name

	resp, err := h.countryService.ListCity(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while listing city", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetCityByCount godoc
// @Summary Get City
// @Description Get cities by country
// @Security BearerAuth
// @Tags City
// @Produce json
// @Param country_id path string true "get cities"
// @Success 200 {object} models.CreateCityResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /city/get-city/{country_id} [get]
func (h *CountriesHandlers) GetCityByCount(c *gin.Context) {
	id := c.Param("country_id")

	resp, err := h.countryService.GetBYCount(context.Background(), &pb.CountryId{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting city", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
