package handler

import (
	pb "api-gateway/genproto/nationality"
	"api-gateway/pkg/minio"
	"api-gateway/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"log/slog"
	"net/http"
	"strconv"
)

type HistoryHandler interface {
	AddHistorical(c *gin.Context)
	UpdateHistoricals(c *gin.Context)
	GetHistoricalByID(c *gin.Context)
	DeleteHistorical(c *gin.Context)
	ListHistorical(c *gin.Context)
	SearchHistorical(c *gin.Context)
	UpdateHisImage(c *gin.Context)
}

type historyHandler struct {
	historyService pb.NationalityServiceClient
	logger         *slog.Logger
}

func NewHistoryHandler(historyService service.Service, logger *slog.Logger) HistoryHandler {
	historyClient := historyService.Nationality()
	if historyClient == nil {
		log.Fatalf("Cannot create history handler")
		return nil
	}
	return &historyHandler{
		historyService: historyClient,
		logger:         logger,
	}
}

// AddHistorical godoc
// @Summary Create Historical
// @Description Create a new Historical
// @Security BearerAuth
// @Tags Historical
// @Accept json
// @Produce json
// @Param Create body models.Historical true "Create Historical"
// @Param file formData file true "Upload image"
// @Success 201 {object} models.HistoricalResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/create [post]
func (h *historyHandler) AddHistorical(c *gin.Context) {
	var his pb.Historical

	if err := c.ShouldBindJSON(&his); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	his.ImageUrl = url

	req, err := h.historyService.AddHistorical(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling AddHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

// UpdateHistoricals godoc
// @Summary Update Historical
// @Description Update Historical
// @Security BearerAuth
// @Tags Historical
// @Accept json
// @Produce json
// @Param Update body models.UpdateHistorical true "Update Historical"
// @Success 200 {object} models.HistoricalResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/update [put]
func (h *historyHandler) UpdateHistoricals(c *gin.Context) {
	var his pb.UpdateHistorical
	if err := c.ShouldBindJSON(&his); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.historyService.UpdateHistoricals(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling UpdateHistoricals", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

// GetHistoricalByID godoc
// @Summary Get Historical by ID
// @Description Get Historical by its ID
// @Security BearerAuth
// @Tags Historical
// @Produce json
// @Param id path string true "Historical ID"
// @Success 200 {object} models.HistoricalResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/getBy/{id} [get]
func (h *historyHandler) GetHistoricalByID(c *gin.Context) {
	id := c.Param("id")

	his := pb.HistoricalId{
		Id: id,
	}

	req, err := h.historyService.GetHistoricalByID(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling GetHistoricalByID", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, req)
}

// DeleteHistorical godoc
// @Summary Delete Historical
// @Description Delete Historical by its ID
// @Security BearerAuth
// @Tags Historical
// @Produce json
// @Param id path string true "Historical ID"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/delete/{id} [delete]
func (h *historyHandler) DeleteHistorical(c *gin.Context) {
	id := c.Param("id")
	his := pb.HistoricalId{
		Id: id,
	}
	req, err := h.historyService.DeleteHistorical(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling DeleteHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

// ListHistorical godoc
// @Summary List Historical
// @Description Get a list of Historical with optional filtering
// @Security BearerAuth
// @Tags Historical
// @Produce json
// @Param filter query models.HistoricalList false "Filter Historical"
// @Success 200 {object} models.HistoricalListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/list [get]
func (h *historyHandler) ListHistorical(c *gin.Context) {
	var post pb.HistoricalList

	limit := c.Query("limit")
	offset := c.Query("offset")

	offsets, err := strconv.Atoi(offset)
	if err != nil {
		offsets = 0
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	post.Limit = int64(limits)
	post.Offset = int64(offsets)

	post.Country = c.Query("country")

	fmt.Println(post)
	res, err := h.historyService.ListHistorical(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while calling ListHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, res)
}

// SearchHistorical godoc
// @Summary List Historical
// @Description Get a list of Historical with optional filtering
// @Security BearerAuth
// @Tags Historical
// @Produce json
// @Param filter query models.HistoricalSearch false "Filter Historical"
// @Success 200 {object} models.HistoricalListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/list_search [get]
func (h *historyHandler) SearchHistorical(c *gin.Context) {
	var post pb.HistoricalSearch
	search := c.Query("search")

	post.Search = search

	res, err := h.historyService.SearchHistorical(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while calling SearchHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddHistoricalImage godoc
// @Summary Add Image to Historical
// @Description Add an image to a Historical by Historical ID
// @Security BearerAuth
// @Tags Historical
// @Accept json
// @Produce json
// @Param id path string true "historical att-n id"
// @Param file formData file true "Upload image"
// @Success 200 {object} models.Message
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/image/{id} [put]
func (h *historyHandler) UpdateHisImage(c *gin.Context) {
	var his pb.HistoricalImage

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

	his.Id = id
	his.Url = url

	req, err := h.historyService.AddHistoricalImage(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling AddHistoricalImage", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}
