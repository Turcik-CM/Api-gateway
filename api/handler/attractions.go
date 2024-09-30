package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type AttractionsHandler interface {
	CreateAttraction(c *gin.Context)
	GetAttractionByID(c *gin.Context)
	UpdateAttraction(c *gin.Context)
	DeleteAttraction(c *gin.Context)
	ListAttractions(c *gin.Context)
	SearchAttractions(c *gin.Context)
	AddImageUrl(c *gin.Context)
	RemoveHistoricalImage(c *gin.Context)
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
// @Summary Create Attraction
// @Description Create a new Attraction
// @Security BearerAuth
// @Tags Attraction
// @Accept json
// @Produce json
// @Param Create body models.Attraction true "Create Attraction"
// @Success 201 {object} models.AttractionResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/create [post]
func (h *attractionsHandler) CreateAttraction(c *gin.Context) {
	var att pb.Attraction
	if err := c.ShouldBindJSON(&att); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.attractionsService.CreateAttraction(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while creating attraction", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, req)
}

// GetAttractionByID godoc
// @Summary Get Attraction by ID
// @Description Get Attraction by its ID
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param id path string true "Attraction ID"
// @Success 200 {object} models.AttractionResponse
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
// @Param Update body models.UpdateAttraction true "Update Attraction"
// @Success 200 {object} models.AttractionResponse
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
	country := c.Query("country")
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
	post.Country = country
	post.Category = category

	req, err := h.attractionsService.ListAttraction(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing attractions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// SearchAttractions godoc
// @Summary List Attraction
// @Description Get a list of Attraction with optional filtering
// @Security BearerAuth
// @Tags Attraction
// @Produce json
// @Param filter query models.AttractionSearch false "Filter Attraction"
// @Success 200 {object} models.AttractionListResponse
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
// @Param image body models.AttractionImage true "Image URL"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /attraction/add-image [post]
func (h *attractionsHandler) AddImageUrl(c *gin.Context) {
	var att pb.AttractionImage
	if err := c.ShouldBindJSON(&att); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(att)
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
