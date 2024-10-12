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
// @Summary Create a new Historical record
// @Description Create a new Historical record, including an image upload
// @Security BearerAuth
// @Tags Historical
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload image file"
// @Param name formData string true "Name of the historical site"
// @Param description formData string true "Description of the historical site"
// @Param city formData string true "City of the historical site"
// @Success 201 {object} nationality.HistoricalResponse "Historical record successfully created"
// @Failure 400 {object} models.Error "Bad request, validation error or invalid file"
// @Failure 500 {object} models.Error "Internal server error"
// @Router /historical/create [post]
func (h *historyHandler) AddHistorical(c *gin.Context) {
	log.Println("Request received for creating a historical record")

	var his models.Historical

	if err := c.ShouldBind(&his); err != nil {
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

	his.ImageURL = url

	res := pb.Historical{
		City:        his.City,
		Name:        his.Name,
		Description: his.Description,
		ImageUrl:    url,
	}

	req, err := h.historyService.AddHistorical(context.Background(), &res)
	if err != nil {
		h.logger.Error("Error occurred while calling AddHistorical", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
// @Param Update body nationality.UpdateHistorical true "Update Historical"
// @Success 200 {object} nationality.HistoricalResponse
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
// @Success 200 {object} nationality.HistoricalResponse
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
// @Success 200 {object} nationality.HistoricalListResponse
// @Failure 400 {object} models.Error
// @Failure 500 {object} models.Error
// @Router /historical/list [get]
func (h *historyHandler) ListHistorical(c *gin.Context) {

	var post pb.HistoricalList
	var offset int

	limit := c.Query("limit")
	p := c.Query("page")

	page, err := strconv.Atoi(p)
	if err != nil {
		page = 0
	}

	limits, err := strconv.Atoi(limit)
	if err != nil {
		limits = 10
	}

	if page == 0 || page == 1 {
		offset = 0
	} else {
		offset = (page - 1) * limits
	}

	post.Limit = int64(limits)
	post.Offset = int64(offset)

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
// @Tags Z-MUST-DELETE
// @Produce json
// @Param filter query nationality.HistoricalSearch false "Filter Historical"
// @Success 200 {object} nationality.HistoricalListResponse
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
