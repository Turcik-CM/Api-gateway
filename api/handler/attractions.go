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

func (h *attractionsHandler) ListAttractions(c *gin.Context) {
	var post pb.AttractionList

	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 1
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	post.Limit = int64(limits)
	post.Offset = int64(offsets)

	post.Country = c.Query("country")
	post.Name = c.Query("Name")
	post.Description = c.Query("Description")
	post.Category = c.Query("Category")

	req, err := h.attractionsService.ListAttraction(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while listing attractions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *attractionsHandler) SearchAttractions(c *gin.Context) {
	var post pb.AttractionSearch

	limit := c.Query("limit")
	offset := c.Query("offset")

	post.Limit = limit
	post.Offset = offset

	req, err := h.attractionsService.SearchAttraction(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while searching attractions", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *attractionsHandler) AddImageUrl(c *gin.Context) {
	var att pb.AttractionImage
	if err := c.ShouldBindJSON(&att); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.attractionsService.AddAttractionImage(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while adding image url", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

func (h *attractionsHandler) RemoveHistoricalImage(c *gin.Context) {
	var att pb.HistoricalImage
	if err := c.ShouldBindJSON(&att); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.attractionsService.RemoveHistoricalImage(context.Background(), &att)
	if err != nil {
		h.logger.Error("Error occurred while removing image url", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}
