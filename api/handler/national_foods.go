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

type NationalFoodHandler interface {
	CreateNationalFood(c *gin.Context)
	UpdateNationalFood(c *gin.Context)
	GetNationalFoodByID(c *gin.Context)
	DeleteNationalFood(c *gin.Context)
	ListNationalFoods(c *gin.Context)
	UpdateImage(c *gin.Context)
}

type nationalFoodHandler struct {
	nationalFoodService pb.NationalityServiceClient
	logger              *slog.Logger
}

func NewNationalFoodHandler(service service.Service, logger *slog.Logger) NationalFoodHandler {
	Clinet := service.Nationality()
	return &nationalFoodHandler{
		nationalFoodService: Clinet,
		logger:              logger,
	}
}

// CreateNationalFood godoc
// @Summary Create a new NationalFood
// @Description Create a new NationalFood, including an optional image upload
// @Security BearerAuth
// @Tags NationalFood
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload image file (optional)"
// @Param food_name formData string true "food_name of the food"
// @Param food_type formData string true "food_type of the food"
// @Param country_id formData string true "country_id of the food"
// @Param description formData string true "description"
// @Param ingredients formData string true "ingredients of the food"
// @Success 201 {object} models.NationalFoodResponse "National food successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error or invalid file"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /national/create [post]
func (h *nationalFoodHandler) CreateNationalFood(c *gin.Context) {
	log.Println("Request received")

	var a models.NationalFood

	if err := c.ShouldBind(&a); err != nil {
		log.Println("Failed to bind form data")
		h.logger.Error("Error occurred while binding form data:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var url string
	file, err := c.FormFile("file")
	if err == nil {
		url, err = minio.UploadNationality(file)
		if err != nil {
			log.Println("Error occurred while uploading file")
			h.logger.Error("Error occurred while uploading file:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		a.ImageURL = url
	} else {
		log.Println("No file uploaded, continuing without an image")
	}
	res := pb.NationalFood{
		FoodName:     a.,
		FoodType:        a.Name,
		Description: a.Description,
		Nationality: a.Nationality,
		ImageUrl:    a.ImageURL,
		Rating:      a.Rating,
		FoodType:    a.FoodType,
		Ingredients: a.Ingredients,
	}

	resp, err := h.nationalFoodService.CreateNationalFood(context.Background(), &res)
	if err != nil {
		log.Println("Error occurred while creating national food in service")
		h.logger.Error("Error occurred while creating national food:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

// UpdateNationalFood godoc
// @Summary Update NationalFood
// @Description Update NationalFood
// @Security BearerAuth
// @Tags NationalFood
// @Accept json
// @Produce json
// @Param Update body models.UpdateNationalFood true "Update NationalFood"
// @Success 200 {object} nationality.NationalFoodResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /national/update [put]
func (h *nationalFoodHandler) UpdateNationalFood(c *gin.Context) {
	var nat pb.UpdateNationalFood
	if err := c.ShouldBindJSON(&nat); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.nationalFoodService.UpdateNationalFoods(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while updating national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"response": resp})
}

// GetNationalFoodByID godoc
// @Summary Get NationalFood by ID
// @Description Get NationalFood by its ID
// @Security BearerAuth
// @Tags NationalFood
// @Produce json
// @Param id path string true "NationalFood ID"
// @Success 200 {object} nationality.NationalFoodResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /national/getBy/{id} [get]
func (h *nationalFoodHandler) GetNationalFoodByID(c *gin.Context) {
	id := c.Param("id")
	nat := pb.NationalFoodId{
		Id: id,
	}
	resp, err := h.nationalFoodService.GetNationalFoodByID(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while getting national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

// DeleteNationalFood godoc
// @Summary Delete NationalFood
// @Description Delete NationalFood by its ID
// @Security BearerAuth
// @Tags NationalFood
// @Produce json
// @Param id path string true "NationalFood ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /national/delete/{id} [delete]
func (h *nationalFoodHandler) DeleteNationalFood(c *gin.Context) {
	id := c.Param("id")
	nat := pb.NationalFoodId{
		Id: id,
	}
	resp, err := h.nationalFoodService.DeleteNationalFood(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while deleting national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

// ListNationalFoods godoc
// @Summary List NationalFood
// @Description Get a list of NationalFood with optional filtering
// @Security BearerAuth
// @Tags NationalFood
// @Produce json
// @Param filter query models.NationalFoodList false "Filter NationalFood"
// @Success 200 {object} nationality.NationalFoodListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /national/list [get]
func (h *nationalFoodHandler) ListNationalFoods(c *gin.Context) {
	var post models.NationalFoodList

	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 0
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 1
	}

	post.Limit = int64(limits)
	post.Offset = int64(offsets)

	post.Country = c.Query("country")
	resp, err := h.nationalFoodService.ListNationalFood(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing national foods", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": resp})
}

// AddImageUrll godoc
// @Summary Add Image to NationalFood
// @Description Add an image to a NationalFood by NationalFood ID
// @Security BearerAuth
// @Tags NationalFood
// @Accept json
// @Produce json
// @Param id path string true "National food id"
// @Param file formData file true "Upload image"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /national/image/{id} [put]
func (h *nationalFoodHandler) UpdateImage(c *gin.Context) {
	var nat pb.NationalFoodImage

	id := c.Param("id")

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Error occurred while getting file from form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := minio.UploadNationality(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading file", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nat.ImageUrl = url
	nat.Id = id

	resp, err := h.nationalFoodService.AddNationalFoodImage(context.Background(), &nat)
	if err != nil {
		h.logger.Error("Error occurred while adding national food", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp.Message = url
	c.JSON(http.StatusCreated, gin.H{"response": resp})
}
