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

type HistoryHandler interface {
	AddHistorical(c *gin.Context)
	UpdateHistoricals(c *gin.Context)
	GetHistoricalByID(c *gin.Context)
	DeleteHistorical(c *gin.Context)
	ListHistorical(c *gin.Context)
	SearchHistorical(c *gin.Context)
	AddHistoricalImage(c *gin.Context)
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

func (h *historyHandler) AddHistorical(c *gin.Context) {
	var his pb.Historical
	if err := c.ShouldBindJSON(&his); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.historyService.AddHistorical(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling AddHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}

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

func (h *historyHandler) ListHistorical(c *gin.Context) {
	var post pb.HistoricalList

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
	post.City = c.Query("city")
	res, err := h.historyService.ListHistorical(context.Background(), &post)
	if err != nil {
		h.logger.Error("Error occurred while calling ListHistorical", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

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

func (h *historyHandler) AddHistoricalImage(c *gin.Context) {
	var his pb.HistoricalImage
	if err := c.ShouldBindJSON(&his); err != nil {
		h.logger.Error("Error occurred while binding json", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req, err := h.historyService.AddHistoricalImage(context.Background(), &his)
	if err != nil {
		h.logger.Error("Error occurred while calling AddHistoricalImage", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, req)
}
