package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/pkg/minio"
	"api-gateway/pkg/models"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type AttractionsHandler interface {
	CreateAttraction(c *gin.Context)
	GetAttractionByID(c *gin.Context)
	UpdateAttraction(c *gin.Context)
	DeleteAttraction(c *gin.Context)
	ListAttractions(c *gin.Context)
	SearchAttractions(c *gin.Context)
	UpdateImage(c *gin.Context)
	RemoveHistoricalImage(c *gin.Context)

	CreateAttractionType(c *gin.Context)
	GetAttractionByIDType(c *gin.Context)
	UpdateAttractionType(c *gin.Context)
	DeleteAttractionType(c *gin.Context)
	ListAttractionsType(c *gin.Context)
}

type attractionsHandler struct {
	attractionsService pb.NationalityServiceClient
	logger             *slog.Logger
}

func NewAttractionsHandler(attrService service.Service, logger *slog.Logger) AttractionsHandler {
	attractionsClent := attrService.Nationality()
	if attractionsClent == nil {
		log.Fatalf("Failed to create attractions handler")
		return nil
	}

	return &attractionsHandler{
		attractionsService: attractionsClent,
		logger:             logger,
	}
}

// CreateAttraction godoc
// @Summary Create a new Attraction
// @Description Create a new Attraction, including an image upload
// @Security BearerAuth
// @Tags Attraction
// @Accept multipart/form-data
// @Produce json
// @Param file formData file false "Upload image file"
// @Param name formData string true "Name of the attraction"
// @Param description formData string true "Description of the attraction"
// @Param location formData string true "Location of the attraction"
// @Param city formData string true "city of the attraction"
// @Param category formData string true "Category of the attraction"
// @Success 201 {object} nationality.AttractionResponse "Attraction successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error or invalid file"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /attraction/create [post]
func (h *attractionsHandler) CreateAttraction(c *gin.Context) {
	log.Println("Request received for creating a new attraction")

	var att models.Attraction

	if err := c.ShouldBind(&att); err != nil {
		h.logger.Error("Error occurred while binding form data", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		h.logger.Error("Error occurred while retrieving file from form", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload is required"})
		return
	}

	url, err := minio.UploadNationality(file)
	if err != nil {
		h.logger.Error("Error occurred while uploading file", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while uploading file"})
		return
	}

	att.ImageURL = url

	res := pb.Attraction{
		City:        att.City,
		Name:        att.Name,
		Description: att.Description,
		Category:    att.Category,
		ImageUrl:    att.ImageURL,
		Location:    att.Location,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}

	req, err := h.attractionsService.CreateAttraction(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while creating attraction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// GetAttractionByID godoc
// @Summary Get Attraction by ID
// @Description Get Attraction by its ID
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param id path string true "Attraction ID"
// @Success 200 {object} nationality.AttractionResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/getBy/{id} [get]
func (h *attractionsHandler) GetAttractionByID(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.attractionsService.GetAttractionByID(context.Background(), &pb.AttractionId{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting attraction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAttraction godoc
// @Summary Update Attraction
// @Description Update attractions
// @Security BearerAuth
// @Tags Attraction
// @Accept json
// @Produce json
// @Param Update body nationality.UpdateAttraction true "Update Attraction"
// @Success 200 {object} nationality.AttractionResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/update [put]
func (h *attractionsHandler) UpdateAttraction(c *gin.Context) {
	var att pb.UpdateAttraction

	if err := c.ShouldBindJSON(&att); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req, err := h.attractionsService.UpdateAttractions(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while updating attraction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// DeleteAttraction godoc
// @Summary Delete Attraction
// @Description Delete Attraction by its ID
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param id path string true "Attraction ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/delete/{id} [delete]
func (h *attractionsHandler) DeleteAttraction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.attractionsService.DeleteAttraction(context.Background(), &pb.AttractionId{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting attraction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAttractions godoc
// @Summary List Attraction
// @Description Get a list of Attraction with optional filtering
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param filter query models.AttractionList false "Filter Attraction"
// @Success 200 {object} models.AttractionListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/list [get]
func (h *attractionsHandler) ListAttractions(c *gin.Context) {
	var post pb.AttractionList

	limit := c.Query("limit")
	offset := c.Query("offset")
	city := c.Query("city")
	category := c.Query("category")
	name := c.Query("name")
	description := c.Query("description")

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
	post.Name = name
	post.Description = description
	post.City = city
	post.Category = category
	fmt.Println(post)

	req, err := h.attractionsService.ListAttraction(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing attractions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	c.JSON(http.StatusOK, req)
}

// SearchAttractions godoc
// @Summary List Attraction
// @Description Get a list of Attraction with optional filtering
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param filter query models.AttractionSearch false "Filter Attraction"
// @Success 200 {object} nationality.AttractionListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/list_search [get]
func (h *attractionsHandler) SearchAttractions(c *gin.Context) {
	var post pb.AttractionSearch

	limit := c.Query("limit")
	offset := c.Query("offset")
	search_term := c.Query("search_term")

	post.Limit = limit
	post.Offset = offset

	post.SearchTerm = search_term

	req, err := h.attractionsService.SearchAttraction(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while searching attractions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// AddImageUrl godoc
// @Summary Add Image to Attraction
// @Description Add an image to a Attraction by Attraction ID
// @Security BearerAuth
// @Tags Attraction
// @Accept json
// @Produce json
// @Param id path string true "attraction id"
// @Param file formData file true "Upload image"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/image/{id} [put]
func (h *attractionsHandler) UpdateImage(c *gin.Context) {
	var att pb.AttractionImage

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

	att.ImageUrl = url
	att.Id = id

	req, err := h.attractionsService.AddAttractionImage(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while adding image url", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// RemoveHistoricalImage godoc
// @Summary Remove Image from Attraction
// @Description Remove an image from a Attraction by Attraction ID
// @Security BearerAuth
// @Tags Attraction
// @Accept json
// @Produce json
// @Param id path string true "Image URL"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/remove-image/{id} [delete]
func (h *attractionsHandler) RemoveHistoricalImage(c *gin.Context) {
	var att pb.HistoricalImage

	id := c.Param("id")

	att.Id = id

	req, err := h.attractionsService.RemoveHistoricalImage(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while removing image url", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// CreateAttractionType godoc
// @Summary Create a new Attraction Type
// @Description Create a new Attraction Type
// @Security BearerAuth
// @Tags AttractionType
// @Accept json
// @Produce json
// @Param CreateAttractionTypeRequest body models.CreateAttractionTypeRequest true "Attraction Type Info"
// @Success 201 {object} models.CreateAttractionTypeResponse "Attraction Type successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /attraction-type/create [post]
func (h *attractionsHandler) CreateAttractionType(c *gin.Context) {
	var req pb.CreateAttractionTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.attractionsService.CreateAttractionType(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while creating attraction type", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// GetAttractionByIDType godoc
// @Summary Get Attraction Type by ID
// @Description Get Attraction Type by its ID
// @Security BearerAuth
// @Tags AttractionType
// @Produce json
// @Param id path string true "Attraction Type ID"
// @Success 200 {object} models.GetAttractionTypeResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction-type/get/{id} [get]
func (h *attractionsHandler) GetAttractionByIDType(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.attractionsService.GetAttractionType(context.Background(), &pb.GetAttractionTypeRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while getting attraction type", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAttractionType godoc
// @Summary Update Attraction Type
// @Description Update Attraction Type
// @Security BearerAuth
// @Tags AttractionType
// @Accept json
// @Produce json
// @Param UpdateAttractionTypeRequest body models.UpdateAttractionTypeRequest true "Update Attraction Type"
// @Success 200 {object} models.UpdateAttractionTypeResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction-type/update [put]
func (h *attractionsHandler) UpdateAttractionType(c *gin.Context) {
	var req pb.UpdateAttractionTypeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error occurred while binding JSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.attractionsService.UpdateAttractionType(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while updating attraction type", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteAttractionType godoc
// @Summary Delete Attraction Type
// @Description Delete Attraction Type by its ID
// @Security BearerAuth
// @Tags AttractionType
// @Produce json
// @Param id path string true "Attraction Type ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction-type/delete/{id} [delete]
func (h *attractionsHandler) DeleteAttractionType(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.attractionsService.DeleteAttractionType(context.Background(), &pb.DeleteAttractionTypeRequest{Id: id})
	if err != nil {
		h.logger.Error("Error occurred while deleting attraction type", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ListAttractionsType godoc
// @Summary List Attraction Types
// @Description Get a list of Attraction Types
// @Security BearerAuth
// @Tags AttractionType
// @Produce json
// @Param filter query models.ListAttractionTypesRequest false "Filter Attraction Types"
// @Success 200 {object} models.ListAttractionTypesResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction-type/list [get]
func (h *attractionsHandler) ListAttractionsType(c *gin.Context) {
	var req pb.ListAttractionTypesRequest
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

	resp, err := h.attractionsService.ListAttractionTypes(context.Background(), &req)
	if err != nil {
		h.logger.Error("Error occurred while listing attraction types", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
